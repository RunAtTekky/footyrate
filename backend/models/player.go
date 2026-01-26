package models

type Player interface {
	GetELO() float32
	GetKfactor() int
	ChangeELO(change float32)
}
