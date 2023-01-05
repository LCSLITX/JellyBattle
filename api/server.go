package api

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/lucassauro/JellyBattle/game"
)

func init() {
	os.Setenv("PORT", ":3333")
}

func Server() error {
	// REST API
	http.HandleFunc("/", GetRoot)
	http.HandleFunc("/ping", GetPing)
	// Player
	http.HandleFunc("/player/create", CreatePlayer)
	// PlayerList
	http.HandleFunc("/playerlist", GetPlayerList)
	// Games
	http.HandleFunc("/games", GetGames)
	http.HandleFunc("/games/getbyid", GetGameByID)
	//Game
	http.HandleFunc("/game/start", StartGame)


	// WEBSOCKET

	
	// Server
	fmt.Printf("Server running on localhost%v...\n", PORT)


	// TODO: Take this horrible thing out of here.
	// This go routine checks if Playerlist has four players and then creates a game.
	go func() {
		for {
			if len(game.PLAYERLIST) >= 4 {

				_, err := game.PLAYERLIST.GroupFourPlayers(game.GROUPS)
				if err != nil {
					fmt.Println(fmt.Errorf(err.Error())) // horrible
					os.Exit(1)
				}
				fmt.Printf("Created a group.\n")

				createdGame, err := game.NewGame(game.GROUPS)
				if err != nil {
					fmt.Println(fmt.Errorf(err.Error())) // horrible
					os.Exit(1)
				}

				fmt.Printf("Created a game: %+v.\n\n", createdGame.GetGame())
				game.GAMES.AddGame(createdGame.GetGame())

				fmt.Printf("Added a game to games: %+v.\n", game.GAMES.GetGames())

			} else {
				fmt.Printf("PlayerList contains %v players.\n", len(game.PLAYERLIST))
			}
			time.Sleep(DEFAULT_INTERVAL)
		}
	}()

	if err := http.ListenAndServe(PORT, nil); err != nil {
		return err
	}
	return nil
}
