package filesystem

import (
	"encoding/json"
	"fmt"

	"github.com/runattekky/footyrate/models"
)

type file_system_player_store struct {
	database *json.Encoder
	players  []models.Footballer
	store    map[string]*models.Footballer
}

func (f *file_system_player_store) GetByID(id string) (*models.Footballer, error) {
	player, ok := f.store[id]
	if !ok {
		return nil, fmt.Errorf("Could not find the player with id %q", id)
	}

	return player, nil
}

func (f *file_system_player_store) Save(player *models.Footballer) error {
	_, err := f.GetByID(player.GetName())
	if err != nil {
		f.players = append(f.players, *player)
		return nil
	}

	return nil
}
