package restapi

import (
	"github.com/sohaha/zlsgo/zutil"
)

func New(o ...func(options *Options)) *Module {
	options := zutil.Optional(Options{
		Prefix:      "__",
		MaxPageSize: defaultMaxPageSize,
	}, o...)
	return &Module{
		options: options,
	}
}
