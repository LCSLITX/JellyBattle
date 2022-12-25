package main

import (
	"fmt"

	"github.com/lucassauro/J.B.Remake/game/boards"
)

// https://stackoverflow.com/questions/28800672/how-to-add-new-methods-to-an-existing-type-in-go

func main() {
	rows := boards.DEFAULT_ROWS
	columns := boards.DEFAULT_COLUMNS
	players := boards.DEFAULT_PLAYERS_NUMBER

	board := boards.NewBoard(rows, columns, players)

	original := board.GetBoard().Rows
	for r := 0; r < len(original); r++ {
		for c := 0; c < len(original[r]); c++ {
			fmt.Printf("ORIGINAL [SPECIALS]: %v [ROW:] %v \n %+v \n", original[r][c].Fulfilled, original[r][c].Row, original[r][c].Special)
		}
	}
	fmt.Printf("\n")


	board.RandomizeBoard()
	randomized := board.GetBoard().Rows
	for r := 0; r < len(randomized); r++ {
		for c := 0; c < len(randomized[r]); c++ {
			fmt.Printf("RANDOMIZED [SPECIALS]: %v [ROW:] %v \n %+v \n", randomized[r][c].Fulfilled, randomized[r][c].Row, randomized[r][c].Special)
		}
	}
	fmt.Printf("\n")

	board.RoundRows()
	nextRound := board.GetBoard().Rows
	for r := 0; r < len(nextRound); r++ {
		for c := 0; c < len(nextRound[r]); c++ {
			fmt.Printf("NEXT ROUND BOARD: [SPECIALS]: %v [ROW:] %v \n %+v \n", nextRound[r][c].Fulfilled, nextRound[r][c].Row, nextRound[r][c].Special)
		}
	}
	fmt.Printf("\n")
	

	board.UpdateBoard()
	updated := board.GetBoard().Rows
	for r := 0; r < len(updated); r++ {
		for c := 0; c < len(updated[r]); c++ {
			fmt.Printf("UPDATED BOARD: [SPECIALS]: %v [ROW:] %v \n %+v \n", updated[r][c].Fulfilled, updated[r][c].Row, updated[r][c].Special)
		}
	}
	fmt.Printf("\n")
}


// TODO: Create Makefile
// TODO: Create unit and integration tests
