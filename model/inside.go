package model

type inside struct {
	oldColumn DealOldColumn
}

type DealOldColumn uint8

const (
	dealOldColumnNone DealOldColumn = iota
	dealOldColumnDelete
	dealOldColumnRename
)

var InsideOption = &inside{}

func (i *inside) OldColumnIgnore() {
	i.oldColumn = dealOldColumnNone
}

func (i *inside) OldColumnDelete() {
	i.oldColumn = dealOldColumnDelete
}

func (i *inside) OldColumnRename() {
	i.oldColumn = dealOldColumnRename
}
