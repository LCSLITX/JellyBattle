package game

import (
	"fmt"
)

func NewGame(b Board, g Group) IGame {
	game := &Game{
		ID:    g.ID, // TODO: Function to Generate ID. Think about game ID and Group ID
		Board: b,
		Group: g,
	}

	if DebugModeGame() {
		fmt.Printf("%v: %+v\n\n", Trace(), game)
	}

	return game
}

func (game *Game) FindGame() Game {
	return Game{}
}
