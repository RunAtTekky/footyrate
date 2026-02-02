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

func main() {
	store, close, err := filesystem.FileSystemPlayerStoreFromFile(dbFileName)
	if err != nil {
		log.Fatalf("Error creating store from file %s %v", dbFileName, err)
	}
	defer close()

	// fmt.Println("Lets compare footballers")
	runat := &models.Footballer{
		Name:     "RunAt",
		ELO:      1550,
		K_Factor: 100,
	}
	store.Save(runat)

	messi := &models.Footballer{
		Name:     "Messi",
		ELO:      2800,
		K_Factor: 100,
	}

	store.Save(messi)

	cris := &models.Footballer{
		Name:     "Cristiano",
		ELO:      2800,
		K_Factor: 100,
	}
	store.Save(cris)

	game := pkg.NewGame(store)
	cli := impl.NewCLI(os.Stdin, os.Stdout, *game)
	cli.Start()
}
