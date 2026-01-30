package filesystem_test

import (
	"testing"

	"github.com/runattekky/footyrate/models"
	filesystem "github.com/runattekky/footyrate/pkg/storage/file-system"
)

const dbFileName = "players.db"

func TestFileStore(t *testing.T) {
	t.Run("", func(t *testing.T) {
		store, closeFunc, err := filesystem.FileSystemPlayerStoreFromFile(dbFileName)
		if err != nil {
			t.Errorf("Error creating Player Store %v", err)
		}
		defer closeFunc()

		messi := models.NewFootballer("Messi", "")
		store.Save(messi)
	})
}
