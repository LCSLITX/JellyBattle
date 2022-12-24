package boards

type IBoard interface {
	// GetBoard returns the board
	GetBoard() Board
	// GenerateBoard creates a board with specified number of rows and columns
	GenerateBoard(rows, columns uint8)
	// RandomizeBoard populates a board with buttons
	RandomizeBoard()

	// RoundRows rotates the board rows every round.
	RoundRows()
	// GenerateRow generate one row at once
	GenerateRow() Row
	
	// CountButtons return the quantity of buttons in the grid.
	countButtons() uint64
	// CountEmptyButtons return the quantity of empty and Fulfilled buttons.
	countEmptyButtons() uint64
}

type IButton interface {
	// RandomizeButton generate an empty button or spell button
	RandomizeButton()
}
