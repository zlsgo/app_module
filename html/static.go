package html

import (
	"embed"
	"fmt"
	"time"

	"github.com/sohaha/zlsgo/zfile"
	"github.com/sohaha/zlsgo/znet"
)

type ZHTML struct {
	accept string
	Url    string
	Origin string
	Target string
	c      *znet.Context
}

func (z *ZHTML) String() string {
	return fmt.Sprintf("accept: %s, url: %s, origin: %s, target: %s", z.accept, z.Url, z.Origin, z.Target)
}

func (z *ZHTML) IsPartial() bool {
	return z.accept == "text/html+partial"
}

func (z *ZHTML) SetTitle(title string) {
	z.c.SetHeader("z-title", title)
}

func (z *ZHTML) Redirect(location string) {
	z.c.SetHeader("z-location", location)
}

func (z *ZHTML) History(url string) {
	z.c.SetHeader("z-history", url)
}

func (z *ZHTML) Swap(value string) {
	z.c.SetHeader("z-swap", value)
}

func (z *ZHTML) SwapPush(value string) {
	z.c.SetHeader("z-swap-push", value)
}

func newZHTML(c *znet.Context) *ZHTML {
	return &ZHTML{
		accept: c.GetHeader("Accept"),
		Url:    c.GetHeader("z-url"),
		Origin: c.GetHeader("z-origin"),
		Target: c.GetHeader("z-target"),
		c:      c,
	}
}

//go:embed static
var static embed.FS

func registerStatic(r *znet.Engine) error {
	r.Use(func(c *znet.Context) {
		c.Injector().Map(newZHTML(c))
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
