package database

import (
	"errors"

	"github.com/zlsgo/app_module/database/partition"
	"github.com/zlsgo/zdb"
)

// Storage 描述一个分区落在哪里。
//
// DB 是该分区的连接，Prefix 是该分区的表前缀。SQLite 按独立文件隔离，
// 此时 Prefix 与默认分区相同（表名不变）；其他数据库共用连接，靠 Prefix
// 把表名分开。
type Storage struct {
	DB     *zdb.DB
	Prefix string
}

var (
	// ErrPartitionsClosed 表示分区管理器已经关闭。
	ErrPartitionsClosed = partition.ErrClosed
	// ErrInvalidKey 表示分区键包含 ASCII 字母、数字、下划线、连字符之外的字符。
	ErrInvalidKey = errors.New("invalid partition key")
	// ErrPartitionLimit 表示分区数量已达配置上限。
	ErrPartitionLimit = partition.ErrLimit
)

// PartitionOptions 定义分区管理器配置。
//
// 内嵌的 Options 用于创建分区连接，缺省取数据库模块当前配置；
// 只有需要与默认连接不同的驱动或参数时才需要改它。
type PartitionOptions struct {
	Options
	// DB 默认分区的连接，缺省取数据库模块当前连接。
	DB *zdb.DB
	// Prefix 默认分区的表前缀。
	Prefix string
	// Max 非默认分区的数量上限，0 表示不限制。
	// 打开失败的分区不占配额，修好底层故障后可以重试。
	Max int
}

// Partitions 按分区键提供存储位置，不解释键的业务含义。
//
// 它只负责分区命名、懒打开、并发首访、数量上限和自有连接的关闭；
// 表结构注册与数据迁移由调用方负责。
//
// 零值不可用，必须通过 NewPartitions 构造。
type Partitions struct {
	pool *partition.Pool[Storage]
}

// NewPartitions 创建分区管理器。
//
// 默认复用数据库模块的当前连接与配置，通常只需要指定表前缀和上限：
//
//	partitions, err := database.NewPartitions(func(o *database.PartitionOptions) {
//		o.Prefix = "im_"
//		o.Max = 100
//	})
//	if err != nil {
//		return err
//	}
//	defer partitions.Close()
//
// 配置在构造时固定，之后的 Reload 或修改回调持有的配置都不会重定向它。
// 默认连接由调用方关闭，Close 只关闭管理器自己打开的连接。
func NewPartitions(opt ...func(*PartitionOptions)) (*Partitions, error) {
	o := PartitionOptions{
		Options: cloneOptions(CurrentOptions()),
		DB:      CurrentDB(),
	}
	for _, fn := range opt {
		if fn != nil {
			fn(&o)
		}
	}

	if o.DB == nil {
		return nil, errors.New("分区存储不可用: 数据库未初始化，请先加载 database 模块或在回调中设置 PartitionOptions.DB")
	}

	driver, err := resolveDriver(o.Options)
	if err != nil {
		return nil, err
	}
	if driver == "" {
		return nil, errors.New("分区存储不可用: 未配置数据库类型，请设置 PartitionOptions.Options.Driver")
	}

	base := Storage{DB: o.DB, Prefix: o.Prefix}
	open := sharedPartitionOpener(base)
	var closeFn func(Storage) error
	if driver == "sqlite" {
		open, closeFn = sqlitePartitionOpener(cloneOptions(o.Options), base)
	}

	pool, err := partition.NewPool(partition.Config[Storage]{
		Base:  base,
		Max:   o.Max,
		Open:  open,
		Close: closeFn,
	})
	if err != nil {
		return nil, err
	}
	return &Partitions{pool: pool}, nil
}

// Get 返回键对应的存储。空键返回默认分区，非法键返回 ErrInvalidKey。
//
// 首次访问会打开分区，并发首访只会打开一次；同一键反复调用返回同一个
// Storage。表结构注册与数据迁移由调用方负责。
func (p *Partitions) Get(key Key) (Storage, error) {
	if !validPartitionKey(key) {
		return Storage{}, ErrInvalidKey
	}
	return p.pool.Get(string(key))
}

// Default 返回默认分区：配置里的原连接与原表前缀。
func (p *Partitions) Default() Storage {
	return p.pool.Base()
}

// Keys 返回本管理器已打开的分区键（含默认空键），按字典序升序。
// 它不扫描磁盘，也不会发现其他进程创建的分区；正在打开或打开失败的键
// 不会出现。
func (p *Partitions) Keys() []Key {
	opened := p.pool.Keys()
	keys := make([]Key, len(opened))
	for i, key := range opened {
		keys[i] = Key(key)
	}
	return keys
}

// Close 拒绝新访问，等待在途创建并关闭本管理器创建的分区连接。
// 并发或重复调用会等待同一次关闭完成；默认分区的连接由调用方负责。
func (p *Partitions) Close() error {
	return p.pool.Close()
}
