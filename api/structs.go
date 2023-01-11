package api

import "github.com/lucassauro/JellyBattle/game"

type PlayerResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Rank uint8  `json:"rank"`
}

type PlayerRequest struct {
	Name string `json:"name"`
}

type GameRequest struct {
	GID          string        `json:"gid"`
	PID          string        `json:"pid"`
	Name         string        `json:"name"`
	JumpPosition game.Position `json:"jumpPosition"`
	Message      game.Message  `json:"message"`
}

type PlayerListResponse struct {
	List []PlayerResponse `json:"list"`
}

type GameResponse struct { // Struct Game refers to a game, composed by a group of players and a board.
	Started  bool // Not sure yet
	Finished bool // Not sure yet
	ID       string
	Deaths   game.Players
	Board    game.Board
	Group    game.Group
	Chat     []game.Message // Not sure yet
	// Timer    time.Duration // Not sure yet
}

type GamesResponse map[string]GameResponse

type GroupsResponse struct {
	List *game.Groups `json:"list"`
}
