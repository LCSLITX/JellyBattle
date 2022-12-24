package main

import (
	"fmt"

	"github.com/lucassauro/J.B.Remake/game/boards"
)

func main() {
	rows := uint8(2)
	columns := uint8(10)
	players := uint8(2)

	board := boards.NewBoard(rows, columns, players)
	board.RandomizeBoard()
	fmt.Printf("[RANDOMIZED BOARD]:\n%+v\n\n", board.GetBoard().Rows)

	board.RoundRows()
	fmt.Printf("[ROUNDED BOARD]:\n%+v\n\n", board.GetBoard().Rows)

}


// TODO: Create Makefile
// TODO: Create unit and integration tests
