package html

import (
	"embed"
	"fmt"
	"time"

	"github.com/sohaha/zlsgo/zfile"
	"github.com/sohaha/zlsgo/znet"
)

type ZViewJS struct {
	accept string
	Url    string
	Origin string
	Target string
	c      *znet.Context
}

func (z *ZViewJS) String() string {
	return fmt.Sprintf("accept: %s, url: %s, origin: %s, target: %s", z.accept, z.Url, z.Origin, z.Target)
}

func (z *ZViewJS) Is() bool {
	return z.Url != ""
}

func (z *ZViewJS) IsPartial() bool {
	return z.accept == "text/html+partial"
}

func (z *ZViewJS) SetTitle(title string) {
	z.c.SetHeader("Z-Title", title)
}

func (z *ZViewJS) Redirect(location string) {
	z.c.SetHeader("Z-Redirect", location)
	if !z.Is() {
		z.c.Redirect(location)
	}
}

func (z *ZViewJS) Location(url string) {
	z.c.SetHeader("Z-Location", url)
}

func (z *ZViewJS) History(url string) {
	z.c.SetHeader("Z-History", url)
}

func (z *ZViewJS) Swap(value string) {
	z.c.SetHeader("Z-Swap", value)
}

func (z *ZViewJS) SwapPush(value string) {
	z.c.SetHeader("Z-Swap-Push", value)
}

func newZHTML(c *znet.Context) *ZViewJS {
	return &ZViewJS{
		accept: c.GetHeader("Accept"),
		Url:    c.GetHeader("Z-Url"),
		Origin: c.GetHeader("Z-Origin"),
		Target: c.GetHeader("Z-Target"),
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
