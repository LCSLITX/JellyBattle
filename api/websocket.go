package api

import (
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
		fmt.Println(err.Error())
	}
	defer wsConn.Close()

	gid := r.URL.Path[len(_WS_START_GID):]

	g, err := game.GAMES.FindGame(gid)
	if err != nil {
		wsConn.WriteMessage(1, []byte("game not found"))
		if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
			fmt.Printf("error: %v\n", err)
		}
		wsConn.Close()
		return
	}

	wsConn.WriteMessage(1, []byte("game found"))

	go g.HandleConnections()

	WSGameHandle(g, wsConn, w, r)

	fmt.Printf("rodou 2\n")
}

func WSGameHandle(g *game.Game, wsConn *websocket.Conn, w http.ResponseWriter, r *http.Request) {
	if DebugModeWebSocket() {
		fmt.Printf("%v: %+v\n\n", game.Trace(""), g)
		fmt.Printf("wsConn.RemoteAddr(): %s\n\n", wsConn.RemoteAddr())
	}

	go g.HandleWSEvents()
}

func WSPlayerList(w http.ResponseWriter, r *http.Request) {
	wsConn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err.Error())
	}

	defer wsConn.Close()

	for {
		if DebugModeWebSocket() {
			fmt.Printf("wsConn.RemoteAddr(): %s\n", wsConn.RemoteAddr())
		}
	}
}

func WSPlayerCreate(w http.ResponseWriter, r *http.Request) {
	wsConn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err.Error())
	}

	defer wsConn.Close()

	var req PlayerRequest
	err = wsConn.ReadJSON(&req)
	if err != nil || req.Name == "" {
		wsConn.WriteMessage(1, []byte("error creating player"))
		fmt.Printf("error: %v\n", err)
		wsConn.Close()
		return
	}

	p := game.NewPlayer(req.Name, wsConn).GetPlayer()
	game.AVAILABLEPLAYERSLIST.AddPlayer(p)
	fmt.Printf("PLAYER: %+v", p)
	str := fmt.Sprintf("Player %v created. Rank: %v.", p.Name, p.Rank)
	wsConn.WriteMessage(1, []byte(str))
}
