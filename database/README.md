# Database 数据库模块

Database 模块提供了统一的数据库连接和管理功能，支持多种数据库驱动和高级特性。

## 支持的数据库

- ✅ MySQL
- ✅ PostgreSQL（需要使用构建标签启用，见下文）
- ✅ SQLite（驱动名为 `sqlite`）

## 功能特性

- 🔌 多数据库支持（MySQL、PostgreSQL、SQLite）
- 🔄 数据库驱动管理
- 📊 基础连接池功能
- 🗃️ 数据库配置管理
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
    // 初始化应用
    app := service.NewApp()(nil)

    // 数据库模块
    dbMod := database.New()

    // 注册全部模块
    err := service.InitModule([]service.Module{dbMod}, app)
    if err != nil {
        panic(err)
    }

    // 使用数据库对象
    err = app.DI.InvokeWithErrorOnly(func(db *zdb.DB) {
        zlog.Info(db.QueryToMaps(`SELECT sqlite_version() AS version;`))

        // 使用数据库
        var result []map[string]interface{}
        err := db.Table("users").Find(&result)
        if err != nil {
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

### 配置文件

数据库模块支持完整的配置文件，可以根据不同的数据库类型进行配置。

#### 完整配置示例

```yaml
# 数据库配置
database:
  driver: "mysql"                     # 数据库驱动类型: mysql, postgres, sqlite
  
  # MySQL 配置（当 driver 为 "mysql" 时使用）
  mysql:
    host: "localhost"                # 数据库主机地址
    port: 3306                       # 数据库端口
    user: "root"                     # 数据库用户名
    password: "password"             # 数据库密码
    db_name: "myapp"                 # 数据库名称
    charset: "utf8mb4"               # 字符集
  
  # PostgreSQL 配置（当 driver 为 "postgres" 时使用）
  postgres:
    host: "localhost"                # 数据库主机地址
    port: 5432                       # 数据库端口
    user: "postgres"                 # 数据库用户名
    password: "password"             # 数据库密码
    db_name: "myapp"                 # 数据库名称
    ssl_mode: "disable"              # SSL 模式: disable, require, verify-ca, verify-full
  
  # SQLite 配置（当 driver 为 "sqlite" 时使用）
  sqlite:
    path: "./data/app.db"            # 数据库文件路径
  
  # 模式配置
  mode:
    delete_column: false             # 是否删除未使用的列
```

### 固定连接事务

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

- `TxOptions.Driver` 留空时自动推导，也可显式传入 `sqlite` / `mysql` / `postgres`。
- `TxOptions.Operation` 非空时会作为前缀拼入错误信息，便于定位失败操作。
- `TxOptions.CleanupTimeout` 控制回滚等清理动作的超时，默认 1 秒。
- `BEGIN` 前的 `SET TRANSACTION` 只作用于该连接的下一个事务；开启失败时连接会被丢弃，防止隔离级别泄漏到连接池。

## 注意事项

### PostgreSQL 构建标签

PostgreSQL 驱动文件带有构建标签：

- `pkg/app_module/database/postgres.go` 使用 `//go:build postgres`

因此在默认构建条件下可能不会包含 PostgreSQL 支持，需要在构建时启用对应 tag。

### delete_column 字段

配置项 `mode.delete_column` 会映射到 `database.Options.Mode.DelteColumn`（结构体字段名存在拼写，但 JSON key 为 `delete_column`，配置不受影响）。

### driver 选择规则

当未指定 `driver` 且同时配置多个数据库类型时，会直接报错并要求显式指定 `driver`。


### SQLite 连接策略

SQLite 仅使用单连接（`MaxOpenConns=1`），避免多连接导致的锁竞争问题。

### 默认方言

`builder.DefaultDriver` 会被最近一次初始化的数据库方言更新。
