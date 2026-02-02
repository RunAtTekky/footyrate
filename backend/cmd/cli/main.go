package main

import (
	"log"
	"os"

	"github.com/runattekky/footyrate/models"
	"github.com/runattekky/footyrate/pkg"
	"github.com/runattekky/footyrate/pkg/impl"
	filesystem "github.com/runattekky/footyrate/pkg/storage/file-system"
)

const dbFileName = "players.db.json"

var players = []*models.Footballer{
	{
		Name:     "Cristiano",
		ELO:      2800,
		K_Factor: 100,
	},
	{
		Name:     "Messi",
		ELO:      2800,
		K_Factor: 100,
	},
	{
		Name:     "RunAt",
		ELO:      1500,
		K_Factor: 100,
	},
	{
		Name:     "Mbappe",
		ELO:      2700,
		K_Factor: 100,
	},
	{
		Name:     "Modric",
		ELO:      2500,
		K_Factor: 100,
	},
	{
		Name:     "Neymar",
		ELO:      2600,
		K_Factor: 100,
	},
}

func main() {
	store, close, err := filesystem.FileSystemPlayerStoreFromFile(dbFileName)
	if err != nil {
		log.Fatalf("Error creating store from file %s %v", dbFileName, err)
	}
	defer close()

	for _, player := range players {
		store.Save(player)
	}

	game := pkg.NewGame(store)
	cli := impl.NewCLI(os.Stdin, os.Stdout, *game)
	cli.Start()
}
