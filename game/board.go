package game

import (
	"fmt"
)

// https://stackoverflow.com/questions/28800672/how-to-add-new-methods-to-an-existing-type-in-go

// NewBoard returns an empty Board instance.
func NewBoard(rows, columns, players uint8) IBoard {
	b := &Board{}
	b.GenerateBoard(rows, columns)
	b.PlayersNumber = players
	b.SpecialFulfillment = DEFAULT_SPECIAL_FULFILLMENT

	return b
}

// GetBoard returns Board object.
func (board *Board) GetBoard() Board {
	if DebugModeBoard() {
		fmt.Printf("%v: %+v\n\n", Trace(), *board)
	}
	return *board
}

// Generateboard generates a board with specified number of rows and columns
func (board *Board) GenerateBoard(rows, columns uint8) {
	board.RowNumber, board.ColumnNumber = rows, columns
	q := uint64(0)
	for r := uint8(0); r < rows; r++ {
		row := Row{} // Create an empty row of buttons
		for c := uint8(0); c < columns; c++ {
			b := Button{Row: r, Column: c} // Create an empty button
			row = append(row, b)           // Add new button to created row of buttons
			q += 1
		}
		board.Rows = append(board.Rows, row) // Add new row of buttons into Rows array.
	}
	board.NumberOfButtons = q
	if DebugModeBoard() {
		fmt.Printf("%v: %+v\n\n", Trace(), board)
	}
}

// GeneratePreviewRow() generates one row at once.
func (board *Board) GeneratePreviewRow() {
	rowsLength, columns := uint8(len(board.Rows)), board.ColumnNumber
	var row Row
	for c := uint8(0); c < columns; c++ {
		b := Button{Row: rowsLength, Column: c} // Create an empty button
		b.RandomizeButton()                     // Randomize it
		row = append(row, b)                    // Add new button to created row of buttons
	}
	if DebugModeBoard() {
		fmt.Printf("%v: %+v\n\n", Trace(), row)
	}
	board.PreviewRow = row
}

// RoundRows rotates the board rows and updates the buttons to correspond correctly to its respective Rows and columns, supposed to be used when changing rounds.
func (board *Board) RoundRows() {
	board.Rows = board.Rows[1:]                       // remove top element index[1]
	board.Rows = append(board.Rows, board.PreviewRow) // add last element index[length - 1]

	rows, columns := board.RowNumber, board.ColumnNumber
	for r := uint8(0); r < rows; r++ {
		for c := uint8(0); c < columns; c++ {
			board.Rows[r][c].UpdateButton(r, c)
		}
	}
	if DebugModeBoard() {
		fmt.Printf("%v: %+v\n\n", Trace(), board.Rows)
	}
}

// CountButtons return the quantity of buttons in the board grid.
func (board *Board) countButtons() uint64 {
	if DebugModeBoard() {
		fmt.Printf("%v: %+v\n\n", Trace(), board.NumberOfButtons)
	}
	return board.NumberOfButtons
}

// CountEmptyButtons return the quantity of empty and Fulfilled buttons.
func (board *Board) countEmptyButtons() uint64 {
	rows, columns := board.RowNumber, board.ColumnNumber
	q := 0
	for r := uint8(0); r < rows; r++ {
		for c := uint8(0); c < columns; c++ {
			if board.Rows[r][c].Fulfilled == false {
				q++
			}
		}
	}
	if DebugModeBoard() {
		fmt.Printf("%v: %+v\n\n", Trace(), uint(q))
	}
	return uint64(q)
}

// TODO: Make it return dinamically depending on the quantity of players and number of columns
func (board *Board) GetStartPosition() {
	row := uint8(1)

	board.StartPositions = Positions{
		{Row: row, Column: 1},
		{Row: row, Column: 5},
		{Row: row, Column: 9},
		{Row: row, Column: 13},
	}
}