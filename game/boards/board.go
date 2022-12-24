package boards

func NewBoard(rows, columns, players uint8) IBoard {
	b := &Board{}
	b.GenerateBoard(rows, columns)
	b.Players = players
	b.SpecialFulfillment = SPECIAL_FULFILLMENT_DEFAULT
	return b
}

func (board *Board) GetBoard() Board {
	return *board
}

// Generate board with specified number of rows and columns
func (board *Board) GenerateBoard(rows, columns uint8) {
	board.RowNumber, board.ColumnNumber = rows, columns

	q := uint64(0)

	for r := uint8(0); r < rows; r++ {
		row := Row{} // Create an empty row of buttons
		for c := uint8(0); c < columns; c++ {
			b := Button{}        // Create an empty button TODO: Implement the logic to create a random button. DEPENDS ON SPELLS
			row = append(row, b) // Add new button to created row of buttons
			q += 1
		}
		board.Rows = append(board.Rows, row) // Add new row of buttons into Rows array.
	}
	board.NumberOfButtons = q
}

// Count Buttons return the quantity of buttons in the grid.
func (board *Board) CountButtons() uint64 {
	return board.NumberOfButtons
}

	// CountEmptyButtons return the quantity of empty and Fulfilled buttons.
func (board *Board) CountEmptyButtons() uint64 {
	// TODO: RETURN CORRECT VALUE
	rows, columns := board.RowNumber, board.ColumnNumber
	
	for r := uint8(0); r < rows; r++ {
		for c := uint8(0); c < columns; c++ {
			board.Rows[r][c].RandomizeButton()
		}
	}
	return board.NumberOfButtons
}

func (board *Board) RandomizeBoard() {
	rows, columns := board.RowNumber, board.ColumnNumber

	for r := uint8(0); r < rows; r++ {
		for c := uint8(0); c < columns; c++ {
			board.Rows[r][c].RandomizeButton()
		}
	}
}