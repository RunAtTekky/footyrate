package models

type Footballer struct {
	Name     string
	ELO      float32
	ImgURL   string
	K_Factor int
	Rounds   int
}

func (f *Footballer) GetELO() float32 {
	return 800.0
}

func (f *Footballer) ChangeELO(change float32) {

}
