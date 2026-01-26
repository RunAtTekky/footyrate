package memory

import (
	"fmt"
	"sync"

	"github.com/runattekky/footyrate/models"
)

type InMemoryPlayerStore struct {
	store map[string]*models.Player
	lock  sync.RWMutex
}

func (f *InMemoryPlayerStore) GetByID(id string) (*models.Player, error) {
	f.lock.RLock()
	defer f.lock.Unlock()

	player, ok := f.store[id]
	if !ok {
		return nil, fmt.Errorf("Player not found")
	}

	return player, nil
}

func (f *InMemoryPlayerStore) Save(player *models.Player) error {
	f.lock.RLock()
	defer f.lock.Unlock()

	name := (*player).GetName()
	f.store[name] = player

	return nil
}
