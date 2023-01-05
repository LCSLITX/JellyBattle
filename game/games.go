package game

import "fmt"

// NewGames() constructor returns a games instance.
func NewGames() (IGames) {
	games := &Games{}

	if DebugModeGame() {
		fmt.Printf("%v: %+v\n\n", Trace(), games)
	}

	return games
}

// AddGame() adds a game to games array instance.
func (games *Games) AddGame(g Game) {
	*games = append(*games, g)
}

// GetGames() returns games array.
func (games *Games) GetGames() Games {
	return *games
}