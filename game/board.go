package game

// NewBoard returns a Board instance.
func NewBoard(rows, columns, players uint8) IBoard {
	b := &Board{}
	b.GenerateBoard(rows, columns)
	b.PlayersNumber = players
	b.SpecialFulfillment = DEFAULT_SPECIAL_FULFILLMENT
	return b
}

// GetBoard returns Board object.
func (board *Board) GetBoard() Board {
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
}

// RandomizeBoard Randomizes all the buttons of the board.
func (board *Board) RandomizeBoard() {
	rows, columns := board.RowNumber, board.ColumnNumber

	for r := uint8(0); r < rows; r++ {
		for c := uint8(0); c < columns; c++ {
			board.Rows[r][c].RandomizeButton()
		}
	}
}

// GenerateRow() generates one row at once.
func (board *Board) GenerateRow() (row Row) {
	rowsLength, columns := uint8(len(board.Rows)), board.ColumnNumber

	for c := uint8(0); c < columns; c++ {
		b := Button{Row: rowsLength, Column: c} // Create an empty button
		b.RandomizeButton()                     // Randomize it
		row = append(row, b)                    // Add new button to created row of buttons
	}
	return row
}

// RoundRows rotates the board rows every round.
func (board *Board) RoundRows() {
	// TODO: Buttons attributes (Rows and Columns) are currently keeping original data. Need to decide if it needs to update Row number x to x+1, for example.
	r := board.GenerateRow()
	board.Rows = board.Rows[1:]        // remove top element
	board.Rows = append(board.Rows, r) // add last element
}

// CountButtons return the quantity of buttons in the board grid.
func (board *Board) countButtons() uint64 {
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
	return uint64(q)
}

// UpdateBoard updates all the buttons of the board by calling updateButton method.
func (board *Board) UpdateBoard() {
	// VALIDATE this
	// DONE: Buttons attributes (Rows and Columns) are currently keeping original data. Need to decide if it needs to update Row number x to x+1, for example.
	rows, columns := board.RowNumber, board.ColumnNumber

	for r := uint8(0); r < rows; r++ {
		for c := uint8(0); c < columns; c++ {
			board.Rows[r][c].UpdateButton(r, c)
		}
	}
}
