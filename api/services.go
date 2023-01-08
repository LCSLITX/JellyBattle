package api

import (
	"fmt"
	"os"
	"time"

	"github.com/lucassauro/JellyBattle/game"
)

// TODO:  This should be called inside a go routine to check if Playerlist has four players and then creates a game.
// An option to not run this inside an eternal loop, would execute this verification everytime a players is added to availableplayerlist.
func VerifyAvailablePlayersList() {
	if len(game.AVAILABLEPLAYERSLIST) >= 4 {

		_, err := game.AVAILABLEPLAYERSLIST.GroupFourPlayers(game.GROUPS)
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
		fmt.Printf("PlayerList contains %v players.\n", len(game.AVAILABLEPLAYERSLIST))
	}
	time.Sleep(DEFAULT_INTERVAL)
}
