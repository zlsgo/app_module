package database

import (
	"strings"
	"testing"

	"github.com/sohaha/zlsgo"
	"github.com/zlsgo/zdb/builder"
	"github.com/zlsgo/zdb/driver"
)

// snapshotDialect 保存并恢复 builder 的全局方言与登记表，避免污染其他用例。
// builder.DefaultDriver 是 zdb 的包级变量，本包所有用例都在同一个进程里跑，
// 测完必须还原，否则后面的 SQLite 用例会拿到这里的假方言。
func snapshotDialect(t *testing.T) {
	t.Helper()
	prev := builder.DefaultDriver
	prevStrict := SetStrictDialect(false)
	resetDialectTracking()
	t.Cleanup(func() {
		resetDialectTracking()
		SetStrictDialect(prevStrict)
		builder.DefaultDriver = prev
	})
}

// TestSetBuilderDialectTracksDialects 覆盖方言登记的常态路径：
// 首次设置、同方言重复设置都不产生冲突，且全局方言随之更新。
func TestSetBuilderDialectTracksDialects(t *testing.T) {
	tt := zlsgo.NewTest(t)
	snapshotDialect(t)

	// 登记表为空 = 本进程还没调用过 setBuilderDialect；此时 BuilderDialect
	// 返回 zdb 自带的默认方言（MySQL），而不是 nil
	tt.Equal(0, len(DialectsInUse()))
	tt.Equal(true, BuilderDialect() != nil)

	tt.NoError(setBuilderDialect(dialectA{}))
	tt.Equal(driver.MySQL, builder.DefaultDriver.Value())
	tt.Equal(1, len(DialectsInUse()))
	tt.Equal(driver.MySQL, DialectsInUse()[0])

	// 同一种方言再设置一次：不算冲突，登记表不重复追加
	tt.NoError(setBuilderDialect(dialectA{}))
	tt.Equal(1, len(DialectsInUse()))
}

// TestSetBuilderDialectConflictWarns 确认默认（非严格）模式下混用方言不会
// 报错，只是覆盖全局方言并打告警 —— 这是历史行为，不能因为加了检测就改变。
func TestSetBuilderDialectConflictWarns(t *testing.T) {
	tt := zlsgo.NewTest(t)
	snapshotDialect(t)

	tt.NoError(setBuilderDialect(dialectA{}))
	tt.NoError(setBuilderDialect(dialectB{}))

	// 后设置的方言生效，登记表记录两种
	tt.Equal(driver.SQLite, builder.DefaultDriver.Value())
	tt.Equal(2, len(DialectsInUse()))
	tt.Equal("MySQL + SQLite", dialectNames(DialectsInUse()))

	// 同样的方言冲突重复触发时不重复告警，但依然可以设置成功
	tt.NoError(setBuilderDialect(dialectA{}))
	tt.Equal(driver.MySQL, builder.DefaultDriver.Value())
}

// TestStrictDialectRejectsMixed 覆盖严格模式：进程内已有别的方言时，
// 切换方言必须失败，且不能污染 builder 的全局方言。
func TestStrictDialectRejectsMixed(t *testing.T) {
	tt := zlsgo.NewTest(t)
	snapshotDialect(t)

	old := SetStrictDialect(true)
	tt.Equal(false, old)

	tt.NoError(setBuilderDialect(dialectA{}))
	err := setBuilderDialect(dialectB{})
	tt.Equal(true, err != nil)
	if err != nil {
		tt.Equal(true, strings.Contains(err.Error(), "方言"))
		tt.Equal(true, strings.Contains(err.Error(), "MySQL"))
		tt.Equal(true, strings.Contains(err.Error(), "SQLite"))
	}

	// 被拒绝后全局方言仍是先前那一种，登记表也没有新增
	tt.Equal(driver.MySQL, builder.DefaultDriver.Value())
	tt.Equal(1, len(DialectsInUse()))

	// 关闭严格模式后同样的切换可以成功
	tt.Equal(true, SetStrictDialect(false))
	tt.NoError(setBuilderDialect(dialectB{}))
	tt.Equal(driver.SQLite, builder.DefaultDriver.Value())
}

// TestSetBuilderDialectNil 守住空方言这道校验：它必须报错而不是把
// builder.DefaultDriver 置空，否则后续所有 SQL 生成都会 panic。
func TestSetBuilderDialectNil(t *testing.T) {
	tt := zlsgo.NewTest(t)
	snapshotDialect(t)

	err := setBuilderDialect(nil)
	tt.Equal(true, err != nil)
	if err != nil {
		tt.Equal(true, strings.Contains(err.Error(), "Dialect"))
	}
	tt.Equal(0, len(DialectsInUse()))
}

// TestDialectNamesEmpty 覆盖 dialectNames 的空列表兜底文本。
func TestDialectNamesEmpty(t *testing.T) {
	tt := zlsgo.NewTest(t)
	tt.Equal("未知", dialectNames(nil))
}

// 编译期确认 fake 方言确实实现了 driver.Dialect。
var (
	_ driver.Dialect = dialectA{}
	_ driver.Dialect = dialectB{}
)
