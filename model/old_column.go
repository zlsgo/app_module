package model

type DealOldColumn uint8

const (
	DealOldColumnNone DealOldColumn = iota
	DealOldColumnDelete
	DealOldColumnRename
)
