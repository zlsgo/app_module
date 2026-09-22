package database

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"errors"
	"time"

	"github.com/sohaha/zlsgo/zerror"
	"github.com/sohaha/zlsgo/ztype"
	"github.com/zlsgo/zdb"
)

type (
	// TxMode 描述事务用途。TxReadSnapshot 在支持的方言上建立一个稳定读快照，
	// TxWrite 使用方言对应的写事务开启方式。
	TxMode uint8

	// TxOptions 控制事务执行行为
	TxOptions struct {
		Mode           TxMode
		Driver         string
		Operation      string
		CleanupTimeout time.Duration
	}

	// Tx 绑定一条已开启事务的固定连接。Exec/Query 直接落在该连接上，
	// 保证 BEGIN 前的一次性设置与事务内全部语句共享同一连接状态。
	Tx struct {
		conn *sql.Conn
	}
)

const (
	// TxWrite 写事务
	TxWrite TxMode = iota
	// TxReadSnapshot 一致性读快照事务
	TxReadSnapshot
)

// TxStatements 返回方言对应的事务开启语句。
// 返回值依次执行在同一固定连接上，用于在 BEGIN 前设置一次性隔离级别。
func TxStatements(driverName string, mode TxMode) ([]string, error) {
	switch normalizeName(driverName) {
	case "sqlite":
		if mode == TxReadSnapshot {
			return []string{"BEGIN"}, nil
		}
		return []string{"BEGIN IMMEDIATE"}, nil
	case "mysql":
		if mode == TxReadSnapshot {
			return []string{
				"SET TRANSACTION ISOLATION LEVEL REPEATABLE READ",
				"START TRANSACTION WITH CONSISTENT SNAPSHOT",
			}, nil
		}
		return []string{"BEGIN"}, nil
	case "postgres", "postgresql":
		if mode == TxReadSnapshot {
			return []string{"BEGIN ISOLATION LEVEL REPEATABLE READ"}, nil
		}
		return []string{"BEGIN"}, nil
	default:
		return nil, errors.New("不支持的数据库驱动[" + driverName + "]")
	}
}

// TxDriverName 将 zdb 方言类型映射为 TxStatements 可识别的驱动名
func TxDriverName(db *zdb.DB) (string, error) {
	if db == nil {
		return "", errors.New("数据库实例为空")
	}
	d := db.GetDriver()
	if d == nil {
		return "", errors.New("数据库方言为空")
	}
	name := d.Value().String()
	switch name {
	case "MySQL", "Doris":
		return "mysql", nil
	case "PostgreSQL":
		return "postgres", nil
	case "SQLite":
		return "sqlite", nil
	default:
		return "", errors.New("不支持的数据库方言[" + name + "]")
	}
}

// Exec 执行写语句
func (tx *Tx) Exec(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	return tx.conn.ExecContext(ctx, query, args...)
}

// Query 执行查询语句
func (tx *Tx) Query(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	return tx.conn.QueryContext(ctx, query, args...)
}

// QueryToMaps 执行查询并返回 map 结果集
func (tx *Tx) QueryToMaps(ctx context.Context, query string, args ...interface{}) (ztype.Maps, error) {
	rows, err := tx.conn.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	result, _, err := zdb.ScanToMap(rows)
	return result, err
}

