package restapi

import (
	"github.com/sohaha/zlsgo/zutil"
)

var options Options

func New(o ...func(options *Options)) *Module {
	options = zutil.Optional(Options{
		Prefix: "__",
	}, o...)
	return &Module{
		options: options,
	}
}
