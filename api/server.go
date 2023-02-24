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
	fileServer := http.FileServer(http.Dir(DOT_FRONTEND))
	http.Handle(_FRONTEND_, http.StripPrefix(_FRONTEND, fileServer))

	// DISABLED Player
	// http.HandleFunc(_PLAYER_CREATE, CreatePlayer)
	// PlayerList
	http.HandleFunc(_PLAYERLIST, GetAvailablePlayersList)
	http.HandleFunc(_PLAYERLIST_CREATEGROUP, GroupFourPlayers)
	// Games
	http.HandleFunc(_GAMES, GetGames)
	http.HandleFunc(_GAMES_GETBYID, GetGameByID)
	// Groups
	http.HandleFunc(_GROUPS, GetGroups)
	// Game
	http.HandleFunc(_GAME_START, StartGame)
	http.HandleFunc(_GAME_CREATE, CreateGame)

	// WEBSOCKET
	http.HandleFunc(_WS_START_GID, WSGame)
	http.HandleFunc(_WS_PLAYERLIST, WSPlayerList)
	http.HandleFunc(_WS_PLAYER_CREATE, WSPlayerCreate)

	// REST API
	http.HandleFunc(_ROOT, GetHome)
	http.HandleFunc(_PING, GetPing)
	http.HandleFunc(_TEST, Test)

	PORT := os.Getenv("PORT")

	go VerifyAvailablePlayersList()

	// Server
	fmt.Printf("Server starting on localhost%s...\n", PORT)
	if err := http.ListenAndServe(PORT, nil); err != nil {
		fmt.Printf("ERROR serving on port %s: %+v.\n", PORT, err)
	}

	return nil
}
