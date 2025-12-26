package model

import (
	"strings"

	"github.com/sohaha/zlsgo/ztype"
	"github.com/zlsgo/zdb/builder"
)

func collectKeyTuples(rows ztype.Maps, keys []string) [][]any {
	if len(rows) == 0 || len(keys) == 0 {
		return nil
	}

	seen := make(map[string]struct{}, len(rows))
	tuples := make([][]any, 0, len(rows))
	keyParts := make([]string, len(keys))

	for i := range rows {
		tuple := make([]any, len(keys))
		ok := true

		for k := range keys {
			v := rows[i].Get(keys[k]).Value()
			if v == nil {
				ok = false
				break
			}
			tuple[k] = v
			keyParts[k] = ztype.ToString(v)
		}

		if !ok {
			continue
		}

		mapKey := strings.Join(keyParts, relationKeySeparator)
		if _, exists := seen[mapKey]; exists {
			continue
		}
		seen[mapKey] = struct{}{}
		tuples = append(tuples, tuple)
	}

	return tuples
}

func buildCompositeFilter(schemaKeys []string, tuples [][]any) ztype.Map {
	if len(schemaKeys) == 0 || len(tuples) == 0 {
		return nil
	}

	if len(schemaKeys) == 1 {
		values := make([]any, 0, len(tuples))
		repeat := make(map[any]struct{}, len(tuples))
		for _, t := range tuples {
			if len(t) == 0 {
				continue
			}
			v := t[0]
			if v == nil {
				continue
			}
			if _, ok := repeat[v]; ok {
				continue
			}
			repeat[v] = struct{}{}
			values = append(values, v)
		}
		if len(values) == 0 {
			return nil
		}
		if len(values) == 1 {
			return ztype.Map{schemaKeys[0]: values[0]}
		}
		return ztype.Map{schemaKeys[0]: values}
	}

	filter := NewFilter().Cond(func(c *builder.BuildCond) string {
		exprs := make([]string, 0, len(tuples))
		for _, t := range tuples {
			if len(t) < len(schemaKeys) {
				continue
			}

			andExprs := make([]string, 0, len(schemaKeys))
			for i := range schemaKeys {
				andExprs = append(andExprs, c.EQ(schemaKeys[i], t[i]))
			}
			if len(andExprs) == 0 {
				continue
			}
			if len(andExprs) == 1 {
				exprs = append(exprs, andExprs[0])
				continue
			}
			exprs = append(exprs, c.And(andExprs...))
		}

		if len(exprs) == 0 {
			return ""
		}
		if len(exprs) == 1 {
			return exprs[0]
		}
		return c.Or(exprs...)
	})

	return ztype.Map(filter)
}
