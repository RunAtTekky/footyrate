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

	compute_result(result)

	json.NewEncoder(w).Encode(result)
}

func compute_result(result *models.Result) {

	winner := &All_Players.Images[result.Winner_ID]
	loser := &All_Players.Images[result.Loser_ID]

	update_ELO(winner, loser)

	All_Players.Update_ELO(winner)
	All_Players.Update_ELO(loser)

	All_Players.Update_Rounds(winner)
	All_Players.Update_Rounds(loser)
}

func update_ELO(winner *models.Image, loser *models.Image) {
	// TODO - Write logic for updating elo

	var difference_ELO float32 = float32(winner.ELO) - float32(loser.ELO)

	expected := 1 / (math.Pow(10, float64(difference_ELO/400)) + 1)

	ELO_change_winner := float64(winner.K_FACTOR) * (1 - expected)
	ELO_change_loser := float64(loser.K_FACTOR) * (1 - expected)

	winner.ELO += int(ELO_change_winner)
	loser.ELO -= int(ELO_change_loser)

	winner.ROUNDS += 1
	loser.ROUNDS += 1

	update_K_factor(winner)
	update_K_factor(loser)
}

func update_K_factor(player *models.Image) {
	switch {
	case player.ROUNDS > 30:
		player.K_FACTOR = 10
	case player.ROUNDS > 20:
		player.K_FACTOR = 20
	case player.ROUNDS > 10:
		player.K_FACTOR = 30
	}
}
