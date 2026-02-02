package pkg

import "github.com/runattekky/footyrate/models"

var players = []*models.Footballer{
	{
		Name:     "Cristiano",
		ELO:      2800,
		K_Factor: 100,
	},
	{
		Name:     "Messi",
		ELO:      2800,
		K_Factor: 100,
	},
	{
		Name:     "RunAt",
		ELO:      1500,
		K_Factor: 100,
	},
	{
		Name:     "Mbappe",
		ELO:      2700,
		K_Factor: 100,
	},
	{
		Name:     "Modric",
		ELO:      2500,
		K_Factor: 100,
	},
	{
		Name:     "Neymar",
		ELO:      2600,
		K_Factor: 100,
	},
}

func LoadDummyPlayers(store models.PlayerStore) {
	for _, player := range players {
		store.Save(player)
	}
}
