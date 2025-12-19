package model

import (
	"errors"
	"fmt"
	"strings"

	"github.com/sohaha/zlsgo/zstring"
	"github.com/sohaha/zlsgo/ztype"
	"github.com/sohaha/zlsgo/zutil"
	"github.com/zlsgo/zdb"
	"github.com/zlsgo/zdb/builder"
)

const (
	placeHolder    = "$"
	placeHolderOR  = "$OR"
	placeHolderAND = "$AND"
	maxParseDepth  = 50 // 最大递归深度限制
)

func (s *SQL) parseExprs(d *builder.BuildCond, filter ztype.Map) (exprs []string, err error) {
	return s.parseExprsWithDepth(d, filter, 0)
}

func (s *SQL) parseExprsWithDepth(d *builder.BuildCond, filter ztype.Map, depth int) (exprs []string, err error) {
	if len(filter) == 0 {
		return nil, nil
	}

	if depth >= maxParseDepth {
		return nil, fmt.Errorf("parseExprs: max recursion depth (%d) exceeded at depth %d", maxParseDepth, depth)
	}

	exprs = make([]string, 0, len(filter))

	for k, value := range filter {
		if k == "" {
			if exprs, err = parseExprsBuildCond(d, value, exprs); err != nil {
				return
			}
			continue
		}

		upperKey := strings.ToUpper(k)
		isPlaceHolderOR := upperKey == placeHolderOR
		isPlaceHolderAND := upperKey == placeHolderAND
		isPlaceHolder := isPlaceHolderOR || isPlaceHolderAND

		if strings.Contains(k, placeHolder) && !isPlaceHolder {
			exprs, err = parseExprsBuildCond(d, value, exprs)
			if err != nil {
				return
			}
			continue
		}

		v := ztype.New(value)

		if isPlaceHolder {
			m := v.Map()
			var cexprs []string
			cexprs, err = s.parseExprsWithDepth(d, m, depth+1)
			if err != nil {
				return nil, err
			}

			if len(cexprs) > 0 {
				if isPlaceHolderOR {
					exprs = append(exprs, d.Or(cexprs...))
				} else {
					exprs = append(exprs, d.And(cexprs...))
				}
			}
			continue
		}

		trimmedKey := k
		if strings.ContainsAny(k, " \t\n\r") {
			trimmedKey = zstring.TrimSpace(k)
		}

		f := strings.SplitN(trimmedKey, " ", 2)

		if len(f) != 2 {
			switch val := v.Value().(type) {
			case ztype.Maps, []ztype.Map:
				m := ztype.ToSlice(val).Maps()
				e := make([]string, 0, len(m))
				for _, mapItem := range m {
					cexprs, err := s.parseExprsWithDepth(d, mapItem, depth+1)
					if err != nil {
						return nil, err
					}
					e = append(e, cexprs...)
				}
				if len(e) > 0 {
					exprs = append(exprs, d.Or(e...))
				}
			case []interface{}, []string, []int64, []int32, []int16, []int8, []int, []uint64, []uint32, []uint16, []uint8, []uint, []float64, []float32:
				values := ztype.ToSlice(v.Value()).Value()
				valuesLen := len(values)
				if valuesLen == 0 {
					continue
				} else if valuesLen == 1 {
					exprs = append(exprs, d.EQ(f[0], values[0]))
				} else {
					exprs = append(exprs, d.In(f[0], values...))
				}
			default:
				exprs = append(exprs, d.EQ(f[0], val))
			}
		} else {
			operator := strings.ToUpper(f[1])
			switch operator {
			case "=":
				exprs = append(exprs, d.EQ(f[0], v.Value()))
			case ">":
				exprs = append(exprs, d.GT(f[0], v.Value()))
			case ">=":
				exprs = append(exprs, d.GE(f[0], v.Value()))
			case "<":
				exprs = append(exprs, d.LT(f[0], v.Value()))
			case "<=":
				exprs = append(exprs, d.LE(f[0], v.Value()))
			case "!=", "<>":
				values := ztype.ToSlice(v.Value()).Value()
				if len(values) == 1 {
					exprs = append(exprs, d.NE(f[0], values[0]))
				} else {
					exprs = append(exprs, d.NotIn(f[0], values...))
				}
			case "LIKE":
				exprs = append(exprs, d.Like(f[0], v.Value()))
			case "IN":
				exprs = append(exprs, d.In(f[0], v.SliceValue()...))
			case "NOTIN", "NOT IN":
				exprs = append(exprs, d.NotIn(f[0], v.SliceValue()...))
			case "IS NULL":
				exprs = append(exprs, d.IsNull(f[0]))
			case "IS NOT NULL":
				exprs = append(exprs, d.IsNotNull(f[0]))
			case "BETWEEN":
				s := v.SliceValue()
				if len(s) != 2 {
					return nil, errors.New("BETWEEN operator need two values")
				}
				exprs = append(exprs, d.Between(f[0], s[0], s[1]))
			default:
				return nil, errors.New("Unknown operator: " + f[1])
			}
		}
	}

	return exprs, nil
}

func (s *SQL) Insert(table string, fields []string, data ztype.Map, fn ...func(*InsertOptions)) (lastId interface{}, err error) {
	o := zutil.Optional(InsertOptions{}, fn...)
	return s.db.Insert(table, data, o.Options)
}

func (s *SQL) InsertMany(table string, fields []string, data ztype.Maps, fn ...func(*InsertOptions)) (lastIds []interface{}, err error) {
	o := zutil.Optional(InsertOptions{}, fn...)
	ids, err := s.db.BatchInsert(table, data, o.Options)
	if err != nil {
		return []interface{}{}, err
	}

	lastIds = make([]interface{}, len(ids))
	for i, id := range ids {
		lastIds[i] = id
	}
	return
}

