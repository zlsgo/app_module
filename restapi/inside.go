package restapi

type inside struct {
	oldColumn DealOldColumn
}

type DealOldColumn uint8

const (
	DealOldColumnNone DealOldColumn = iota
	DealOldColumnDelete
	DealOldColumnRename
)

var Inside = &inside{}

func (i *inside) DeleteOldColumn(b DealOldColumn) {
	i.oldColumn = b
}
