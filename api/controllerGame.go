package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	
	"github.com/lucassauro/JellyBattle/game"
)

func CreateGame(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("[%s] %s%s\n", r.Method, r.Host, r.URL)

	if r.Method != "POST" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	// create a player, add it to playerList to wait for games and return its name and Rank (0).
	g, err := game.NewGame(game.GROUPS)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createdGame := g.GetGame()

	game.GAMES.AddGame(createdGame)

	// set response headers to json
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(GameResponse{
		Game: createdGame, // TODO: must return only success or failure not game.
	})
}

func StartGame(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("[%s] %s%s\n", r.Method, r.Host, r.URL)

	if r.Method != "GET" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	id := r.URL.Query().Get("id")

	var g game.Game
	for _, v := range game.GAMES.GetGames() {
		if v.ID == id {
			g = v
			break
		}
	}

	if g.ID == "" {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	boolean := g.StartGame()
	if !boolean {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}