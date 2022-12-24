package boards

type IBoard interface {
	// GetBoard returns the board
	GetBoard() Board
	
	// GenerateBoard creates a board with specified number of rows and columns
	GenerateBoard(rows, columns uint8)

	// RandomizeBoard populates a board with buttons
	RandomizeBoard()

	// CountButtons return the quantity of buttons in the grid.
	CountButtons() uint64

	// CountEmptyButtons return the quantity of empty and Fulfilled buttons.
	CountEmptyButtons() uint64
}

type IButton interface {
	// RandomizeButton generate an empty button or spell button
	RandomizeButton()
}
