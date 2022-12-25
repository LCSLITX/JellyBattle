package game

import (
	"github.com/lucassauro/J.B.Remake/game/utils"
)

// NewButton returns a Button instance.
func NewButton(row, column uint8) (button Button) {
	b := Button{
		Row: row,
	}
	return b
}

// RandomizeButton uses probabilistic randomization to fulfill special buttons.
func (button *Button) RandomizeButton() {
	rand := utils.RandomNumber(ZERO_TO_NINETY_NINE) // 100 possible random numbers

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

// EraseButtonSpecial updates a button by removing its special
// TODO: Method to erase button. To be used when a player uses this very button. DEPENDS ON PLAYER AND GAME IMPLEMENTATIONS.
func (button *Button) EraseButtonSpecial() {
	button.Special = Special{}
	button.Fulfilled = false
}
