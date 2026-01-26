package models

type Footballer struct {
	Name     string
	ELO      float32
	ImgURL   string
	K_Factor int
	Rounds   int
}

func (f *Footballer) GetName() string {
	return f.Name
}

func (f *Footballer) GetELO() float32 {
	return f.ELO
}

func (f *Footballer) GetImgURL() string {
	return f.ImgURL
}

func (f *Footballer) GetKfactor() int {
	return f.K_Factor
}

func (f *Footballer) GetRounds() int {
	return f.Rounds
}

func (f *Footballer) SetName(name string) {
	f.Name = name
}

func (f *Footballer) SetELO(elo float32) {
	f.ELO = elo
}

func (f *Footballer) SetImgURL(imgURL string) {
	f.ImgURL = imgURL
}

func (f *Footballer) SetKfactor(k_factor int) {
	f.K_Factor = k_factor
}

func (f *Footballer) SetRounds(rounds int) {
	f.Rounds = rounds
}

func (f *Footballer) ChangeELO(change float32) {
	f.ELO += change
}
