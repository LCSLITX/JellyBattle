package api

import "github.com/lucassauro/JellyBattle/game"

type PlayerResponse struct {
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
	Message      string        `json:"message"`
}

type PlayerListResponse struct {
	List []PlayerResponse `json:"list"`
}

type GameResponse struct {
	Game game.Game `json:"game"`
}

type GamesResponse struct {
	List game.Games `json:"list"`
}
