package boards

type IBoard interface {
	// GetBoard returns the board
	GetBoard() Board
	// GenerateBoard creates a board with specified number of rows and columns
	GenerateBoard(rows, columns uint8)
	// RandomizeBoard populates a board with buttons
	RandomizeBoard()
	// UpdateBoard updates the board buttons
	UpdateBoard()

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
	// RandomizeButton generate an empty button or special button
	RandomizeButton()
	// UpdateButton updates row and columns inside each button
	UpdateButton(row, column uint8)
	// EraseButtonSpecial updates a button by removing its special
	EraseButtonSpecial()
}
