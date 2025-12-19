package model

import (
	"github.com/sohaha/zlsgo/zarray"
	"github.com/sohaha/zlsgo/zjson"
	"github.com/sohaha/zlsgo/zstring"
	"github.com/sohaha/zlsgo/ztype"
)

// PageData 分页数据结构
type PageData struct {
	Items    ztype.Maps `json:"items"`
	Page     PageInfo   `json:"page"`
	pagesize uint       `json:"-"`
}

// String 返回格式化的 JSON 字符串
func (p *PageData) String() string {
	json, err := zjson.Marshal(p)
	if err != nil {
		return ""
	}
	return zstring.Bytes2String(zjson.Format(json))
}

// Map 对分页结果进行批量映射转换
// fn: 转换函数，parallel: 并发数（默认为 pagesize）
func (p *PageData) Map(fn func(index int, item ztype.Map) ztype.Map, parallel ...uint) *PageData {
	if len(parallel) == 0 {
		parallel = []uint{p.pagesize}
	}
	p.Items = zarray.Map(p.Items, fn, parallel[0])

	return p
}

// Pages 分页查询（公开 API）
func Pages[T filter](
	m *Schema,
	page, pagesize int,
	filter T,
	fn ...func(*CondOptions),
) (*PageData, error) {
	return pages(m, page, pagesize, getFilter(m, filter), true, fn...)
}

// pages 分页查询内部实现
// cryptId: 是否需要加密/解密 ID
func pages(
	m *Schema,
	page, pagesize int,
	filter ztype.Map,
	cryptId bool,
	fn ...func(*CondOptions),
) (pagedata *PageData, err error) {
	if cryptId {
		_ = m.DeCrypt(filter)
	}

	var (
		childRelationson map[string][]string
		foreignKeys      []string
		data             = &PageData{pagesize: uint(pagesize)}
	)

	rows, pages, err := m.Storage.Pages(
		m.GetTableName(),
		m.GetFields(),
		page,
		pagesize,
		filter,
		func(so *CondOptions) {
			if len(fn) > 0 {
				fn[0](so)
			}

			childRelationson, foreignKeys = relationson(m, so)
			if len(so.Fields) > 0 && len(so.Join) == 0 {
				so.Fields = m.filterFields(so.Fields)
			} else if len(so.Fields) == 0 {
				so.Fields = m.GetFields()
			}
		},
	)

	data.Items = rows
	data.Page = pages

	if err != nil {
		return data, err
	}

	data.Items, err = handlerRelationson(m, data.Items, childRelationson, foreignKeys)
	if err != nil {
		return data, err
	}

	afterProcess := m.afterProcess
	if len(afterProcess) == 0 {
		return data, nil
	}

	for i := range data.Items {
		row := &data.Items[i]
		for k, v := range afterProcess {
			if _, ok := (*row)[k]; ok {
				(*row)[k], err = v[0](row.Get(k).String())
				if err != nil {
					return data, err
				}
			}
		}

		if cryptId && *m.define.Options.CryptID {
			_ = m.EnCrypt(row)
		}
	}

	return data, nil
}
