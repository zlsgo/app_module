package html

import (
	"github.com/zlsgo/app_module/html/el"
	"github.com/zlsgo/app_module/html/zview"
)

// defaultStaticPrefix 是 zcss/zview 静态资源的默认路由前缀。
const defaultStaticPrefix = "/__static_html"

type Options struct {
	// ErrorPage is rendered with HTTP 500 when a handler returns an error or
	// the page element cannot be rendered.
	ErrorPage *el.Element

	// StaticPrefix 指定 zcss.js / zview.js 等静态资源的路由前缀，
	// 为空时使用默认值 /__static_html。前端引用时需与此保持一致。
	StaticPrefix string
}

// ZViewJS is kept as a compatibility alias for the name used by older
// documentation. New code may import html/zview and use zview.Context.
type ZViewJS = zview.Context

func (Options) ConfKey() string {
	return "html"
}

func (Options) DisableWrite() bool {
	return true
}
