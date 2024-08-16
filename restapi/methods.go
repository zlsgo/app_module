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

// Insert 添加数据
func Insert(
	c *znet.Context,
	mod *model.Model,
	fn func(data ztype.Map) (ztype.Map, error),
	o ...func(io *model.InsertOptions),
) (ztype.Map, error) {
	j, err := c.GetJSONs()
	if err != nil {
		return nil, err
	}

	var data ztype.Map
	if j.IsArray() {
		data = j.Get("0").Map()
	} else {
		data = j.Map()
	}

	if fn != nil {
		var err error
		data, err = fn(data)
		if err != nil {
			return nil, err
		}
	}

	id, err := mod.Insert(data, o...)
	if err != nil {
		return nil, err
	}

	return ztype.Map{"id": id}, nil
}

// InsertMany 添加多条数据
func InsertMany(
	c *znet.Context,
	mod *model.Model,
	fn func(i int, data ztype.Map) (ztype.Map, error),
	o ...func(io *model.InsertOptions),
) (ztype.Map, error) {
	j, err := c.GetJSONs()
	if err != nil {
		return nil, err
	}

	var data ztype.Maps
	if j.IsArray() {
		data = j.Maps()
	} else {
		data = ztype.Maps{j.Map()}
	}

	if fn != nil {
		d := make(ztype.Maps, 0, len(data))
		for i := range data {
			data, err := fn(i, data.Index(i))
			if err != nil {
				return nil, err
			}
			if data.IsEmpty() {
				continue
			}
			d = append(d, data)
		}
		data = d
	}
	id, err := mod.InsertMany(data, o...)
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
