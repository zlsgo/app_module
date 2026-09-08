//go:build nostatic
// +build nostatic

package html

import (
	"github.com/sohaha/zlsgo/znet"
	"github.com/zlsgo/app_module/html/zview"
)

// registerStatic 在 `-tags nostatic` 构建下的空实现：
// 仅注入 zview.Context 依赖，不引入 zcss.js / zview.js 静态资源，
// 也不注册对应的静态路由。
func registerStatic(r *znet.Engine, _ string) error {
	r.Use(func(c *znet.Context) {
		c.Injector().Map(zview.New(c))
		c.Next()
	})
	return nil
}
