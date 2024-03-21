package restapi

import (
	"errors"
	"strings"

	"github.com/sohaha/zlsgo/zstring"
	"github.com/sohaha/zlsgo/ztype"
	"github.com/zlsgo/zdb"
	"github.com/zlsgo/zdb/builder"
)

func (s *SQL) parseExprs(d *builder.BuildCond, filter ztype.Map) (exprs []string, err error) {
	if len(filter) > 0 {
		for k := range filter {
			value := filter[k]
			if value == nil {
				exprs = append(exprs, k)
				continue
			}

			if k == "" {
				switch val := value.(type) {
				case func(*builder.BuildCond) string:
					exprs = append(exprs, val(d))
				case func() string:
					exprs = append(exprs, val())
				default:
					err = errors.New("unknown type")
					return
				}

				continue
			}
			upperKey := strings.ToUpper(k)
			v := ztype.New(value)
			if upperKey == "" || upperKey == "$OR" || upperKey == "$AND" {
				m := v.Map()
				cexprs, err := s.parseExprs(d, m)
				if err != nil {
					return nil, err
				}

				switch upperKey {
				case "$OR":
					exprs = append(exprs, d.Or(cexprs...))
				default:
					exprs = append(exprs, d.And(cexprs...))
				}

				continue
			}

			f := strings.SplitN(zstring.TrimSpace(k), " ", 2)
			l := len(f)
			if l != 2 {
				switch val := v.Value().(type) {
				case ztype.Maps, []ztype.Map:
					for _, v := range ztype.ToSlice(val).Maps() {
						cexprs, err := s.parseExprs(d, v)
						if err != nil {
							return nil, err
						}
						exprs = append(exprs, d.Or(cexprs...))
					}
				case []interface{}, []string, []int64, []int32, []int16, []int8, []int, []uint64, []uint32, []uint16, []uint8, []uint, []float64, []float32:
					exprs = append(exprs, d.In(f[0], ztype.ToSlice(val).Value()...))
				default:
					exprs = append(exprs, d.EQ(f[0], val))
				}
			} else {
				switch strings.ToLower(f[1]) {
				default:
					err = errors.New("Unknown operator: " + f[1])
					return
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
				case "!=":
					exprs = append(exprs, d.NE(f[0], v.Value()))
				case "like":
					exprs = append(exprs, d.Like(f[0], v.Value()))
				case "in":
					exprs = append(exprs, d.In(f[0], v.Slice().Value()...))
				case "notin":
					exprs = append(exprs, d.NotIn(f[0], v.Slice().Value()...))
				case "between":
					s := v.Slice()
					exprs = append(exprs, d.Between(f[0], s.Index(0), s.Index(1)))
				}
			}
		}
	}

	return
}

func (s *SQL) Insert(table string, data ztype.Map) (lastId interface{}, err error) {
	return s.db.Insert(table, data)
}

func (s *SQL) InsertMany(table string, data ztype.Maps) (lastIds []interface{}, err error) {
	ids, err := s.db.BatchInsert(table, data)
	if err != nil {
		return []interface{}{}, err
	}
	for _, id := range ids {
		lastIds = append(lastIds, id)
	}
	return
}

func (s *SQL) Delete(table string, filter ztype.Map, fn ...func(*CondOptions) error) (int64, error) {
	o := CondOptions{}
	for _, f := range fn {
		if err := f(&o); err != nil {
			return 0, err
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

		return nil
	})
}

func (s *SQL) First(table string, filter ztype.Map, fn ...func(*CondOptions) error) (ztype.Map, error) {
	rows, err := s.Find(table, filter, func(so *CondOptions) error {
		so.Limit = 1
		if len(fn) > 0 {
			return fn[0](so)
		}
		return nil
	})

	if err == nil && rows.Len() > 0 {
		return rows[0], nil
	}

	return ztype.Map{}, err
}

func (s *SQL) Find(table string, filter ztype.Map, fn ...func(*CondOptions) error) (ztype.Maps, error) {
	o := CondOptions{}
	for i := range fn {
		if fn[i] == nil {
			continue
		}
		if err := fn[i](&o); err != nil {
			return nil, err
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

func (s *SQL) Pages(table string, page, pagesize int, filter ztype.Map, fn ...func(*CondOptions) error) (ztype.Maps, PageInfo, error) {
	o := CondOptions{}
	for _, f := range fn {
		if err := f(&o); err != nil {
			return nil, PageInfo{}, err
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

func (s *SQL) Update(table string, data ztype.Map, filter ztype.Map, fn ...func(*CondOptions) error) (int64, error) {
	o := CondOptions{}
	for _, f := range fn {
		if err := f(&o); err != nil {
			return 0, err
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
