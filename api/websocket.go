package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/lucassauro/JellyBattle/game"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: check,
}

// TODO:  check() must check the header and return bool depending of whether its from an accepted origin.
// https://pkg.go.dev/github.com/gorilla/websocket#hdr-Origin_Considerations
func check(r *http.Request) bool {
	if DebugModeWebSocket() {
		fmt.Printf("%v: %+v\n\n", game.Trace(), r)
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
		
		ok, err := json.Marshal("ok")
		if err != nil {
			fmt.Printf("%s\n\n", err.Error())
			break
		}

		err = wsConn.WriteMessage(1, ok)
		if err != nil {
			fmt.Printf("write: %s\n", err)
			break
		}

		if DebugModeWebSocket() {
			fmt.Printf("%v: %+v\n\n", game.Trace(), req)
			fmt.Printf("%v: %+v\n\n", game.Trace(), ok)
			fmt.Printf("wsConn.RemoteAddr(): %s\n", wsConn.RemoteAddr())
		}
	}
}
