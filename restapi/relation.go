package restapi

import (
	"github.com/sohaha/zlsgo/zarray"
	"github.com/sohaha/zlsgo/znet"
	"github.com/sohaha/zlsgo/ztype"
	"github.com/zlsgo/app_module/model"
)

type Relation struct {
	Operation *model.Store
}

func HanderPageRelation(c *znet.Context, oper *model.Store, filter model.Filter, relations map[string]Relation) (*model.PageData, error) {
	data, err := Page(c, oper, filter, nil)
	return data, err
}

func HanderPageRelation2(data *model.PageData, relations map[string]Relation) (*model.PageData, error) {
	// relations := map[string]any{}

	// size := len(data.Items)
	// relationOper := oper.Operation()
	// relationKey := "attachment"
	// relationName := relationKey + "_relation"
	// relationName = relationKey
	// relationField := "id"
	// relationValues := make([]any, 0, size)

	data.Items = zarray.Map(data.Items, func(_ int, v ztype.Map) ztype.Map {
		// zlog.Debug(v)
		// zlog.Error(v.Get(relationKey).Value())
		// relationValues = append(relationValues, v.Get(relationKey).Value())
		// relationValue := v.Get(relationKey).Value()
		// relationFilter := ztype.Map{
		// 	relationField: relationValue,
		// }
		// relationFields := []string{"id", "path", "size"}
		// v[relationName], _ = relationOper.Find(relationFilter, func(co *model.CondOptions) {
		// 	co.Fields = relationFields
		// })

		return v
	}, 100)
	// relations

	// zlog.Dump(relationValue)
	// rows, err := relationOper.Find(ztype.Map{
	// 	relationField: relationValue,
	// })
	// zlog.Debug(relationValue)
	// zlog.Debug(rows, err)

	return data, nil
}
