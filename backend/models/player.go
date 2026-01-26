package models

type Player interface {
	GetName() string
	GetELO() float32
	GetImgURL() string
	GetKfactor() int
	GetRounds() int

	SetName(string)
	SetELO(float32)
	SetImgURL(string)
	SetKfactor(int)
	SetRounds(int)

	ChangeELO(change float32)
	UpdateKfactor(rounds int)
}
