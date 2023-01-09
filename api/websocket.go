package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"github.com/lucassauro/JellyBattle/game"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: check,
}

const (
	// Time allowed to read the next pong message from the peer.
	pongWait = 15 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 512
)

// TODO:  check() must check the header and return bool depending of whether its from an accepted origin.
// https://pkg.go.dev/github.com/gorilla/websocket#hdr-Origin_Considerations
func check(r *http.Request) bool {
	if DebugModeWebSocket() {
		fmt.Printf("%v: %+v\n\n", game.Trace(""), r)
	}

	return true
}

func WSGame(w http.ResponseWriter, r *http.Request) {
	wsConn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Printf(err.Error())
	}

	defer wsConn.Close()

	for {
		var req GameRequest
		err = wsConn.ReadJSON(&req)
		if err != nil {
			fmt.Printf("%s\n\n", err.Error())
			break
		}

		if req.Message != "" {
			// message flow 
			fmt.Printf("Received MESSAGE: %s\n", req.Message)
		}

		if req.JumpPosition.Row != 0 && req.JumpPosition.Column != 0 {
			// player movement flow
			fmt.Printf("Received POSITION: %+v\n", req.JumpPosition)
		}

		// Response
		ok, err := json.Marshal("ok")
		if err != nil {
			fmt.Printf("%s\n\n", err.Error())
			break
		}

		err = wsConn.WriteMessage(1, ok) // TODO:  Check what first argument means
		if err != nil {
			fmt.Printf("write: %s\n", err)
			break
		}

		if DebugModeWebSocket() {
			fmt.Printf("%v: %+v\n\n", game.Trace(""), req)
			fmt.Printf("%v: %+v\n\n", game.Trace(""), ok)
			fmt.Printf("wsConn.RemoteAddr(): %s\n", wsConn.RemoteAddr())
		}
	}
}

func WSPlayerList(w http.ResponseWriter, r *http.Request) {
	wsConn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Printf(err.Error())
	}

	defer wsConn.Close()

	for {
		if DebugModeWebSocket() {
			fmt.Printf("wsConn.RemoteAddr(): %s\n", wsConn.RemoteAddr())
		}
	}
}
