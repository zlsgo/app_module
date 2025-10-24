package html

import "github.com/zlsgo/app_module/html/el"

type Options struct {
	ErrorPage *el.Element
}

func (Options) ConfKey() string {
	return "html"
}

func (Options) DisableWrite() bool {
	return true
}

var options = Options{}
