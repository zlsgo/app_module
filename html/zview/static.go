package zview

import (
	"fmt"

	"github.com/sohaha/zlsgo/znet"
)

type Context struct {
	accept string
	url    string
	origin string
	target string
	c      *znet.Context
}

func (z *Context) String() string {
	return fmt.Sprintf("accept: %s, url: %s, origin: %s, target: %s", z.accept, z.url, z.origin, z.target)
}

func (z *Context) GetUrl(fallback ...string) string {
	if z.url == "" && len(fallback) > 0 {
		return fallback[0]
	}
	return z.url
}

func (z *Context) GetOrigin() string {
	return z.origin
}

func (z *Context) GetTarget() string {
	return z.target
}

func (z *Context) Is() bool {
	return z.url != ""
}

func (z *Context) IsPartial() bool {
	return z.accept == "text/html+partial"
}

func (z *Context) SetTitle(title string) {
	z.c.SetHeader("Z-Title", title)
}

func (z *Context) SetRedirect(location string) {
	z.c.SetHeader("Z-Redirect", location)
	if !z.Is() {
		z.c.Redirect(location)
	}
}

func (z *Context) SetLocation(url string) {
	z.c.SetHeader("Z-Location", url)
}

func (z *Context) SetHistory(url string) {
	z.c.SetHeader("Z-History", url)
}

func (z *Context) SetHistoryReplace(url string) {
	z.c.SetHeader("Z-History-Replace", url)
}

func (z *Context) SetSwap(value Swap) {
	z.c.SetHeader("Z-Swap", string(value))
}

func (z *Context) SetTarget(value string) {
	z.c.SetHeader("Z-Target", value)
}

func New(c *znet.Context) *Context {
	return &Context{
		accept: c.GetHeader("Accept"),
		url:    c.GetHeader("Z-Url"),
		origin: c.GetHeader("Z-Origin"),
		target: c.GetHeader("Z-Target"),
		c:      c,
	}
}
