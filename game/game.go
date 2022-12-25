package game

import (
	"fmt"

	"github.com/lucassauro/J.B.Remake/game/utils"
)

func NewGame(b Board, g Group) IGame {
	game := &Game{
		ID:    g.ID, // TODO: Function to Generate ID. Think about game ID and Group ID
		Board: b,
		Group: g,
	}

	if utils.DebugMode() {
		fmt.Printf("%v: %+v\n\n", utils.Trace(), game)
	}

	return game
}

func (game *Game) FindGame() Game {
	return Game{}
}
