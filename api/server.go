package api

import (
	"fmt"
	"net/http"
	"os"
)

func Server() error {
	// Chcek it out to know more: https://forum.golangbridge.org/t/cant-get-a-background-image-to-display-using-go-webserver/6940
	// Create a file server which serves static files out of the "./images" directory.
	// Use the http.Handle() function to register the file server as the handler for all URL paths that start with "/images/". For matching
	// paths, we strip the "/image" prefix before the request reaches the file server.
	fileServer := http.FileServer(http.Dir("./images"))
	http.Handle("/images/", http.StripPrefix("/images", fileServer))

	// REST API
	http.HandleFunc("/home", GetHome)
	http.HandleFunc("/ping", GetPing)
	http.HandleFunc("/test", Test)
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

	// WEBSOCKET
	http.HandleFunc("/ws/start", WSGame)
	http.HandleFunc("/ws/playerlist", WSPlayerList)

	PORT := os.Getenv("PORT")

	// Server
	fmt.Printf("Server starting on localhost%s...\n", PORT)
	if err := http.ListenAndServe(PORT, nil); err != nil {
		fmt.Printf("ERROR serving on port %s: %+v.\n", PORT, err)
	}

	go VerifyAvailablePlayersList()

	return nil
}
