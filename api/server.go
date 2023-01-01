package api

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/lucassauro/JellyBattle/game"
)

func Server() error {
	http.HandleFunc("/", GetRoot)
	http.HandleFunc("/ping", GetPing)

	http.HandleFunc("/player/create", CreatePlayer)

	http.HandleFunc("/playerlist", GetPlayerList)

	http.HandleFunc("/games", GetGames)
	http.HandleFunc("/games/getbyid", GetGameByID)

	http.HandleFunc("/game/start", StartGame)

	os.Setenv("PORT", ":3333")
	fmt.Printf("Server running on localhost%v...\n", os.Getenv("PORT"))

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
			time.Sleep(game.DEFAULT_INTERVAL)
		}
	}()

	if err := http.ListenAndServe(os.Getenv("PORT"), nil); err != nil {
		return err
	}
	return nil
}
