package server

import (
	"encoding/json"
	"net/http"

	"github.com/runattekky/footyrate/models"
	"github.com/runattekky/footyrate/pkg"
)

type PlayerServer struct {
	game *pkg.Game
	http.Handler
}

type compareResponse struct {
	PlayerOne *models.Footballer `json:"p1"`
	PlayerTwo *models.Footballer `json:"p2"`
}

func NewPlayerServer(store models.PlayerStore) *PlayerServer {
	p := new(PlayerServer)
	p.game = pkg.NewGame(store)

	router := http.NewServeMux()
	router.Handle("/players", http.HandlerFunc(p.playersHandle))
	router.Handle("/compare", http.HandlerFunc(p.compareHandle))

	p.Handler = router
	return p
}

func (p *PlayerServer) playersHandle(w http.ResponseWriter, r *http.Request) {
	players, _ := p.game.Store.GetAllPlayers()
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(players)
}

func (p *PlayerServer) compareHandle(w http.ResponseWriter, r *http.Request) {
	p1, p2, _ := p.game.GetTwoOpps()
	response := compareResponse{
		PlayerOne: p1,
		PlayerTwo: p2,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
