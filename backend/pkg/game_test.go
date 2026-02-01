package pkg_test

import (
	"fmt"
	"testing"

	"github.com/runattekky/footyrate/models"
	"github.com/runattekky/footyrate/pkg"
)

type stubPlayerStore struct {
	players        []*models.Footballer
	mapFootballers map[string]*models.Footballer
	ratingGroup    map[int][]*models.Footballer
}

func (s *stubPlayerStore) GetByID(id string) (*models.Footballer, error) {
	player, ok := s.mapFootballers[id]
	if ok {
		return player, nil
	} else {
		return nil, fmt.Errorf("No player found with id %s", id)
	}
}

func (s *stubPlayerStore) GetAllPlayers() ([]*models.Footballer, error) {
	return s.players, nil
}

func (s *stubPlayerStore) GetRatingGroup() (map[int][]*models.Footballer, error) {
	ratingGroup := make(map[int][]*models.Footballer)

	for _, player := range s.players {
		eloRange := int((player.GetELO() / 100) * 100)
		ratingGroup[eloRange] = append(ratingGroup[eloRange], player)
	}

	return ratingGroup, nil
}

func (s *stubPlayerStore) Save(player *models.Footballer) error {
	if p, ok := s.mapFootballers[player.GetName()]; ok {
		*p = *player
	} else {
		s.players = append(s.players, player)
		s.mapFootballers[player.GetName()] = player
	}
	return nil
}

func TestOpps(t *testing.T) {
	store := &stubPlayerStore{
		players: []*models.Footballer{
			{
				Name: "RunAt",
				ELO:  1500,
			},
			{
				Name: "Cristiano",
				ELO:  2800,
			},
		},
	}
	game := pkg.NewGame(store)

	p1, p2, err := game.GetTwoOpps()
	if err != nil {
		t.Fatalf("Error getting two opponents %v", err)
	}

	assertPlayerNotEqual(t, p1, p2)
}
func assertPlayerNotEqual(t *testing.T, p1, p2 *models.Footballer) {
	t.Helper()
	if p1 == p2 {
		t.Errorf("Expected different players, got same player, %v", p1)
	}
}
