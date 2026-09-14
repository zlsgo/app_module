//go:build !nostatic

// 静态资源由脚本生成：go generate ./html
//go:generate go run gen.go

package html

import (
	"strconv"
	"strings"
	"time"

	"github.com/sohaha/zlsgo/zfile"
	"github.com/sohaha/zlsgo/znet"
	"github.com/zlsgo/app_module/html/zview"
)

// registerStatic 注册 zview.Context 依赖注入，并通过 znet 直接输出由
// html/gen.go 生成的静态资源（zcss.js / zview.js），前缀路由可通过
// Options.StaticPrefix 自定义（默认 /__static_html）。
//
// 使用 `-tags nostatic` 构建可完全剔除这些静态资源（参见 static_disabled.go）。
func registerStatic(r *znet.Engine, prefix string) error {
	r.Use(func(c *znet.Context) {
		c.Injector().Map(zview.New(c))
		c.Next()
	})

	if prefix == "" {
		prefix = defaultStaticPrefix
	}
	prefix = "/" + strings.Trim(prefix, "/")
	if prefix == "/" {
		prefix = ""
	}
	prefix += "/"

	now := time.Now()
	for i := range staticFiles {
		file := &staticFiles[i]
		r.GET(prefix+file.Name, func(c *znet.Context) {
			if !znet.Utils.IsModified(c, now) {
				return
			}
			c.SetContentType(zfile.GetMimeType(file.Name, file.Data))
			c.Byte(200, file.Data)
		})
		r.HEAD(prefix+file.Name, func(c *znet.Context) {
			if !znet.Utils.IsModified(c, now) {
				return
			}
			c.SetContentType(zfile.GetMimeType(file.Name, file.Data))
			c.SetHeader("Content-Length", strconv.Itoa(len(file.Data)))
			c.Byte(200, nil)
		})
	}
	return nil
}
