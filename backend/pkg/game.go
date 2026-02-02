package pkg

import (
	"fmt"
	"math/rand"

	"github.com/runattekky/footyrate/models"
)

type Game struct {
	store models.PlayerStore
}

func getRandomPlayer(players []*models.Footballer) *models.Footballer {
	n := len(players)
	randomIdx := rand.Int() % n
	return players[randomIdx]
}

func getRandomPlayerFromGroup(ratingGroup map[int][]*models.Footballer) *models.Footballer {
	randomGroupELO := (rand.Int() % 4000) / 100 * 100
	for len(ratingGroup[randomGroupELO]) == 0 {
		randomGroupELO = (rand.Int() % 4000) / 100 * 100
	}

	return getRandomPlayer(ratingGroup[randomGroupELO])
}

func (g *Game) GetTwoOpps() (player1, player2 *models.Footballer, err error) {
	allPlayers, err := g.store.GetAllPlayers()
	if err != nil {
		return nil, nil, fmt.Errorf("Could not get all players list %v", err)
	}
	player1 = getRandomPlayer(allPlayers)

	ratingGroup, err := g.store.GetRatingGroup()
	if err != nil {
		return nil, nil, fmt.Errorf("Could not get rating group %v", err)
	}

	for player2 = getRandomPlayerFromGroup(ratingGroup); player1.GetName() == player2.GetName(); {
		player2 = getRandomPlayerFromGroup(ratingGroup)
	}

	return player1, player2, nil
}

func NewGame(store models.PlayerStore) *Game {
	return &Game{store: store}
}
