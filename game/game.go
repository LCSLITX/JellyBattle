package game

import (
	"fmt"
	"time"
)

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
	}

	if DebugModeGame() {
		fmt.Printf("%v: %+v\n\n", Trace(), game)
	}

	return game, nil
}

func (game *Game) GetGame() Game {
	return *game
}

// TODO: StartGame wont be used as games do not have time limit, apparently. But rounds does. it'll be useful there.
func (game *Game) StartGame() bool {
	game.Started = true

	game.DistributePlayers()
	game.Board.GeneratePreviewRow()
	game.Board.RoundRows()

	ticker := time.NewTicker(DEFAULT_ROUND_TIME * time.Second)
	timer := time.NewTimer(DEFAULT_GAME_TIME * time.Minute)

	go func() {
		for {
			select {
			case <-Finish:
				game.FinishGame()
				ticker.Stop()
			case <-timer.C:
				game.FinishGame()
				ticker.Stop()
			case <-ticker.C:
				game.RoundingUp()
			}
		}
	}()

	if DebugModeGame() {
		fmt.Printf("%v: %+v\n\n", Trace(), game)
	}

	return game.Started
}

func (game *Game) RoundingUp() {
	game.MovePlayers()
	game.Board.GeneratePreviewRow()
	game.Board.RoundRows()
}

func (game *Game) FinishGame() bool {
	game.Finished = true
	return game.Finished
}

func (game *Game) GetWinners() Group {
	return game.Group
}

// Dead function should re-order game.Group so the first player in group is the winner and the last is the loser.
func (game *Game) Dead(p Player) {
	// TODO: Create Logic to rank players.
	game.Deaths++
	if game.Deaths == 3 {
		Finish <- true
	}
}

// TODO: think about how DoDamage is gonna work. Apparently it will need Board and special.
// Because every weapon has one type of calculating the attacked player.
// Lasers do damage horizontally and vertically
// Nuke do damage to all players
// Air Strike seems to have a percentage to do damage to every visible player in the board
// Random Drop do damage to one visible random player
// Handbag do damage to all players within one square
// Blast do damage to all players, while its increases distance, decrease blast damage.
// CalculateAttackablePlayers receives Special and Player (who used the special) and returns a list of Players who are supposed to be damaged.

func (game *Game) MovePlayers() {
	for _, p := range game.Group.Players {
		p.JumpTo()
		if DebugModeBoard() {
			fmt.Printf("%v: %+v\n\n", Trace(), p)
		}
	}
}

func (game *Game) DistributePlayers() {
	game.Board.GetStartPosition()
	b := game.Board.StartPositions
	for k := range game.Group.Players {
		game.Group.Players[k].Position = b[k]
	}
}
