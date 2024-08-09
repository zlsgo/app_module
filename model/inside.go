package model

type inside struct {
	oldColumn        DealOldColumn
	softDeleteIsTime bool
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

func (i *inside) SoftDeleteIsTime(b ...bool) {
	is := true

	if len(b) > 0 {
		is = b[0]
	}
	i.softDeleteIsTime = is
}
