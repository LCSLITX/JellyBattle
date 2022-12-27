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

// TODO: StartGame wont be used as games do not have time limit, apparently. But rounds does. it'll be useful there.
func (game *Game) StartGame() bool {
	timer := time.NewTimer(10 * time.Second)
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
