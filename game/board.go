package game

import (
	"fmt"
)


// NewBoard() constructor returns an empty Board instance.
func NewBoard(rows, columns, players uint8) IBoard {
	// To know more about Object Oriented Programming (OOP) in Go and how to add methods to a "class":
	// https://stackoverflow.com/questions/28800672/how-to-add-new-methods-to-an-existing-type-in-go
	b := &Board{}
	b.GenerateBoard(rows, columns)
	b.PlayersNumber = players
	b.SpecialFulfillment = DEFAULT_SPECIAL_FULFILLMENT

	return b
}

// GetBoard() returns Board object.
func (board *Board) GetBoard() Board {
	if DebugModeBoard() {
		fmt.Printf("%v: %+v\n\n", Trace(), *board)
	}
	return *board
}

// Generateboard() generates a board with specified number of rows and columns.
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

// GeneratePreviewRow() generates one row at once with randomized buttons and set it up as board.PreviewRow.
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

// GetPreviewRow() returns current preview row.
func (board *Board) GetPreviewRow() Row {
	if DebugModeBoard() {
		fmt.Printf("%v: %+v\n\n", Trace(), board.PreviewRow)
	}
	return board.PreviewRow
}

// RoundRows() rotates the rows and update its buttons to correctly correspond to its respective new row. Supposed to be used when changing rounds.
// For example, considering a matrix of 5 (rows) by 20 (columns). The first row (at index 0) will be the bottom one, and the fifth (at index 4) will be the top one.
// RoundRows will delete the first index, so Row X will now be Row X - 1. And it will also append the PreviewRow to the last index.
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

// CountButtons() return the quantity of buttons in the board grid.
func (board *Board) countButtons() uint64 {
	if DebugModeBoard() {
		fmt.Printf("%v: %+v\n\n", Trace(), board.NumberOfButtons)
	}
	return board.NumberOfButtons
}

// CountEmptyButtons() return the quantity of empty and Fulfilled buttons.
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

// GetStartPosition() returns the players start positions.
// TODO: Make it return dinamically depending on the quantity of players and number of columns.
func (board *Board) GetStartPosition() {
	row := uint8(1)

	board.StartPositions = Positions{
		{Row: row, Column: 1},
		{Row: row, Column: 5},
		{Row: row, Column: 9},
		{Row: row, Column: 13},
	}
}