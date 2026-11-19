package database

import (
	"errors"
	"sync"

	"github.com/sohaha/zlsgo/zlog"
	"github.com/zlsgo/zdb/builder"
	"github.com/zlsgo/zdb/driver"
)

// zdb 把「方言」放在 builder.DefaultDriver 这个包级变量里，builder 每次
// 构造语句（Select/Insert/Update/Delete…）都读它来决定占位符与标识符
// 引号风格。也就是说方言本来就是进程级全局的，而不是连接级。
//
// 于是同一进程里只要出现第二种方言，后设置的就会覆盖先设置的：先建的
// 那个库之后生成的 SQL 会带错语法（最典型的是 Postgres 的 $1 占位符
// 被发到 MySQL）。隔离需要 zdb 内部支持（builder 按库持有方言），本包
// 做不到，但可以把隐患变成「可见、可控」：
//
//   - 记录本进程设置过的方言，出现第二种时打一条告警（每种方言只报一次，
//     避免配置热重载刷屏）；
//   - 提供 SetStrictDialect，让无法接受这种隐性错乱的部署直接拒绝启动；
//   - 暴露 BuilderDialect / DialectsInUse 供启动自检与测试断言。
var (
	dialectMu     sync.Mutex
	dialectOrder  []driver.Typ
	dialectSeen   = make(map[driver.Typ]struct{}, 2)
	dialectWarned = make(map[driver.Typ]struct{}, 2)
	strictDialect bool
)

// BuilderDialect 返回当前对 zdb builder 生效的方言。
//
// 注意它读的是 builder.DefaultDriver，也就是「最后设置的那个」，与具体
// 连接无关：多方言进程里它只代表最近一次初始化。zdb 自带一个默认方言
// （MySQL），所以本进程没设置过时它也不是 nil —— 想判断「是否设置过」
// 请看 DialectsInUse。
func BuilderDialect() driver.Dialect {
	dialectMu.Lock()
	defer dialectMu.Unlock()
	return builder.DefaultDriver
}

// DialectsInUse 返回本进程设置过的全部方言，按首次出现顺序排列。
//
// 长度大于 1 就说明进程内混用了多种方言，此时只有最后设置的那一种能
// 正常工作，其余连接的 SQL 语法是错的。
func DialectsInUse() []driver.Typ {
	dialectMu.Lock()
	defer dialectMu.Unlock()
	out := make([]driver.Typ, len(dialectOrder))
	copy(out, dialectOrder)
	return out
}

// SetStrictDialect 打开或关闭「方言严格模式」，返回原值。
//
// 严格模式下，如果本进程已经用过一种方言，再设置另一种方言会直接返回
// 错误，而不是悄悄覆盖全局方言。默认关闭，保持与历史行为（最后设置生效）
// 兼容；单库部署打开它可以尽早暴露配置错误。
func SetStrictDialect(enabled bool) bool {
	dialectMu.Lock()
	defer dialectMu.Unlock()
	old := strictDialect
	strictDialect = enabled
	return old
}

// resetDialectTracking 清空方言登记，仅供测试使用。
func resetDialectTracking() {
	dialectMu.Lock()
	defer dialectMu.Unlock()
	dialectOrder = nil
	dialectSeen = make(map[driver.Typ]struct{}, 2)
	dialectWarned = make(map[driver.Typ]struct{}, 2)
}

// setBuilderDialect 把方言设为 zdb builder 的全局方言，并登记使用记录。
func setBuilderDialect(dialect driver.Dialect) error {
	if dialect == nil {
		return errors.New("数据库配置失败: driver Dialect 为空")
	}

	dialectMu.Lock()
	defer dialectMu.Unlock()

	t := dialect.Value()
	_, seen := dialectSeen[t]

	// 已登记过另一种方言：全局 dialect 一旦切换，先前那个库的 SQL 就会
	// 带错语法。严格模式直接拒绝；默认模式只告警，保持历史行为不变。
	if !seen && len(dialectSeen) > 0 {
		if strictDialect {
			return errors.New("数据库配置失败: 同一进程已使用 " + dialectNames(dialectOrder) +
				" 方言, 无法再切换为 " + t.String() +
				"; zdb 的 SQL builder 共用全局方言, 混合方言请拆分为独立进程")
		}
		if _, warned := dialectWarned[t]; !warned {
			dialectWarned[t] = struct{}{}
			zlog.Warn("数据库方言冲突: 进程内已使用 " + dialectNames(dialectOrder) +
				" 方言, 现在切换为 " + t.String() +
				"; zdb 的 SQL builder 共用全局方言, 先创建的连接会拿到错误语法。若这是配置热重载切换请忽略, 多方言并存请拆进程或用 SetStrictDialect 拦截")
		}
	}

	if !seen {
		dialectSeen[t] = struct{}{}
		dialectOrder = append(dialectOrder, t)
	}

	builder.DefaultDriver = dialect
	return nil
}

// dialectNames 把方言列表拼成 "MySQL + SQLite" 这样的可读文本。
func dialectNames(ts []driver.Typ) string {
	if len(ts) == 0 {
		return "未知"
	}
	s := ""
	for i, t := range ts {
		if i > 0 {
			s += " + "
		}
		s += t.String()
	}
	return s
}
