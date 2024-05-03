package utils

import (
	"github.com/zlsgo/app_module/quick/define"
)

func IsDisableMigratioField(d *define.Define, name string) bool {
	for n := range d.Fields {
		if name != n {
			continue
		}
		if d.Fields[n].Options.DisableMigration {
			return true
		}
	}
	return false
}
