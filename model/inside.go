package model

type inside struct {
	oldColumn        DealOldColumn
	softDeleteIsNull bool
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

func (i *inside) SoftDeleteIsNull(b ...bool) {
	softDeleteIsNull := true

	if len(b) > 0 {
		softDeleteIsNull = b[0]
	}
	i.softDeleteIsNull = softDeleteIsNull
}
