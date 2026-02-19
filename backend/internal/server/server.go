package server

import (
	"encoding/json"
	"net/http"

	"github.com/runattekky/footyrate/models"
)

type PlayerServer struct {
	store models.PlayerStore
	http.Handler
}

func NewPlayerServer(store models.PlayerStore) *PlayerServer {
	p := new(PlayerServer)
	p.store = store

	router := http.NewServeMux()
	router.Handle("/players", http.HandlerFunc(p.playersHandle))

	p.Handler = router

	return p
}

func (p *PlayerServer) playersHandle(w http.ResponseWriter, r *http.Request) {
	players, _ := p.store.GetAllPlayers()
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(players)
}
