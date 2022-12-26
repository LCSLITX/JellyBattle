package main

import (
	"fmt"
	"os"

	"github.com/lucassauro/J.B.Remake/game"
)

// https://stackoverflow.com/questions/28800672/how-to-add-new-methods-to-an-existing-type-in-go


func main() {
	// os.Setenv("DEBUG_MODE_BOARD", "true")
	os.Setenv("DEBUG_MODE_GROUPS", "true")
	os.Setenv("DEBUG_MODE_GAME", "true")
	os.Setenv("DEBUG_MODE_PLAYER", "true")
	os.Setenv("DEBUG_MODE_PLAYERLIST", "true")
	// os.Setenv("DEBUG_MODE_SPECIAL", "true")

	rows := game.DEFAULT_ROWS
	columns := game.DEFAULT_COLUMNS
	players := game.DEFAULT_PLAYERS_NUMBER

	board := game.NewBoard(rows, columns, players)
	board.RandomizeBoard()

	p1 := game.NewPlayer("Lucas")
	p2 := game.NewPlayer("P2")
	p3 := game.NewPlayer("P3")
	p4 := game.NewPlayer("P4")

	playerList := game.NewPlayerList()
	playerList.AddPlayer(p1.GetPlayer())
	playerList.GetPlayerList()
	playerList.AddPlayer(p2.GetPlayer())
	playerList.GetPlayerList()
	playerList.AddPlayer(p3.GetPlayer())
	playerList.GetPlayerList()
	playerList.AddPlayer(p4.GetPlayer())
	playerList.GetPlayerList()

	// playerList.FindPlayer(game.NewPlayer("B").GetPlayer()) // -1
	// playerList.FindPlayer(p1.GetPlayer())                  // 0
	// playerList.GetPlayerList()

	groups := game.NewGroups().GetGroups()

	group, err := playerList.GroupFourPlayers(&groups)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	groups.GetGroups()
	
	playerList.GetPlayerList()

	// groups.RemoveGroup(group)

	// groups.GetGroups()

	game.NewGame(board.GetBoard(), group)

}

// TODO: Create Makefile
// TODO: Create unit and integration tests
