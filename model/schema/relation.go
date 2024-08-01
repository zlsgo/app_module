package schema

type RelationType string

const (
	RelationOne  RelationType = "o2o"
	RelationMany RelationType = "o2m"
)
