package memory

import "github.com/runattekky/footyrate/models"

type InMemoryPlayerStore struct {
}

func (f *InMemoryPlayerStore) GetByID(id string) (*models.Player, error) {
	return nil, nil
}

func (f *InMemoryPlayerStore) Save(player *models.Player) error {
	return nil
}
