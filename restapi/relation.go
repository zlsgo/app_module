package restapi

import (
	"github.com/sohaha/zlsgo/znet"
	"github.com/zlsgo/app_module/model"
)

// Relation 描述一个关联关系：operation 为关联目标模型。
//
// Deprecated: 关联数据请通过查询参数 with/relations 配合 Options.AllowRelations
// 装载，restapi 会在路由内部将关联写入 CondOptions.Relations。
type Relation struct {
	Operation *model.Store
}

// HanderPageRelation 基于请求上下文对指定 Store 执行分页查询并返回结果。
//
// 说明：该 helper 仅负责分页，不做关联装载；需装载关联时请走通配路由
// （with/relations 参数），关联逻辑由路由内部实现。
func HanderPageRelation(c *znet.Context, oper *model.Store, filter model.Filter, _ map[string]Relation) (*model.PageData, error) {
	return Page(c, oper, filter, nil)
}
