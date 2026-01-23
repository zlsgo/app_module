package html

import (
	"embed"
	"time"

	"github.com/sohaha/zlsgo/zfile"
	"github.com/sohaha/zlsgo/znet"
	"github.com/zlsgo/app_module/html/zview"
)

//go:embed static
var static embed.FS

func registerStatic(r *znet.Engine) error {
	r.Use(func(c *znet.Context) {
		c.Injector().Map(zview.New(c))
		c.Next()
	})

	now := time.Now()
	r.GET("/__static_html/*", func(c *znet.Context) {
		path := c.GetParam("*")
		f, err := static.ReadFile("static/" + path)
		if err != nil {
			c.String(404, "Not Found")
			return
		}

		if !znet.Utils.IsModified(c, now) {
			return
		}

		ctype := zfile.GetMimeType(path, f)
		c.SetContentType(ctype)
		c.Byte(200, f)
	})
	return nil
}
