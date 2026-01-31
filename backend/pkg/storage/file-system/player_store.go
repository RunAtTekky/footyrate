package filesystem

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/runattekky/footyrate/models"
)

type file_system_player_store struct {
	database       *json.Encoder
	players        []*models.Footballer
	mapFootballers map[string]*models.Footballer
}

func FileSystemPlayerStoreFromFile(path string) (*file_system_player_store, func(), error) {
	db, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		return nil, nil, fmt.Errorf("Problem opening file %s, %v", path, err)
	}

	closeFunc := func() {
		db.Close()
	}

	store, err := NewFileSystemPlayerStore(db)
	if err != nil {
		return nil, nil, fmt.Errorf("Error creating file system player store %v", err)
	}

	return store, closeFunc, nil

}

func NewFileSystemPlayerStore(database io.ReadWriteSeeker) (*file_system_player_store, error) {
	database.Seek(0, 0)
	players, err := NewPlayers(database)
	if err != nil {
		return nil, fmt.Errorf("Problem loading players %v", err)
	}

	return &file_system_player_store{
		database:       json.NewEncoder(database),
		players:        players,
		mapFootballers: createMap(players),
	}, nil
}

func (f *file_system_player_store) GetByID(id string) (*models.Footballer, error) {
	player, ok := f.mapFootballers[id]
	if !ok {
		return nil, fmt.Errorf("Could not find the player with id %q", id)
	}

	return player, nil
}

func (f *file_system_player_store) Save(player *models.Footballer) error {
	p, err := f.GetByID(player.GetName())
	if err != nil {
		f.players = append(f.players, player)
		return nil
	}

	*p = *player
	return nil
}

func createMap(players []*models.Footballer) map[string]*models.Footballer {
	mapFootballers := make(map[string]*models.Footballer)
	for _, player := range players {
		mapFootballers[player.GetName()] = player
	}

	return mapFootballers
}

func NewPlayers(reader io.ReadWriteSeeker) ([]*models.Footballer, error) {
	var players []*models.Footballer

	_, err := reader.Seek(0, io.SeekStart)
	if err != nil {
		return nil, fmt.Errorf("Problem seeking to the start: %v", err)
	}

	err = json.NewDecoder(reader).Decode(&players)
	if err == io.EOF {
		return players, nil
	}

	if err != nil {
		return nil, fmt.Errorf("Problem parsing players %v", err)
	}

	return players, err
}
