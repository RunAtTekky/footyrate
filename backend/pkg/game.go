package pkg

import (
	"fmt"
	"math/rand"

	"github.com/runattekky/footyrate/models"
)

type Game struct {
	store models.PlayerStore
}

func (g *Game) GetTwoOpps() (player1, player2 *models.Footballer, err error) {
	allPlayers, err := g.store.GetAllPlayers()
	if err != nil {
		return nil, nil, fmt.Errorf("Could not get all players list %v", err)
	}
	n := len(allPlayers)

	randomIdx := rand.Int() % n
	player1 = allPlayers[randomIdx]

	ratingGroup, err := g.store.GetRatingGroup()
	if err != nil {
		return nil, nil, fmt.Errorf("Could not get rating group %v", err)
	}

	n = len(ratingGroup)
	randomGroupELO := (rand.Int() % 4000) / 100 * 100
	for len(ratingGroup[randomGroupELO]) == 0 {
		randomGroupELO = (rand.Int() % 4000) / 100 * 100
	}

	m := len(ratingGroup[randomGroupELO])
	randomPlayerIdx := rand.Int() % m
	for ratingGroup[randomGroupELO][randomPlayerIdx] != player1 {
		randomPlayerIdx = rand.Int() % m
	}

	player2 = ratingGroup[randomGroupELO][randomPlayerIdx]
	return player1, player2, err
}

func NewGame(store models.PlayerStore) *Game {
	return &Game{store: store}
}