// RunPinnedTx 在固定连接上执行方言感知的事务。
// zdb.DB.Transaction 只能发起标准 BEGIN，无法在 BEGIN 前设置
// 一次性隔离级别或选择 BEGIN IMMEDIATE；本函数补齐该能力。
func RunPinnedTx(ctx context.Context, db *zdb.DB, opts TxOptions, fn func(tx *Tx) error) error {
	if ctx == nil {
		ctx = context.Background()
	}
	if db == nil {
		return txPhaseError(opts.Operation, "校验事务失败", errors.New("数据库实例为空"))
	}
	if fn == nil {
		return txPhaseError(opts.Operation, "校验事务失败", errors.New("事务回调为空"))
	}
	if err := ctx.Err(); err != nil {
		return err
	}

	driverName := opts.Driver
	if driverName == "" {
		var err error
		driverName, err = TxDriverName(db)
		if err != nil {
			return txPhaseError(opts.Operation, "识别数据库驱动失败", err)
		}
	}

	begin, err := TxStatements(driverName, opts.Mode)
	if err != nil {
		return txPhaseError(opts.Operation, "识别数据库驱动失败", err)
	}

	sqlDB, err := db.GetSQLDB()
	if err != nil {
		return txPhaseError(opts.Operation, "固定数据库连接失败", err)
	}
	conn, err := sqlDB.Conn(ctx)
	if err != nil {
		return txPhaseError(opts.Operation, "固定数据库连接失败", err)
	}
	defer conn.Close()

	cleanup := opts.CleanupTimeout
	if cleanup <= 0 {
		cleanup = time.Second
	}

	for i, stmt := range begin {
		if _, err := conn.ExecContext(ctx, stmt); err != nil {
			cleanupErr := discardBadConn(opts.Operation, conn, err)
			// SET TRANSACTION 只作用于该连接上的下一个事务；
			// 后续 BEGIN 失败时丢弃连接，防止隔离级别泄漏。
			if i > 0 && !errors.Is(err, driver.ErrBadConn) {
				cleanupErr = errors.Join(cleanupErr, discardConn(opts.Operation, conn))
			}
			return errors.Join(txPhaseError(opts.Operation, "开启事务失败", err), cleanupErr)
		}
	}

	tx := &Tx{conn: conn}
	active := true
	defer func() {
		if active {
			_ = rollbackTx(opts.Operation, conn, cleanup, false)
		}
	}()

	if err := fn(tx); err != nil {
		cleanupErr := rollbackTx(opts.Operation, conn, cleanup, false)
		active = false
		if opts.Mode == TxReadSnapshot {
			if parentErr := ctx.Err(); parentErr != nil {
				return errors.Join(parentErr, cleanupErr)
			}
		}
		return errors.Join(err, cleanupErr)
	}
	if opts.Mode == TxReadSnapshot {
		if parentErr := ctx.Err(); parentErr != nil {
			cleanupErr := rollbackTx(opts.Operation, conn, cleanup, false)
			active = false
			return errors.Join(parentErr, cleanupErr)
		}
	}

	if _, err := conn.ExecContext(ctx, "COMMIT"); err != nil {
		commitErr := txPhaseError(opts.Operation, "提交事务失败", err)
		cleanupErr := rollbackTx(opts.Operation, conn, cleanup, true)
		active = false
		if opts.Mode == TxReadSnapshot {
			if parentErr := ctx.Err(); parentErr != nil {
				return errors.Join(parentErr, cleanupErr)
			}
		}
		return errors.Join(commitErr, cleanupErr)
	}
	active = false
	return nil
}

func rollbackTx(operation string, conn *sql.Conn, timeout time.Duration, discardAlways bool) error {
	cleanupCtx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	_, rollbackErr := conn.ExecContext(cleanupCtx, "ROLLBACK")

	var discardErr error
	if rollbackErr != nil || discardAlways {
		discardErr = discardConn(operation, conn)
	}
	return errors.Join(
		txPhaseError(operation, "回滚事务失败", rollbackErr),
		discardErr,
	)
}

func discardBadConn(operation string, conn *sql.Conn, err error) error {
	if !errors.Is(err, driver.ErrBadConn) {
		return nil
	}
	return discardConn(operation, conn)
}

func discardConn(operation string, conn *sql.Conn) error {
	err := conn.Raw(func(interface{}) error { return driver.ErrBadConn })
	if err == nil || errors.Is(err, driver.ErrBadConn) {
		return nil
	}
	return txPhaseError(operation, "丢弃数据库连接失败", err)
}

// txPhaseError 按项目错误风格包装事务各阶段错误，
// Operation 非空时作为前缀拼入阶段描述。
func txPhaseError(operation, phase string, err error) error {
	if err == nil {
		return nil
	}
	if operation != "" {
		phase = operation + ": " + phase
	}
	return zerror.With(err, phase)
}
