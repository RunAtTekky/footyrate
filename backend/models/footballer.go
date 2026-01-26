package models

type Footballer struct {
	Name     string
	ELO      float32
	ImgURL   string
	K_Factor int
	Rounds   int
}

func (f *Footballer) ProcessWin(opponent Player) {

}

func (f *Footballer) ProcessLoss(opponent Player) {

}
