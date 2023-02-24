package game

import "fmt"


func (g *Game) HandleConnections() {
	fmt.Printf("rodou 3\n")

	for {
		select {
		case conn := <-g.Register:
			fmt.Printf("Registering")

			g.Connections[conn] = nil
		case conn := <-g.Unregister:
			fmt.Printf("Unregistering")
			if _, ok := g.Connections[conn]; ok {
				delete(g.Connections, conn)
				conn.Close()
			}
		case message := <-g.Broadcast:
			fmt.Printf("Broadcasting")
			for k, player := range g.Connections {
				select {
				case player.Send <- message:
				default:
					close(player.Send)
					delete(g.Connections, k)
				}
			}
		}
	}
}

func (g *Game) HandleWSEvents() {
	fmt.Printf("G.CONNECTIONS: %v\n", g.Connections)
	for {
		for conn := range g.Connections {
			var req GameRequest
			err := conn.ReadJSON(&req)
			if err != nil {
				fmt.Printf("%s\n\n", err.Error())
				break
			}
			if len(req.Message.Msg) != 0 {
				// message flow
				fmt.Printf("Received MESSAGE: %+v\n", req.Message.Msg)
				g.Broadcast <- req.Message.Msg
				conn.WriteMessage(1, []byte("Message received and broadcasted"))
			} else {
				// player movement flow
				fmt.Printf("Received POSITION: %+v\n", req.JumpPosition)
				g.Group.Players[req.Name].CurrentPosition = req.JumpPosition
				conn.WriteMessage(1, []byte("Movement received and broadcasted"))
			}
		}
	}
}
