package html

import "github.com/zlsgo/app_module/html/el"

// defaultStaticPrefix 是 zcss/zview 静态资源的默认路由前缀。
const defaultStaticPrefix = "/__static_html"

type Options struct {
	ErrorPage *el.Element

	// StaticPrefix 指定 zcss.js / zview.js 等静态资源的路由前缀，
	// 为空时使用默认值 /__static_html。前端引用时需与此保持一致。
	StaticPrefix string
}

func (Options) ConfKey() string {
	return "html"
}

func (Options) DisableWrite() bool {
	return true
}

var options = Options{}
