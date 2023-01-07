package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/lucassauro/JellyBattle/game"
)

func GetGames(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("[%s] %s%s\n", r.Method, r.Host, r.URL)

	if r.Method != "GET" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	
	// set response headers to json
	w.Header().Set("Content-Type", "application/json")
	g := game.GAMES.GetGames()
	
	if err := json.NewEncoder(w).Encode(GamesResponse{
		List: g,
	}); err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func GetGameByID(w http.ResponseWriter, r *http.Request) {
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

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(GameResponse{
		Game: g,
	}); err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}