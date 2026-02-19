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

type compareRequest struct {
	Winner string `json:"winner"`
	Loser  string `json:"loser"`
}

func NewPlayerServer(store models.PlayerStore) *PlayerServer {
	p := new(PlayerServer)
	p.game = pkg.NewGame(store)

	router := http.NewServeMux()
	router.HandleFunc("/players", p.getPlayerList)
	router.HandleFunc("GET /compare", p.getTwoPlayers)
	router.HandleFunc("POST /compare", p.saveVote)
	router.HandleFunc("POST /add", p.addPlayer)

	p.Handler = router
	return p
}

func (p *PlayerServer) getPlayerList(w http.ResponseWriter, r *http.Request) {
	players, _ := p.game.Store.GetAllPlayers()
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(players)
}

func (p *PlayerServer) getTwoPlayers(w http.ResponseWriter, r *http.Request) {
	p1, p2, _ := p.game.GetTwoOpps()
	response := compareResponse{
		PlayerOne: p1,
		PlayerTwo: p2,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (p *PlayerServer) saveVote(w http.ResponseWriter, r *http.Request) {
	var req compareRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	winner, _ := p.game.Store.GetByID(req.Winner)
	loser, _ := p.game.Store.GetByID(req.Loser)
	p.game.SaveChoice(winner, loser)

	response := map[string]string{"status": "recorded"}
	json.NewEncoder(w).Encode(response)

}

func (p *PlayerServer) addPlayer(w http.ResponseWriter, r *http.Request) {
	var player *models.Footballer
	err := json.NewDecoder(r.Body).Decode(&player)
	if err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	p.game.Store.Save(player)
	w.Header().Set("Content-Type", "application/json")
	response := map[string]string{"status": "added"}
	json.NewEncoder(w).Encode(response)
}
