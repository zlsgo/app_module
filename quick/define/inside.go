package define

type inside struct {
	oldColumn         MigrationOldColumn
	deleteFieldPrefix string // 删除字段前缀
	idKey             string
	createdAtKey      string
	updatedAtKey      string
	deletedAtKey      string
}

type MigrationOldColumn uint8

const (
	DealOldColumnNone MigrationOldColumn = iota
	DealOldColumnDelete
	DealOldColumnRename
)

var Inside = &inside{
	deleteFieldPrefix: "__del__",
	idKey:             "_id",
	createdAtKey:      "created_at",
	updatedAtKey:      "updated_at",
	deletedAtKey:      "deleted_at",
}

func (i *inside) IDKey() string {
	return i.idKey
}

func (i *inside) CreatedAtKey() string {
	return i.createdAtKey
}

func (i *inside) UpdatedAtKey() string {
	return i.updatedAtKey
}

func (i *inside) DeletedAtKey() string {
	return i.deletedAtKey
}

func (i *inside) DeleteFieldPrefix() string {
	return i.deleteFieldPrefix
}
