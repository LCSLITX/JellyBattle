package main

import (
	"os"

	"github.com/lucassauro/J.B.Remake/game"
)

// https://stackoverflow.com/questions/28800672/how-to-add-new-methods-to-an-existing-type-in-go

// func main() {
// 	os.Setenv("DEBUG_MODE", "true")
// 	rows := game.DEFAULT_ROWS
// 	columns := game.DEFAULT_COLUMNS
// 	players := game.DEFAULT_PLAYERS_NUMBER
// 	board := game.NewBoard(rows, columns, players)
// 	board.RandomizeBoard()
// 	board.RoundRows()
// 	board.UpdateBoard()
// }

func main() {
	os.Setenv("DEBUG_MODE", "true")
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

	playerList.FindPlayer(game.NewPlayer("B").GetPlayer()) // -1
	playerList.FindPlayer(p1.GetPlayer()) // 0

	playerList.RemovePlayer(p4.GetPlayer())

	playerList.GetPlayerList()
}




// TODO: Create Makefile
// TODO: Create unit and integration tests
