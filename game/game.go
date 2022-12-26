package game

import (
	"fmt"
	"time"
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

func (game *Game) StartGame() bool {
	timer := time.NewTimer(5 * time.Minute)
	game.Started = true

	go func() {
		<-timer.C
		game.FinishGame()
	}()

	return game.Started
}

func (game *Game) FinishGame() bool {
	game.Finished = true
	return game.Finished
}

func (game *Game) GetWinners() Group {
	return Group{}
}
