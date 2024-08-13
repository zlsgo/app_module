package restapi

import (
	"github.com/sohaha/zlsgo/zerror"
	"github.com/sohaha/zlsgo/znet"
	"github.com/sohaha/zlsgo/ztype"
	"github.com/zlsgo/app_module/model"
)

// FindById 获取一条数据
func FindById(
	c *znet.Context,
	mod *model.Model,
	id string,
	fn func(o *model.CondOptions),
) (ztype.Map, error) {
	res, err := find(c, mod, id, model.Filter{}, fn)
	if err != nil {
		return nil, err
	}

	return res.(ztype.Map), nil
}

// Page 获取分页数据
func Page(
	c *znet.Context,
	mod *model.Model,
	filter model.Filter,
	fn func(o *model.CondOptions),
) (*model.PageData, error) {
	res, err := find(c, mod, "", filter, fn)
	if err != nil {
		return nil, err
	}

	return res.(*model.PageData), nil
}

// Find 获取多条数据
func Find(
	c *znet.Context,
	mod *model.Model,
	filter model.Filter,
	fn func(o *model.CondOptions),
) (ztype.Maps, error) {
	res, err := find(c, mod, "*", filter, fn)
	if err != nil {
		return nil, err
	}

	return res.(ztype.Maps), nil
}

// Insert 数据，如果 json 数据是数组则批量添加
func Insert(
	c *znet.Context,
	mod *model.Model,
	fn func(data ztype.Map) (ztype.Map, error),
	o ...func(io *model.InsertOptions),
) (any, error) {
	j, err := c.GetJSONs()
	if err != nil {
		return nil, err
	}
	var id any
	if j.IsArray() {
		data := j.Maps()
		if fn != nil {
			for i := range data {
				ndata, err := fn(data.Index(i))
				if err != nil {
					return nil, err
				}
				data[i] = ndata
			}
		}

		id, err = mod.InsertMany(data, o...)
	} else {
		data := j.Map()
		if fn != nil {
			var err error
			data, err = fn(data)
			if err != nil {
				return nil, err
			}
		}

		id, err = mod.Insert(data, o...)
	}

	if err != nil {
		return nil, err
	}

	return ztype.Map{"id": id}, nil
}

// DeleteById 删除一条数据
func DeleteById(
	c *znet.Context,
	mod *model.Model,
	id string,
	handler func(old ztype.Map) error,
) (any, error) {
	if len(id) == 0 {
		return nil, zerror.InvalidInput.Text("id cannot empty")
	}

	return Delete(c, mod, model.Filter{model.IDKey(): id}, handler)
}

// Delete 删除多条数据
func Delete(
	c *znet.Context,
	mod *model.Model,
	filter model.Filter,
	handler func(old ztype.Map) error,
) (any, error) {
	if handler != nil {
		rows, err := Find(c, mod, filter, nil)
		if err != nil {
			return nil, err
		}

		for i := range rows {
			if err = handler(rows[i]); err != nil {
				return nil, err
			}
		}
	}

	total, err := mod.DeleteMany(filter)
	return ztype.Map{"total": total}, err
}

// UpdateById 更新一条数据
func UpdateById(
	c *znet.Context,
	mod *model.Model,
	id string,
	handler func(old ztype.Map, data ztype.Map) (ztype.Map, error),
) (any, error) {
	if id == "" {
		return nil, zerror.InvalidInput.Text("id cannot empty")
	}

	j, err := c.GetJSONs()
	if err != nil {
		return nil, err
	}

	data := j.Map()

	if handler != nil {
		info, err := mod.FindOneByID(id)
		if err != nil {
			return nil, err
		}

		if info.IsEmpty() {
			return nil, zerror.InvalidInput.Text("id not found")
		}

		data, err = handler(info, data)
		if err != nil {
			return nil, err
		}
	}

	total, err := mod.UpdateByID(id, data)
	return ztype.Map{"total": total}, err
}
