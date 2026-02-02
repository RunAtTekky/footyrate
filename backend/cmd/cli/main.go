package main

import (
	"log"
	"os"

	"github.com/runattekky/footyrate/pkg"
	"github.com/runattekky/footyrate/pkg/impl"
	filesystem "github.com/runattekky/footyrate/pkg/storage/file-system"
)

const dbFileName = "players.db.json"

func main() {
	store, close, err := filesystem.FileSystemPlayerStoreFromFile(dbFileName)
	if err != nil {
		log.Fatalf("Error creating store from file %s %v", dbFileName, err)
	}
	defer close()

	pkg.LoadDummyPlayers(store)

	game := pkg.NewGame(store)
	cli := impl.NewCLI(os.Stdin, os.Stdout, *game)
	cli.Start()
}
