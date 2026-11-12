# Database 数据库模块

Database 模块提供了统一的数据库连接与管理能力，支持多种数据库驱动、命名键分区以及固定连接事务等高级特性。

## 支持的数据库

- ✅ MySQL
- ✅ PostgreSQL（需要使用构建标签启用，见下文）
- ✅ SQLite（驱动名为 `sqlite`）

## 功能特性

- 🔌 多数据库支持（MySQL、PostgreSQL、SQLite）
- 🔄 数据库驱动管理
- 📊 基础连接池功能
- 🗃️ 数据库配置管理
- 🧩 业务无关的命名键分区（SQLite 文件 / 其他数据库表前缀）
- 🔒 固定连接事务（一致性读快照 / 写事务）
- 🛠️ 模块化设计

## 模块结构

```
database/
├── database.go        # 数据库初始化
├── driver.go          # 驱动管理
├── mysql.go           # MySQL 配置
├── postgres.go        # PostgreSQL 配置
├── sqlite3.go         # SQLite 配置
├── options.go         # 配置选项
├── service.go         # 服务封装
├── assign.go          # 单数据库连接
├── module.go          # 模块定义
├── tx.go              # 固定连接事务（读快照 / 写事务）
├── partition.go       # 分区公开 API 与存储路由
├── partition_key.go   # 分区键格式化与校验
├── partition_open.go  # SQLite 独立文件 / 共享连接前缀策略
├── partition/         # 懒打开、并发单飞、数量上限与关闭内核
└── partition.md       # 分区使用指南与跨分区查询说明
```

## 快速开始

### 基本使用

```go
package main

import (
    "fmt"

    "github.com/sohaha/zlsgo/zlog"
    "github.com/zlsgo/app_module/database"
    "github.com/zlsgo/app_core/service"
    "github.com/zlsgo/zdb"
)

func main() {
    app := service.NewApp()(nil)

    dbMod := database.New()
    err := service.InitModule([]service.Module{dbMod}, app)
    if err != nil {
        panic(err)
    }

    err = app.DI.InvokeWithErrorOnly(func(db *zdb.DB) {
        zlog.Info(db.QueryToMaps(`SELECT sqlite_version() AS version;`))

        var result []map[string]interface{}
        if err := db.Table("users").Find(&result); err != nil {
            zlog.Error(err)
            return
        }

        fmt.Printf("用户数量: %d\n", len(result))
    })
    if err != nil {
        panic(err)
    }
}
```

## 配置说明

### 完整配置示例

```yaml
database:
  driver: "mysql" # 数据库驱动类型: mysql, postgres, sqlite

  # MySQL 配置（当 driver 为 "mysql" 时使用）
  mysql:
    host: "localhost"
    port: 3306
    user: "root"
    password: "password"
    db_name: "myapp"
    charset: "utf8mb4"

  # PostgreSQL 配置（当 driver 为 "postgres" 时使用）
  postgres:
    host: "localhost"
    port: 5432
    user: "postgres"
    password: "password"
    db_name: "myapp"
    ssl_mode: "disable" # disable, require, verify-ca, verify-full

  # SQLite 配置（当 driver 为 "sqlite" 时使用）
  sqlite:
    path: "./data/app.db"

  # 模式配置
  mode:
    delete_column: false # 是否删除未使用的列
```

## 通用分区

分区模块按通用 `Key` 把数据落到不同的库文件或表名空间，不理解租户、应用等业务身份。只有显式调用 `NewPartitions` 才启用分区，默认连接与配置不会因分区功能本身发生变化。

完整用法与可运行示例见 [partition.md](./partition.md)。

### 基本示例

```go
// 默认复用数据库模块的当前连接与配置，只调分区自己的参数。
partitions, err := database.NewPartitions(func(o *database.PartitionOptions) {
    o.Prefix = "data_"
    o.Max = 100
})
if err != nil {
    return err
}
defer partitions.Close()

defaultStorage := partitions.Default() // 原连接与原表前缀
keyStorage, err := partitions.Get(database.KeyOf(1001))
if err != nil {
    return err
}

// 用 keyStorage.DB 与 keyStorage.Prefix 创建或绑定业务 schema。
// 表结构注册与迁移由调用方负责。
fmt.Println(defaultStorage.Prefix, keyStorage.Prefix, partitions.Keys())
```

### `PartitionOptions`

| 字段 | 说明 |
| --- | --- |
| `DB` | 默认分区的连接，零值取 `database.CurrentDB()` |
| `Prefix` | 默认分区的表前缀，其余分区按约定派生 |
| `Max` | 非默认分区数量上限，`0` 或负数表示不限制 |
| `Options` | 建分区连接用的数据库配置，零值取 `database.CurrentOptions()` |

可选回调按传入顺序执行，配置在构造时固定；之后的 `Reload` 或修改回调持有的配置都不会重定向已创建的管理器。默认连接由调用方关闭，`Close` 只关闭管理器自己打开的连接。

