package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/lucassauro/JellyBattle/game"
)

func CreatePlayer(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("[%s] %s%s\n", r.Method, r.Host, r.URL)

	if r.Method != "POST" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	// Decode request.Body into req.
	var req game.Player
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// set response headers to json
	w.Header().Set("Content-Type", "application/json")

	// create a player, add it to playerList to wait for games and return its name and Rank (0).
	p := game.NewPlayer(req.Name).GetPlayer()
	game.PLAYERLIST.AddPlayer(p)

	if err := json.NewEncoder(w).Encode(PlayerResponse{
		Name: p.Name,
		Rank: p.Rank,
	}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
