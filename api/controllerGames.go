package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/lucassauro/JellyBattle/game"
)

func GetGames(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("[%s] %s%s\n", r.Method, r.Host, r.URL)

	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	// set response headers to json
	w.Header().Set("Content-Type", "application/json")
	g := game.GAMES.GetGames()

	grs := GamesResponse{}
	for k, v := range g {
		gr := GameResponse{}
		gr.Started = v.Started
		gr.Finished = v.Finished
		gr.ID = v.ID
		gr.Deaths = v.Deaths
		gr.Board = v.Board
		gr.Group = v.Group
		gr.Chat = v.Chat
		grs[k] = gr
	}

	if err := json.NewEncoder(w).Encode(grs); err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func GetGameByID(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("[%s] %s%s\n", r.Method, r.Host, r.URL)

	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	id := r.URL.Query().Get("id")

	var g GameResponse
	for _, v := range game.GAMES.GetGames() {
		if v.ID == id {
			g.Started = v.Started
			g.Finished = v.Finished
			g.ID = v.ID
			g.Deaths = v.Deaths
			g.Board = v.Board
			g.Group = v.Group
			g.Chat = v.Chat
			break
		}
	}

	if g.ID == "" {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(g); err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
