package game

// NewButton returns a Button instance.
func NewButton(row, column uint8) (button Button) {
	b := Button{
		Row: row,
	}
	return b
}

// RandomizeButton uses probabilistic randomization to fulfill buttons with specials or to leave it empty.
func (button *Button) RandomizeButton() {
	rand := RandomNumber(ZERO_TO_NINETY_NINE) // 100 possible random numbers

	if rand > DEFAULT_SPECIAL_FULFILLMENT {
		button.Fulfilled = false
		button.Special = Special{}
	} else {
		button.Fulfilled = true
		button.Special = GenerateSpecial()
	}
}

// UpdateButton updates row and columns inside a button
func (button *Button) UpdateButton(row, column uint8) {
	button.Row, button.Column = row, column
}

// VALIDATE
// DONE: Method to erase button. To be used when a player uses this very button. DEPENDS ON PLAYER AND GAME IMPLEMENTATIONS.
// EraseButtonSpecial updates a button by removing its special
func (button *Button) EraseButtonSpecial() {
	button.Special = Special{}
	button.Fulfilled = false
}
