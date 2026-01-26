package memory

import (
	"testing"

	"github.com/runattekky/footyrate/models"
)

type stubPlayerStore struct {
	Players []models.Player
}

func (s *stubPlayerStore) GetByID(id string) (*models.Player, error) {
	return nil, nil
}

func (s *stubPlayerStore) Save(player *models.Footballer) error {
	s.Players = append(s.Players, player)
	return nil
}

func TestInMemoryPlayerStore(t *testing.T) {
	t.Run("Save RunAt and Minato to PlayerStore", func(t *testing.T) {
		runAt := models.NewFootballer("RunAt", "")
		minato := models.NewFootballer("Minato", "")

		store := &stubPlayerStore{}
		store.Save(runAt)
		store.Save(minato)

		if (len(store.Players) < 2) || (store.Players[0] != runAt && store.Players[1] != minato) {
			t.Fatalf("Wanted %v and %v\nGot %v and %v", runAt, minato, store.Players[0], store.Players[1])
		}
	})
}
