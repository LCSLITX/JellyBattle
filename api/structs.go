package api

import "github.com/lucassauro/JellyBattle/game"

type PlayerResponse struct {
	Name string `json:"name"`
	Rank uint8  `json:"rank"`
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
