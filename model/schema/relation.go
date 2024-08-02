package schema

type RelationType string

const (
	// RelationOne 单对单
	RelationOne RelationType = "o2o"
	// RelationOneMerge 单对单，结果合并
	RelationOneMerge RelationType = "o2o_merge"
	// RelationMany 单对多
	RelationMany RelationType = "o2m"
)
