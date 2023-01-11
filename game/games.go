package game

import "fmt"

// NewGames() constructor returns a games instance.
func NewGames() IGames {
	games := &Games{}

	if DebugModeGame() {
		fmt.Printf("%v: %+v\n\n", Trace(""), games)
	}

	return games
}

// AddGame() adds a game to games map instance.
func (games *Games) AddGame(g Game) {
	// *games = append(*games, g)
	(*games)[g.ID] = &g
}

// GetGames() returns games map.
func (games *Games) GetGames() Games {
	return *games
}

// FindGame() returns a specific game of games map or an error if game is not existent.
func (games *Games) FindGame(id string) (g *Game, err error) {
	g, ok := (*games)[id]
	if !ok {
		return &Game{}, fmt.Errorf("Game ID not found.")
	}
	return
}

