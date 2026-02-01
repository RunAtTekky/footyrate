package main

import (
	"fmt"
	"log"

	"github.com/runattekky/footyrate/pkg"
	filesystem "github.com/runattekky/footyrate/pkg/storage/file-system"
)

const dbFileName = "players.db.json"

func main() {
	store, close, err := filesystem.FileSystemPlayerStoreFromFile(dbFileName)
	if err != nil {
		log.Fatalf("Error creating store from file %s %v", dbFileName, err)
	}
	defer close()

	fmt.Println("Lets compare footballers")
	// runat := &models.Footballer{Name: "RunAt"}
	// store.Save(runat)
	//
	// cris := &models.Footballer{Name: "Cristiano"}
	// store.Save(cris)
	//
	game := pkg.NewGame(store)
	p1, p2, err := game.GetTwoOpps()
	if err != nil {
		log.Fatalf("Could not get two opps %v", err)
	}

	fmt.Printf("p1:\n%v\np2:\n%v", p1, p2)
}
