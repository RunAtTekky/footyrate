package filesystem_test

import (
	"fmt"
	"io"
	"strings"
	"testing"

	"github.com/runattekky/footyrate/models"
	filesystem "github.com/runattekky/footyrate/pkg/storage/file-system"
)

type stubReadWriteSeeker struct {
	io.ReadSeeker
}

func (s *stubReadWriteSeeker) Write(p []byte) (n int, err error) {
	return 0, nil
}

type StubFileSystemPlayerStore struct {
	players        []*models.Footballer
	mapFootballers map[string]*models.Footballer
}

func (s *StubFileSystemPlayerStore) GetByID(id string) (*models.Footballer, error) {
	player, ok := s.mapFootballers[id]
	if !ok {
		return nil, fmt.Errorf("Could not find the player with id %q", id)
	}

	return player, nil
}

func (s *StubFileSystemPlayerStore) Save(player *models.Footballer) error {
	p, err := s.GetByID(player.GetName())
	if err != nil {
		s.players = append(s.players, player)
		return nil
	}

	*p = *player
	return nil
}

func TestFileStore(t *testing.T) {
	t.Run("Testing FileSystem Parsing and Saving", func(t *testing.T) {
		data := `[{"name": "Messi"}]`
		database := &stubReadWriteSeeker{strings.NewReader(data)}
		store, err := filesystem.NewFileSystemPlayerStore(database)
		if err != nil {
			t.Fatalf("Failed to create store %v", err)
		}

		messi := models.NewFootballer("RunAt", "")
		store.Save(messi)

		got, err := store.GetByID("Messi")
		if err != nil {
			t.Errorf("Expected to find Messi, but got error %v", err)
		}

		want := "Messi"

		if got.GetName() != want {
			t.Errorf("Got %s but want %s", got.GetName(), want)
		}
	})
}
