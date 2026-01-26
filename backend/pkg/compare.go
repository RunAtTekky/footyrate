package pkg

import "github.com/runattekky/footyrate/models"

func UpdateELO(winner, loser models.Player) {
	// TODO: Write the compare logic
	getELOchange(winner.GetELO(), loser.GetELO())
}
