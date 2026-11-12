# 数据库分区使用指南

`database` 模块的分区能力按通用 `Key` 把数据落到不同的库文件或表名空间。
它不解释租户、应用、年份等业务身份，只负责「同一个键恒定映射到同一个存储」。

适用场景：

- 按应用 / 租户拆 SQLite 文件，降低跨应用的写锁排队。
- 按业务键拆 MySQL / PostgreSQL 表名前缀，控制单表体量。
- 需要在运行时按需创建分区，而不是启动时一次性建好全部库表。

不适用场景：

- 需要跨分区 JOIN、跨分区事务或全局唯一约束。
- 期望自动迁移存量数据（表结构注册、数据搬迁仍由调用方负责）。
- 需要把分区当成权限隔离（表前缀不是数据库权限边界）。

## 概念

| 名称 | 含义 |
| --- | --- |
| `Key` | 分区键。空键是默认分区；非空键只允许 ASCII 字母、数字、`_`、`-` |
| `Storage` | 分区落点：`DB` 是连接，`Prefix` 是表前缀 |
| `Partitions` | 分区管理器：懒打开、并发单飞、数量上限、连接生命周期 |
| `PartitionOptions` | 构造配置：`DB`、`Prefix`、`Max`，以及建分区连接用的 `Options` |

## 快速开始

### 复用数据库模块的连接

数据库模块 `Load` 之后，`CurrentDB()` 与 `CurrentOptions()` 已有值，可以零参数构造：

```go
partitions, err := database.NewPartitions()
if err != nil {
    return err
}
defer partitions.Close()

defaultStorage := partitions.Default() // 原连接 + 原表前缀
keyStorage, err := partitions.Get(database.KeyOf(1001))
if err != nil {
    return err
}
fmt.Println(defaultStorage.Prefix, keyStorage.Prefix)
```

通常至少设置表前缀与上限：

```go
partitions, err := database.NewPartitions(func(o *database.PartitionOptions) {
    o.Prefix = "app_"
    o.Max = 100 // 0 或负数表示不限制
})
```

`Max` 是硬上限，不会自动淘汰旧分区：达到上限后，已打开的分区继续可用，新的
非空键返回 `ErrPartitionLimit`；默认分区不受影响。要释放名额只能 `Close()` 后
重建管理器。

### 使用自建连接

默认连接不是由数据库模块创建时，显式提供连接与建分区连接所需的配置：

```go
partitions, err := database.NewPartitions(func(o *database.PartitionOptions) {
    o.DB = baseDB
    o.Prefix = "app_"
    o.Options = database.Options{
        Driver: "sqlite",
        Sqlite: &database.Sqlite{Path: "data/app.db"},
    }
})
```

默认连接由创建者关闭，`Partitions.Close()` 只关闭它自己打开的分区连接。

## 按 Key 自动路由（以多租户为例）

分区键可以是租户、应用、年份等任意业务 ID。业务代码不应该在每个 handler 里
手写 `Get`，用一个薄封装把业务键收口：

```go
type TenantStore struct {
    partitions *database.Partitions
}

func (s *TenantStore) Save(ctx context.Context, tenantID int64, note Note) error {
    storage, err := s.partitions.Get(database.KeyOf(tenantID))
    if err != nil {
        return err
    }
    _, err = storage.DB.ExecContext(ctx,
        "INSERT INTO "+storage.Prefix+"notes (content) VALUES (?)", note.Content)
    return err
}
```

要点：

- `Get` 幂等且自带并发单飞，可以在每次请求直接调用，不必自己缓存 `*zdb.DB`。
- 每个分区是独立的库或表空间，表结构也要在每个分区建一次。第一次访问该键时
  建表并缓存结果（失败下次重试），不要每个请求都跑迁移。
- 租户 ID 从鉴权上下文取，不要直接信任请求体。
- 没有业务键语境的全局数据落默认分区：`partitions.Default()` 或 `partitions.Get("")`。
- `KeyOf(0) == "0"` 是一个普通分区，不是默认分区；需要默认分区时必须用空键。
- 保存失败要区分 `ErrInvalidKey` / `ErrPartitionLimit`，不要静默落回默认分区。
- SQLite 下每个分区是独立文件，跨分区统计不能 `JOIN`，按 `Keys()` 逐分区查询后合并。

完整可运行版本见 `database/partition_example_test.go` 的
`ExamplePartitions_multiTenant`：两个租户各写一份数据，互相不可见。

## API 速查

| 方法 | 说明 |
| --- | --- |
| `NewPartitions(opt ...func(*PartitionOptions)) (*Partitions, error)` | 构造管理器；配置在构造时复制并固定 |
| `Get(key Key) (Storage, error)` | 取分区；空键返回默认分区，首次访问会懒打开 |
| `Default() Storage` | 直接取默认分区，不打开任何东西 |
| `Keys() []Key` | 已成功打开的键，含空键，按字典序升序 |
| `Close() error` | 拒绝新访问、等待在途创建、关闭自有连接 |

