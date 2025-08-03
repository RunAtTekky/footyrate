package handlers

import (
	"encoding/json"
	"fmt"
	"image_compare/models"
	"math"
	"net/http"
)

func Handle_result(w http.ResponseWriter, r *http.Request) {

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	w.WriteHeader(http.StatusOK)

	result := new(models.Result)

	if err := json.NewDecoder(r.Body).Decode(result); err != nil {
		fmt.Printf("Error decoding result")
		return
	}

	winner, loser := compute_result(result)

	response := models.Response{
		Player1: winner,
		Player2: loser,
	}

	json.NewEncoder(w).Encode(response)
}

func compute_result(result *models.Result) (models.Player, models.Player) {

	var winner_idx int = -1
	var loser_idx int = -1

	for idx := range models.All_Players.Player_List {
		if models.All_Players.Player_List[idx].ID == result.Winner_ID {
			winner_idx = idx
		}
		if models.All_Players.Player_List[idx].ID == result.Loser_ID {
			loser_idx = idx
		}

		if winner_idx != -1 && loser_idx != -1 {
			break
		}
	}

	winner := &models.All_Players.Player_List[winner_idx]
	loser := &models.All_Players.Player_List[loser_idx]

	update_ELO(winner, loser)

	models.All_Players.Update_ELO(winner)
	models.All_Players.Update_ELO(loser)

	models.All_Players.Update_K_Factor(winner)
	models.All_Players.Update_K_Factor(loser)

	models.All_Players.Update_Rounds(winner)
	models.All_Players.Update_Rounds(loser)

	return *winner, *loser
}

// func update_ELO(winner *models.Player, loser *models.Player) {
// 	expected_winner_score := 1 / (math.Pow(10, float64(float32(loser.ELO)-float32(winner.ELO))/400) + 1)
//
// 	elo_change_winner := float64(winner.K_FACTOR) * (1 - expected_winner_score)
//
// 	expected_loser_score := 1 - expected_winner_score
//
// 	elo_change_loser := float64(loser.K_FACTOR) * (0 - expected_loser_score)
//
// 	winner.ELO += int(math.Round(elo_change_winner))
//
// 	new_loser_elo := loser.ELO + int(math.Round(elo_change_loser))
//
// 	loser.ELO = max(0, new_loser_elo)
//
// 	winner.ROUNDS += 1
// 	loser.ROUNDS += 1
//
// 	update_K_factor(winner)
// 	update_K_factor(loser)
// }

func update_ELO(winner *models.Player, loser *models.Player) {
	var difference_ELO float32 = float32(loser.ELO) - float32(winner.ELO)

	expected_A := 1 / (math.Pow(10, float64(difference_ELO/400)) + 1)
	expected_B := 1 / (math.Pow(10, float64(-1*difference_ELO/400)) + 1)

	ELO_change_winner := float64(winner.K_FACTOR) * (1 - expected_A)
	ELO_change_loser := float64(loser.K_FACTOR) * (0 - expected_B)

	winner.ELO += int(ELO_change_winner)
	loser.ELO += int(ELO_change_loser)

	winner.ROUNDS += 1
	loser.ROUNDS += 1

	update_K_factor(winner)
	update_K_factor(loser)
}

func update_K_factor(player *models.Player) {
	switch {
	case player.ROUNDS > 30:
		player.K_FACTOR = 10
	case player.ROUNDS > 20:
		player.K_FACTOR = 20
	case player.ROUNDS > 10:
		player.K_FACTOR = 30
	}
}
