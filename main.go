package main

import (
	"fmt"
	"os"

	"github.com/lucassauro/JellyBattle/api"
	"github.com/lucassauro/JellyBattle/game"
)

func main() {
	os.Setenv("DEBUG_MODE_STRUCTS", "true")
	os.Setenv("DEBUG_MODE_GROUPS", "true")
	os.Setenv("DEBUG_MODE_GAME", "true")
	os.Setenv("DEBUG_MODE_PLAYER", "true")
	os.Setenv("DEBUG_MODE_PLAYERLIST", "true")
	os.Setenv("DEBUG_MODE_BOARD", "true")
	os.Setenv("DEBUG_MODE_SPECIAL", "true")

	p1 := game.NewPlayer("Lucas")
	game.AVAILABLEPLAYERSLIST.AddPlayer(p1.GetPlayer())

	p2 := game.NewPlayer("P2")
	game.AVAILABLEPLAYERSLIST.AddPlayer(p2.GetPlayer())

	p3 := game.NewPlayer("P3")
	game.AVAILABLEPLAYERSLIST.AddPlayer(p3.GetPlayer())

	p4 := game.NewPlayer("P4")
	game.AVAILABLEPLAYERSLIST.AddPlayer(p4.GetPlayer())

	_, err := game.AVAILABLEPLAYERSLIST.GroupFourPlayers(game.GROUPS.GetGroups())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	game.GROUPS.GetGroups()

	g, err := game.NewGame(game.GROUPS.GetGroups())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	gm := g.GetGame()

	game.GAMES.AddGame(gm) 

	fmt.Println(gm)
	
	api.Server()
}
