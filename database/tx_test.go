package database

import (
	"context"
	"errors"
	"strings"
	"testing"

	"github.com/sohaha/zlsgo"
	"github.com/zlsgo/app_module/internal/testutil"
)

func TestTxStatements(t *testing.T) {
	tt := zlsgo.NewTest(t)

	for _, c := range []struct {
		driver string
		mode   TxMode
		want   []string
	}{
		{"sqlite", TxWrite, []string{"BEGIN IMMEDIATE"}},
		{"sqlite3", TxReadSnapshot, []string{"BEGIN"}},
		{" SQLite ", TxReadSnapshot, []string{"BEGIN"}},
		{"mysql", TxWrite, []string{"BEGIN"}},
		{"MySQL", TxReadSnapshot, []string{
			"SET TRANSACTION ISOLATION LEVEL REPEATABLE READ",
			"START TRANSACTION WITH CONSISTENT SNAPSHOT",
		}},
		{"postgres", TxWrite, []string{"BEGIN"}},
		{"PostgreSQL", TxReadSnapshot, []string{"BEGIN ISOLATION LEVEL REPEATABLE READ"}},
	} {
		got, err := TxStatements(c.driver, c.mode)
		tt.NoError(err)
		tt.Equal(c.want, got)
	}

	_, err := TxStatements("oracle", TxWrite)
	tt.Equal(true, err != nil)
	tt.Equal(true, strings.Contains(err.Error(), "oracle"))
}

func TestTxDriverName(t *testing.T) {
	tt := zlsgo.NewTest(t)

	_, err := TxDriverName(nil)
	tt.Equal(true, err != nil)

	db := testutil.NewMemoryDB(t)
	defer func() { _ = db.Close() }()

	name, err := TxDriverName(db)
	tt.NoError(err)
	tt.Equal("sqlite", name)
}

func TestRunPinnedTxCommitAndHelpers(t *testing.T) {
	tt := zlsgo.NewTest(t)

	db := testutil.NewMemoryDB(t)
	defer func() { _ = db.Close() }()

	_, err := db.Exec("CREATE TABLE pinned (id INTEGER PRIMARY KEY, name TEXT)")
	tt.NoError(err)

	err = RunPinnedTx(context.Background(), db, TxOptions{Operation: "创建用户"}, func(tx *Tx) error {
		driverName, err := TxDriverName(db)
		tt.NoError(err)
		tt.Equal("sqlite", driverName)

		if _, err := tx.Exec(context.Background(), "INSERT INTO pinned (name) VALUES (?)", "a"); err != nil {
			return err
		}
		if _, err := tx.Exec(context.Background(), "INSERT INTO pinned (name) VALUES (?)", "b"); err != nil {
			return err
		}

		rows, err := tx.Query(context.Background(), "SELECT name FROM pinned ORDER BY id")
		if err != nil {
			return err
		}
		defer func() { _ = rows.Close() }()
		var names []string
		for rows.Next() {
			var name string
			if err := rows.Scan(&name); err != nil {
				return err
			}
			names = append(names, name)
		}
		if err := rows.Err(); err != nil {
			return err
		}
		tt.Equal([]string{"a", "b"}, names)

		maps, err := tx.QueryToMaps(context.Background(), "SELECT name FROM pinned WHERE name = ?", "b")
		if err != nil {
			return err
		}
		tt.Equal(1, len(maps))
		tt.Equal("b", maps[0].Get("name").String())

		// 查询语法错误要能透出，而不是被忽略
		if _, err := tx.QueryToMaps(context.Background(), "SELECT * FROM not_exists"); err == nil {
			return errors.New("expected query error")
		}
		return nil
	})
	tt.NoError(err)

	maps, err := db.QueryToMaps("SELECT name FROM pinned ORDER BY id")
	tt.NoError(err)
	tt.Equal(2, len(maps))
	tt.Equal("a", maps[0].Get("name").String())
	tt.Equal("b", maps[1].Get("name").String())
}

func TestRunPinnedTxRollbackOnCallbackError(t *testing.T) {
	tt := zlsgo.NewTest(t)

	db := testutil.NewMemoryDB(t)
	defer func() { _ = db.Close() }()

	_, err := db.Exec("CREATE TABLE pinned_rollback (id INTEGER PRIMARY KEY, name TEXT)")
	tt.NoError(err)

	expected := errors.New("业务失败")
	err = RunPinnedTx(context.Background(), db, TxOptions{Operation: "回滚用例"}, func(tx *Tx) error {
		if _, err := tx.Exec(context.Background(), "INSERT INTO pinned_rollback (name) VALUES (?)", "x"); err != nil {
			return err
		}
		return expected
	})
	tt.Equal(true, errors.Is(err, expected))

	maps, err := db.QueryToMaps("SELECT name FROM pinned_rollback")
	tt.NoError(err)
	tt.Equal(0, len(maps))
}

func TestRunPinnedTxReadSnapshot(t *testing.T) {
	tt := zlsgo.NewTest(t)

	db := testutil.NewMemoryDB(t)
	defer func() { _ = db.Close() }()

	_, err := db.Exec("CREATE TABLE pinned_snapshot (id INTEGER PRIMARY KEY, name TEXT)")
	tt.NoError(err)
	_, err = db.Exec("INSERT INTO pinned_snapshot (name) VALUES ('before')")
	tt.NoError(err)

	err = RunPinnedTx(context.Background(), db, TxOptions{Mode: TxReadSnapshot, Operation: "快照读"}, func(tx *Tx) error {
		maps, err := tx.QueryToMaps(context.Background(), "SELECT name FROM pinned_snapshot")
		if err != nil {
			return err
		}
		if len(maps) != 1 || maps[0].Get("name").String() != "before" {
			return errors.New("快照读结果异常")
		}
		return nil
	})
	tt.NoError(err)
}

func TestRunPinnedTxValidation(t *testing.T) {
	tt := zlsgo.NewTest(t)

	db := testutil.NewMemoryDB(t)
	defer func() { _ = db.Close() }()

	noop := func(*Tx) error { return nil }

	tt.Equal(true, RunPinnedTx(context.Background(), nil, TxOptions{}, noop) != nil)
	tt.Equal(true, RunPinnedTx(context.Background(), db, TxOptions{}, nil) != nil)

	// 显式传入不支持的驱动时在开事务前就失败，并给出阶段描述
	err := RunPinnedTx(context.Background(), db, TxOptions{Driver: "oracle"}, noop)
	tt.Equal(true, err != nil)
	tt.Equal(true, strings.Contains(err.Error(), "识别数据库驱动失败"))

	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	tt.Equal(context.Canceled, RunPinnedTx(ctx, db, TxOptions{}, noop))
}

func TestRunPinnedTxCommitFailureDiscardsConn(t *testing.T) {
	tt := zlsgo.NewTest(t)

	db := testutil.NewMemoryDB(t)
	defer func() { _ = db.Close() }()

	// 回调里主动回滚，外层 COMMIT 必然失败，从而走到「提交失败 + 丢弃连接」的清理路径
	err := RunPinnedTx(context.Background(), db, TxOptions{Operation: "提交失败"}, func(tx *Tx) error {
		_, err := tx.Exec(context.Background(), "ROLLBACK")
		return err
	})
	tt.Equal(true, err != nil)

	// 连接被丢弃后，库仍然可用
	_, err = db.Exec("CREATE TABLE after_commit_failure (id INTEGER PRIMARY KEY)")
	tt.NoError(err)
}
