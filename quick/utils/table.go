package utils

import (
	"strings"

	"github.com/sohaha/zlsgo/ztype"
)

func FillFilterTablePrefix(f ztype.Map, table string) ztype.Map {
	if table == "" {
		return f
	}

	for k := range f {
		if k == "" {
			continue
		}
		if !strings.ContainsRune(k, '.') {
			f[table+k] = f[k]
			delete(f, k)
		}
	}

	return f
}

func FillFieldsTablePrefix(f []string, table string) []string {
	if table == "" {
		return f
	}

	for i := range f {
		if !strings.ContainsRune(f[i], '.') {
			f[i] = table + f[i]
		}
	}

	return f
}
