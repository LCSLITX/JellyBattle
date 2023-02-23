package api

import (
	"fmt"
	"net/http"
	"os"
)

func Server() error {
	// Check it out to know more: https://forum.golangbridge.org/t/cant-get-a-background-image-to-display-using-go-webserver/6940
	// Create a file server which serves static files out of the "./frontend" directory.
	// Use the http.Handle() function to register the file server as the handler for all URL paths that start with "/frontend/". For matching
	// paths, we strip the "/frontend" prefix before the request reaches the file server.
	fileServer := http.FileServer(http.Dir("./frontend"))
	http.Handle("/frontend/", http.StripPrefix("/frontend", fileServer))

	// Player
	http.HandleFunc("/player/create", CreatePlayer)
	// PlayerList
	http.HandleFunc("/playerlist", GetAvailablePlayersList)
	http.HandleFunc("/playerlist/creategroup", GroupFourPlayers)
	// Games
	http.HandleFunc("/games", GetGames)
	http.HandleFunc("/games/getbyid", GetGameByID)
	// Groups
	http.HandleFunc("/groups", GetGroups)
	// Game
	http.HandleFunc("/game/start", StartGame)
	http.HandleFunc("/game/create", CreateGame)
	// REST API
	http.HandleFunc("/", GetHome)
	http.HandleFunc("/ping", GetPing)
	http.HandleFunc("/test", Test)
	
	// WEBSOCKET
	http.HandleFunc("/ws/start", WSGame)
	http.HandleFunc("/ws/playerlist", WSPlayerList)

	PORT := os.Getenv("PORT")

	go VerifyAvailablePlayersList()

	// Server
	fmt.Printf("Server starting on localhost%s...\n", PORT)
	if err := http.ListenAndServe(PORT, nil); err != nil {
		fmt.Printf("ERROR serving on port %s: %+v.\n", PORT, err)
	}

	return nil
}
