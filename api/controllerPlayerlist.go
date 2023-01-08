package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/lucassauro/JellyBattle/game"
)

func GetPlayerList(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("[%s] %s%s\n", r.Method, r.Host, r.URL)

	if r.Method != "GET" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	// set response headers to json
	w.Header().Set("Content-Type", "application/json")

	pl := game.AVAILABLEPLAYERSLIST.GetPlayerList()

	list := make([]PlayerResponse, 0)

	for _, v := range pl {
		var p PlayerResponse
		p.Name, p.Rank = v.Name, v.Rank
		list = append(list, p)
	}

	if err := json.NewEncoder(w).Encode(PlayerListResponse{
		List: list,
	}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
