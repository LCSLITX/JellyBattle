package game

// NewButton() constructor returns a Button instance.
func NewButton(row, column uint8) (button Button) {
	b := Button{
		Row: row,
	}
	return b
}

// RandomizeButton() uses probabilistic randomization to fulfill buttons with specials or to leave it empty.
func (button *Button) RandomizeButton() {
	rand := RandomNumber(ZERO_TO_NINETY_NINE) // generate a random number from 0 to 99. 100 possible random numbers.

	// Different comparisons between the random number and a given number will result in a probability.
	// For example, there is 30% chance of a random number between 1 and 100 be bigger than 70. Because there is 70 possible numbers smaller than or equal to 70. And only 30 possible numbers bigger than 70.
	if rand > DEFAULT_SPECIAL_FULFILLMENT { 
		button.Fulfilled = false
		button.Special = Special{}
	} else {
		button.Fulfilled = true
		button.Special = GenerateSpecial()
	}
}

// UpdateButton() updates row and columns inside a button
func (button *Button) UpdateButton(row, column uint8) {
	button.Row, button.Column = row, column
}

// VALIDATE: 
// DONE:  Method to erase button. To be used when a player uses a button.
// EraseButtonSpecial() updates a button by removing its special. Supposed to be used after a players uses a button.
func (button *Button) EraseButtonSpecial() {
	button.Special = Special{}
	button.Fulfilled = false
}