调用方自行建连时可以直接指定：

```go
partitions, err := database.NewPartitions(func(o *database.PartitionOptions) {
    o.DB = baseDB
    o.Prefix = "data_"
    o.Options = database.Options{Driver: "sqlite", Sqlite: &database.Sqlite{Path: "data/db.db"}}
})
```

### 分区键与命名约定

| 输入 / 后端 | 行为 |
| --- | --- |
| 空键 `""` | `Get("")` 与 `Default()` 返回传入的默认连接与表前缀，不占 `Max` 配额 |
| 非空 `Key` | 只允许 `[A-Za-z0-9_-]+`，不裁剪或改写，非法键返回 `ErrInvalidKey` |
| `KeyOf(int64)` | 仅做十进制转换，`KeyOf(0) == "0"`、`KeyOf(-1) == "-1"`，二者都不是默认分区 |
| SQLite，路径 `data/db.db` | 键 `tenant_West-2` 使用 `data/db_tenant_West-2.db`，表前缀不变 |
| SQLite，无扩展名路径 `data/db` | 使用 `data/db_tenant_West-2`，不接受内存库 |
| 其他数据库，前缀 `data_` | 共用默认连接，键 `tenant_West-2` 使用前缀 `data_tenant_West-2_` |

注意事项：

- `Keys()` 仅返回此管理器已成功打开的键，始终包含默认空键，按字典序排序。
- SQLite 分区是独立库文件：不能跨分区 `JOIN`，也不能在一个事务里读写多个分区；需要全局结果时按 `Keys()` 逐个查询后在应用层合并。
- `Max` 限制非默认分区数量，达到上限后已打开分区仍可用，新的非空键返回 `ErrPartitionLimit`；不会自动淘汰。
- 驱动名按 `normalizeName` 处理（去首尾空白 + 转小写）。
- 非 SQLite 的前缀按 `Prefix + Key + "_"` 原样拼接，建议 `Prefix` 以 `_` 结尾。
- `Close()` 拒绝后续访问，等待在途创建完成，再关闭自己创建的连接；关闭后 `Keys()` 仍可读取已打开键的快照。

## 固定连接事务

`database.RunPinnedTx` 将事务内全部语句绑定到同一连接，并按方言下发对应的开启语句：

| 驱动 | 写事务（`TxWrite`） | 读快照（`TxReadSnapshot`） |
| --- | --- | --- |
| sqlite | `BEGIN IMMEDIATE` | `BEGIN` |
| mysql | `BEGIN` | `SET TRANSACTION ISOLATION LEVEL REPEATABLE READ` + `START TRANSACTION WITH CONSISTENT SNAPSHOT` |
| postgres | `BEGIN` | `BEGIN ISOLATION LEVEL REPEATABLE READ` |

```go
err := database.RunPinnedTx(ctx, db, database.TxOptions{
    Mode:      database.TxReadSnapshot,
    Operation: "查询报表",
}, func(tx *database.Tx) error {
    rows, err := tx.QueryToMaps(ctx, "SELECT * FROM orders")
    if err != nil {
        return err
    }
    fmt.Println(rows)
    return nil
})
```

说明：

- `TxOptions.Driver` 留空时自动推导，也可显式传入 `sqlite` / `mysql` / `postgres`。
- `TxOptions.Operation` 非空时会作为前缀拼入错误信息，便于定位失败操作。
- `TxOptions.CleanupTimeout` 控制回滚等清理动作的超时，默认 1 秒。
- `BEGIN` 前的 `SET TRANSACTION` 只作用于该连接的下一个事务；开启失败时连接会被丢弃，防止隔离级别泄漏到连接池。

## 注意事项

### PostgreSQL 构建标签

PostgreSQL 驱动文件带有构建标签（`//go:build postgres`），默认构建条件下可能不包含 PostgreSQL 支持，需要在构建时启用对应 tag：

```bash
go build -tags postgres ./...
```

### `delete_column` 字段

配置项 `mode.delete_column` 会映射到 `database.Options.Mode.DelteColumn`（结构体字段名存在拼写，但 JSON key 为 `delete_column`，配置不受影响）。

### driver 选择规则

当未指定 `driver` 且同时配置多个数据库类型时，会直接报错并要求显式指定 `driver`。

### SQLite 连接策略

SQLite 仅使用单连接（`MaxOpenConns=1`），避免多连接导致的锁竞争问题。

### 默认方言

`builder.DefaultDriver` 会被最近一次初始化的数据库方言更新。

## 最佳实践

- 明确指定 `driver`，避免多数据库配置并存时产生歧义。
- 生产环境根据业务负载调整连接池与超时参数。
- 使用分区时提前规划键命名规则与上限，避免运行期命名冲突或资源耗尽。
- 需要跨分区统计时在应用层聚合，避免依赖 SQLite 不支持的跨库操作。