func (s *SQL) Delete(table string, fields []string, filter ztype.Map, fn ...func(*CondOptions)) (int64, error) {
	o := acquireCondOptions()
	defer releaseCondOptions(o)
	for _, f := range fn {
		if f != nil {
			f(o)
		}
	}

	return s.db.Delete(table, func(b *builder.DeleteBuilder) error {
		var fieldPrefix string
		hasJoin := len(o.Join) > 0
		if hasJoin {
			fieldPrefix = table + "."
		}

		exprs, err := s.parseExprs(b.Cond, fillFilterTablePrefix(filter, fieldPrefix))
		if err != nil {
			return err
		}

		if len(exprs) > 0 {
			b.Where(exprs...)
		}

		b.OrderBy(sqlOrderBy(o.OrderBy, fieldPrefix)...)

		if o.Limit > 0 {
			b.Limit(o.Limit)
		}

		return nil
	})
}

func (s *SQL) First(table string, fields []string, filter ztype.Map, fn ...func(*CondOptions)) (ztype.Map, error) {
	rows, err := s.Find(table, fields, filter, func(so *CondOptions) {
		so.Limit = 1
		so.Offset = 0
		if len(fn) > 0 {
			fn[0](so)
		}
	})

	if err == nil && rows.Len() > 0 {
		return rows[0], nil
	}

	return ztype.Map{}, err
}

func (s *SQL) Find(table string, fields []string, filter ztype.Map, fn ...func(*CondOptions)) (ztype.Maps, error) {
	o := acquireCondOptions()
	defer releaseCondOptions(o)
	for _, f := range fn {
		if f != nil {
			f(o)
		}
	}
	items, err := s.db.Find(table, func(b *builder.SelectBuilder) error {
		var fieldPrefix string
		hasJoin := len(o.Join) > 0
		if hasJoin {
			fieldPrefix = table + "."
		}

		if len(o.Fields) > 0 {
			b.Select(fillFieldsTablePrefix(o.Fields, fieldPrefix)...)
		}

		exprs, err := s.parseExprs(b.Cond, fillFilterTablePrefix(filter, fieldPrefix))
		if err != nil {
			return err
		}

		if len(exprs) > 0 {
			b.Where(exprs...)
		}

		if hasJoin {
			for _, v := range o.Join {
				b.JoinWithOption(v.ModelOptions, b.As(v.Table, v.As), v.Expr)
			}
		}

		b.OrderBy(sqlOrderBy(o.OrderBy, fieldPrefix)...)

		if o.Limit > 0 {
			b.Limit(o.Limit)
		}

		if o.Offset > 0 {
			b.Offset(o.Offset)
		}

		if len(o.GroupBy) > 0 {
			b.GroupBy(fillFieldsTablePrefix(o.GroupBy, fieldPrefix)...)
		}

		return nil
	})

	if err != nil && err != zdb.ErrNotFound {
		return items, err
	}

	return items, nil
}

func (s *SQL) Pages(table string, fields []string, page, pagesize int, filter ztype.Map, fn ...func(*CondOptions)) (ztype.Maps, PageInfo, error) {
	o := acquireCondOptions()
	defer releaseCondOptions(o)
	for _, f := range fn {
		if f != nil {
			f(o)
		}
	}

	rows, p, err := s.db.Pages(table, page, pagesize, func(b *builder.SelectBuilder) error {
		var fieldPrefix string
		hasJoin := len(o.Join) > 0
		if hasJoin {
			fieldPrefix = table + "."
		}

		if len(o.Fields) > 0 {
			b.Select(fillFieldsTablePrefix(o.Fields, fieldPrefix)...)
		}

		exprs, err := s.parseExprs(b.Cond, fillFilterTablePrefix(filter, fieldPrefix))
		if err != nil {
			return err
		}

		if len(exprs) > 0 {
			b.Where(exprs...)
		}

		b.OrderBy(sqlOrderBy(o.OrderBy, fieldPrefix)...)

		if hasJoin {
			for _, v := range o.Join {
				b.JoinWithOption(v.ModelOptions, b.As(v.Table, v.As), v.Expr)
			}
		}

		if o.Limit > 0 {
			b.Limit(o.Limit)
		}

		if o.Offset > 0 {
			b.Offset(o.Offset)
		}

		if len(o.GroupBy) > 0 {
			b.GroupBy(fillFieldsTablePrefix(o.GroupBy, fieldPrefix)...)
		}

		return nil
	})

	if err != nil && err != zdb.ErrNotFound {
		return rows, PageInfo{}, err
	}

	return rows, PageInfo{
		p,
	}, nil
}

func (s *SQL) Update(table string, fields []string, data ztype.Map, filter ztype.Map, fn ...func(*CondOptions)) (int64, error) {
	o := acquireCondOptions()
	defer releaseCondOptions(o)
	for _, f := range fn {
		if f != nil {
			f(o)
		}
	}

	return s.db.Update(table, data, func(b *builder.UpdateBuilder) error {
		var fieldPrefix string
		hasJoin := len(o.Join) > 0
		if hasJoin {
			fieldPrefix = table + "."
		}

		exprs, err := s.parseExprs(b.Cond, fillFilterTablePrefix(filter, fieldPrefix))
		if err != nil {
			return err
		}
		if len(exprs) > 0 {
			b.Where(exprs...)
		}

		if o.Limit > 0 {
			b.Limit(o.Limit)
		}

		b.OrderBy(sqlOrderBy(o.OrderBy, fieldPrefix)...)

		return nil
	})
}
