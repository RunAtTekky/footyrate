package memory

import (
	"fmt"
	"sync"

	"github.com/runattekky/footyrate/models"
)

func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{
		map[string]*models.Footballer{},
		sync.RWMutex{},
	}
}

type InMemoryPlayerStore struct {
	store map[string]*models.Footballer
	lock  sync.RWMutex
}

func (f *InMemoryPlayerStore) GetByID(id string) (*models.Footballer, error) {
	f.lock.RLock()
	defer f.lock.RUnlock()

	player, ok := f.store[id]
	if !ok {
		return nil, fmt.Errorf("Player not found")
	}

	return player, nil
}

func (f *InMemoryPlayerStore) Save(player *models.Footballer) error {
	f.lock.Lock()
	defer f.lock.Unlock()

	name := player.GetName()
	f.store[name] = player

	return nil
}
