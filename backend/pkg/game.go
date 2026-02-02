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
	// TODO: Only give players with rating difference of +-200
	n := rand.Int() % len(ratingGroup)
	for _, playersWithRating := range ratingGroup {
		if n == 0 {
			return getRandomPlayer(playersWithRating)
		}
		n--
	}

	return nil
}

func (g *Game) GetTwoOpps() (player1, player2 *models.Footballer, err error) {
	allPlayers, err := g.store.GetAllPlayers()
	if err != nil {
		return nil, nil, fmt.Errorf("Could not get all players list %v", err)
	}
	if len(allPlayers) == 0 {
		return nil, nil, fmt.Errorf("No player is there, can not give two opps")
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

func (g *Game) SaveChoice(winner, loser *models.Footballer) {
	UpdateELO(winner, loser)
	g.store.Save(winner)
	g.store.Save(loser)
}

func NewGame(store models.PlayerStore) *Game {
	return &Game{store: store}
}