错误判定：

```go
_, err := partitions.Get("../etc")
errors.Is(err, database.ErrInvalidKey)       // 键含非法字符
errors.Is(err, database.ErrPartitionLimit)   // 达到 Max
errors.Is(err, database.ErrPartitionsClosed) // 已 Close
```

键工具：

- `database.KeyOf(int64)`：十进制格式化，`KeyOf(0) == "0"`，不是默认分区。
- `Key.Int64()`：反解十进制键；空键或非十进制键返回错误，不校验规范拼写。
- 字符串键可以直接写：`partitions.Get(database.Key("tenant-west"))`。

## 后端行为

| 后端 | 落点 | 表名示例（前缀 `app_`，键 `1001`） |
| --- | --- | --- |
| SQLite，路径 `data/app.db` | 独立文件 `data/app_1001.db`，表前缀不变 | `app_users`（在 `app_1001.db` 里） |
| SQLite，路径 `data/app` | 独立文件 `data/app_1001` | `app_users` |
| MySQL / PostgreSQL | 共用默认连接，表前缀变为 `app_1001_` | `app_1001_users` |

约束：

- SQLite 内存库（`:memory:`、`file::memory:?cache=shared`）不支持文件分区。
- 非 SQLite 的前缀按 `Prefix + Key + "_"` 原样拼接，不会自动补下划线；
  建议 `Prefix` 以 `_` 结尾。`Prefix` 为空且键是纯数字时，表名会以数字开头，
  MySQL 需要额外引用，共用连接时请给非空前缀。
- 键会原样进入文件名或表前缀，需自行保证部署环境里不会碰撞，并留足表名长度。
- 非 SQLite 只共享连接与方言，不会为分区新建数据库。

## 跨分区查询（重要）

SQLite 的每个分区是独立的 `*zdb.DB` 与独立文件，一条 SQL 只能看到当前连接所在的库：

- 不能跨分区 `JOIN`，也不能在一个分区的事务里读写另一个分区。
- 唯一索引、外键、聚合都只在单个分区内成立。
- 不要用 `ATTACH DATABASE` 绕过；分区管理器不保证路径与文件生命周期，
  这类连接也不会跟着 `Close()` 一起收口。

MySQL / PostgreSQL 虽然共用一个连接，但每个分区是独立表名前缀。跨前缀 JOIN 等于
把不同业务键的数据混在一起，同样应该避免。

需要全局视图时，按 `Keys()` 逐个分区查询后在应用层合并：

```go
var all ztype.Maps
for _, key := range partitions.Keys() {
    storage, err := partitions.Get(key)
    if err != nil {
        return err
    }
    rows, err := storage.DB.QueryToMaps("SELECT * FROM " + storage.Prefix + "orders")
    if err != nil {
        return err
    }
    all = append(all, rows...)
}
```

默认分区也是一个独立落点：写进默认分区的数据不会出现在任何非空键的 `Storage` 里。
需要同时读默认分区时，显式查询 `Default()`。

## 生命周期与并发

- `Partitions` 必须用 `NewPartitions` 构造，零值不可用。
- `Get` 可以并发调用；同一个键的首次访问只打开一次，其余调用等待同一结果。
- 打开失败不会占用 `Max` 配额，修好底层路径后可以直接重试同一个键。
- 驱动在打开过程中 panic 会转成错误，同样不占配额，可以重试。
- `Keys()` 不扫描磁盘，只返回本管理器已经打开的键；其他进程创建的文件不会出现。
  首次打开还在进行中的键不会出现，打开失败的分区会被移除，下一次 `Get` 重新尝试。
- `Close()` 并发调用会等待同一次关闭完成；关闭后 `Get` 返回 `ErrPartitionsClosed`，
  `Keys()` 仍可读取快照，`Default()` 仍返回原存储。

## 配合 model 使用

分区的 `Storage` 可以直接作为模型存储的输入：

```go
storage, err := partitions.Get(database.KeyOf(appID))
if err != nil {
    return err
}

schemas := model.NewSchemas(di, model.NewSQL(storage.DB, storage.Prefix), model.DefaultSchemaOptions)
// 之后按现有方式注册 schema、执行迁移。
```

表结构注册与数据迁移由业务侧负责；分区管理器不会自动建表或搬数据。

## 完整可运行示例

`database/partition_example_test.go` 包含两个示例：

- `ExamplePartitions`：创建临时 SQLite 库，写入默认分区与两个应用分区并打印结果。
- `ExamplePartitions_multiTenant`：用 `zcache.FastCache` 缓存分区连接，演示按租户
  保存时自动路由。

运行：

```bash
go test ./database -run ExamplePartitions -v
```
