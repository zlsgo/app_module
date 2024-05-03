package process

import (
	"github.com/zlsgo/app_module/database/hashid"
	"github.com/zlsgo/app_module/quick/define"
)

type Process struct {
	Hashid        *hashid.HashID
	Define        *define.Define
	AfterProcess  map[string][]AfterProcess
	BeforeProcess map[string][]BeforeProcess
	CryptKeys     map[string]CryptProcess
}
