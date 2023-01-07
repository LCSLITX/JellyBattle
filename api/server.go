package api

import (
	"fmt"
	"net/http"
	"os"
)

func Server() error {
	os.Setenv("DEBUG_MODE_WEBSOCKET", "true")

	// Chcek it out to know more: https://forum.golangbridge.org/t/cant-get-a-background-image-to-display-using-go-webserver/6940
	// Create a file server which serves static files out of the "./images" directory.
	// Use the http.Handle() function to register the file server as the handler for all URL paths that start with "/images/". For matching
	// paths, we strip the "/image" prefix before the request reaches the file server.
	fileServer := http.FileServer(http.Dir("./images"))
	http.Handle("/images/", http.StripPrefix("/images", fileServer))

	// REST API
	http.HandleFunc("/", GetHome)
	http.HandleFunc("/ping", GetPing)
	// Player
	http.HandleFunc("/player/create", CreatePlayer)
	// PlayerList
	http.HandleFunc("/playerlist", GetPlayerList)
	// Games
	http.HandleFunc("/games", GetGames)
	http.HandleFunc("/games/getbyid", GetGameByID)
	// Game
	http.HandleFunc("/game/start", StartGame)

	// WEBSOCKET
	http.HandleFunc("/ws/start", WSGame)
	
	PORT := os.Getenv("PORT")
	
	// Server
	fmt.Printf("Server starting on localhost%s...\n", PORT)
	if err := http.ListenAndServe(PORT, nil); err != nil {
		fmt.Printf("ERROR serving on port %s: %+v.\n", PORT, err)
	}

	return nil
}


// TODO:  Take this horrible thing out of here.
// This go routine checks if Playerlist has four players and then creates a game.
// go func() {
// 	for {
// 		if len(game.PLAYERLIST) >= 4 {

// 			_, err := game.PLAYERLIST.GroupFourPlayers(game.GROUPS)
// 			if err != nil {
// 				fmt.Println(fmt.Errorf(err.Error())) // horrible
// 				os.Exit(1)
// 			}
// 			fmt.Printf("Created a group.\n")

// 			createdGame, err := game.NewGame(game.GROUPS)
// 			if err != nil {
// 				fmt.Println(fmt.Errorf(err.Error())) // horrible
// 				os.Exit(1)
// 			}

// 			fmt.Printf("Created a game: %+v.\n\n", createdGame.GetGame())
// 			game.GAMES.AddGame(createdGame.GetGame())

// 			fmt.Printf("Added a game to games: %+v.\n", game.GAMES.GetGames())

// 		} else {
// 			fmt.Printf("PlayerList contains %v players.\n", len(game.PLAYERLIST))
// 		}
// 		time.Sleep(DEFAULT_INTERVAL)
// 	}
// }()