package pkg

import "github.com/runattekky/footyrate/models"

func UpdateELO(winner, loser models.Player) {
	expectedScoreWinner := GetExpectedScore(winner.GetELO(), loser.GetELO())
	expectedScoreLoser := 1 - expectedScoreWinner

	ratingChangeWinner := GetRatingChange(1, expectedScoreWinner, winner.GetKfactor())
	ratingChangeLoser := GetRatingChange(0, expectedScoreLoser, loser.GetKfactor())

	winner.ChangeELO(ratingChangeWinner)
	loser.ChangeELO(ratingChangeLoser)
}
