package game

import "fmt"

func NewGames() (IGames) {
	games := &Games{}

	if DebugModeGame() {
		fmt.Printf("%v: %+v\n\n", Trace(), games)
	}

	return games
}

func (games *Games) AddGame(g Game) {
	*games = append(*games, g)
}

func (games *Games) GetGames() Games {
	return *games
}