package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/runattekky/footyrate/internal/server"
	filesystem "github.com/runattekky/footyrate/pkg/storage/file-system"
)

const (
	dbFileName = "players.db.json"
	PORT       = 8080
)

func main() {
	store, close, err := filesystem.FileSystemPlayerStoreFromFile(dbFileName)
	if err != nil {
		log.Fatalf("Error creating player store from file %s, %v", dbFileName, err)
	}

	defer close()

	svr := server.NewPlayerServer(store)

	log.Printf("Server starting on %d\n", PORT)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", PORT), svr); err != nil {
		log.Fatalf("Could not listen on PORT %d, %v", PORT, err)
	}
}
