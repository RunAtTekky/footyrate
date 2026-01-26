package models

type Player interface {
	GetELO() float32
	ChangeELO(change float32)
}
