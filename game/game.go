package game

import (
	"fmt"
	"time"
)

// NewGame() constructor returns a game instance.
func NewGame(g *Groups) (IGame, error) {
	if len(*g) == 0 {
		return nil, fmt.Errorf("There's no group to create a game.")
	} else if len((*g)[0].Players) == 0 {
		return nil, fmt.Errorf("Cannot use an empty group.")
	}

	b := NewBoard(DEFAULT_ROWS, DEFAULT_COLUMNS, DEFAULT_PLAYERS_NUMBER).GetBoard()

	game := &Game{
		ID:    (*g)[0].ID, // TODO: Function to Generate ID. Think about game ID and Group ID
		Board: b,
		Group: (*g)[0],
		Chat: make([]Message, 0),
		Finish: make(chan bool),
		Deaths: make(Players),
	}

	if DebugModeGame() {
		fmt.Printf("%v: %+v\n\n", Trace(""), game)
	}

	return game, nil
}

// GetGame() returns Game object.
func (game *Game) GetGame() Game {
	return *game
}

// StartGame() set variable Started to true and proceeds with a few specific procedures proper to initial moment as distribute Players to its starting positions, generate a few rows, start a ticker to count round time etc.
func (game *Game) StartGame() bool {
	game.Started = true

	game.DistributePlayers()
	game.Board.GeneratePreviewRow()
	game.Board.RoundRows()

	ticker := time.NewTicker(DEFAULT_ROUND_TIME * time.Second)

	// TODO: Think if game won't have time limit.
	// timer := time.NewTimer(DEFAULT_GAME_TIME * time.Minute)

	go func() {
		for {
			select {
			case <-game.Finish:
				game.FinishGame()
				ticker.Stop()
			// case <-timer.C:
			// 	game.FinishGame()
			// 	ticker.Stop()
			case msg := <-game.Broadcast:
				game.Chat = append(game.Chat, msg)
				game.Send <- true
			case <-ticker.C:
				game.RoundUp()
			}
		}
	}()

	if DebugModeGame() {
		fmt.Printf("%v: %+v\n\n", Trace(""), game)
	}

	return game.Started
}

// RoundUp() will execute the functions responsible for in-between round actions.
func (game *Game) RoundUp() {
	game.MovePlayers()
	game.Board.GeneratePreviewRow()
	game.Board.RoundRows()
}

// Games will only finish with FORFEIT or DEATH.
// FinishGame() is supposed to be called with 3 DEATHS and FORFEIT if there are only two players alive and set variable Finished to true.
func (game *Game) FinishGame() bool {
	game.Finished = true
	return game.Finished
}

// Think if it will be really necessary.
func (game *Game) GetWinners() Group {
	return game.Group
}

// Think if death is gonna happen this way.
// Dead function should re-order game.Group so the first player in group is the winner and the last is the loser.
func (game *Game) Dead(p Player) {
	game.Deaths[p.ID] = &p
}

// MovePlayers() moves players to its jump positions.
func (game *Game) MovePlayers() {
	for _, p := range game.Group.Players {
		p.JumpTo()
		if DebugModeBoard() {
			fmt.Printf("%v: %+v\n\n", Trace(""), p)
		}
	}
}

// DistributePlayers() distributes players to its starting positions over the board.
func (game *Game) DistributePlayers() {
	game.Board.GetStartPosition()
	b := game.Board.StartPositions
	i := 0
	for k := range game.Group.Players {
		game.Group.Players[k].Position = b[i]
		i ++
	}
}

