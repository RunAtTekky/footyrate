package models

type Footballer struct {
	Name     string  `json:"name"`
	ELO      float32 `json:"elo"`
	ImgURL   string  `json:"imgURL"`
	K_Factor int     `json:"kFactor"`
	Rounds   int     `json:"rounds"`
}

var k_factors = []int{100, 90, 80, 70, 60, 50, 40, 30, 20, 10}

const (
	roundsPerLevel = 10
	InitialELO     = 1000
)

func NewFootballer(name string, ImgURL string) *Footballer {
	footballer := &Footballer{
		Name:     name,
		ELO:      InitialELO,
		ImgURL:   ImgURL,
		Rounds:   0,
		K_Factor: k_factors[0],
	}

	footballer.UpdateKfactor(footballer.Rounds)
	return footballer
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
	f.Rounds++
	f.UpdateKfactor(f.Rounds)
}

func (f *Footballer) UpdateKfactor(rounds int) {
	lvl := rounds / roundsPerLevel
	if lvl >= len(k_factors) {
		lvl = len(k_factors) - 1
	}

	f.K_Factor = k_factors[lvl]
}
