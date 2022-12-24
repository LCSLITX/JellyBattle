package boards

import (
	"github.com/lucassauro/J.B.Remake/game_logic/utils"
)

// func NewButton(row, column uint8) (button Button) {
// 	b := Button{
// 		Row: row,
// 	}

// 	return b
// }

// RandomizeButton uses probabilistic randomization to fulfill special buttons.
func (button *Button) RandomizeButton() {
	rand := utils.RandomNumber(ZERO_TO_NINETY_NINE) // 100 possible random numbers

	if rand > SPECIAL_FULFILLMENT_DEFAULT {
		button.Fulfilled = false
		button.Special = Special{}
	} else {
		button.Fulfilled = true
		button.Special = GenerateSpecial()
	}
}
