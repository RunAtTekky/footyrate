package models

type Footballer struct {
	Name     string
	ELO      float32
	ImgURL   string
	K_Factor int
	Rounds   int
}

func (f *Footballer) GetELO() float32 {
	return f.ELO
}

func (f *Footballer) ChangeELO(change float32) {
	f.ELO += change
}

func (f *Footballer) GetKfactor() int {
	return f.K_Factor
}
