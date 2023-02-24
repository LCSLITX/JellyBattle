package api

import (
	"os"
	"time"
)

func init() {
	os.Setenv("PORT", ":3333")
}

const (
	DEFAULT_INTERVAL = time.Second * 15

	// FRONTEND FILESERVER
	DOT_FRONTEND = "./frontend"
	_FRONTEND_ = "/frontend/"
	_FRONTEND = "/frontend"


	//  ENDPOINTS
	_PLAYER_CREATE = "/player/create"
	
	_PLAYERLIST     = "/playerlist"
	_PLAYERLIST_CREATEGROUP = "/playerlist/creategroup"

	_GAMES = "/games"
	_GAMES_GETBYID = "/games/getbyid"
	
	_GROUPS = "/groups"
	
	_GAME_START = "/game/start"
	_GAME_CREATE = "/game/create"

	_ROOT = "/"
	_PING = "/ping"
	_TEST = "/test"

	// WS
	_WS_START_GID = "/ws/start/"
	_WS_PLAYER_CREATE = "/ws/player/create"
	_WS_PLAYERLIST = "/ws/playerlist/"

)
