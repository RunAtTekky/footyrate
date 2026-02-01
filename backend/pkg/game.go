package pkg

import (
	"fmt"

	"github.com/runattekky/footyrate/models"
)

type Game struct {
	store models.PlayerStore
}

func (g *Game) GetTwoOpps() (player1, player2 *models.Footballer, err error) {
	player1, err = g.store.GetByID("RunAt")
	if err != nil {
		return nil, nil, fmt.Errorf("Could not find the player with name %s, %v", player1.GetName(), err)
	}
	player2, err = g.store.GetByID("Cristiano")
	if err != nil {
		return nil, nil, fmt.Errorf("Could not find the player with name %s, %v", player2.GetName(), err)
	}

	return player1, player2, err
}

func NewGame(store models.PlayerStore) *Game {
	return &Game{store: store}
}
