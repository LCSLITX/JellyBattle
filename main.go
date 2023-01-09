package main

import (
	"github.com/lucassauro/JellyBattle/api"
)

func main() {
	api.Server()
}

// func main() {
// os.Setenv("DEBUG_MODE_STRUCTS", "true")
// os.Setenv("DEBUG_MODE_GROUPS", "true")
// os.Setenv("DEBUG_MODE_GAME", "true")
// os.Setenv("DEBUG_MODE_PLAYER", "true")
// os.Setenv("DEBUG_MODE_PLAYERLIST", "true")
// os.Setenv("DEBUG_MODE_BOARD", "true")
// os.Setenv("DEBUG_MODE_SPECIAL", "true")

// 	rows := game.DEFAULT_ROWS
// 	columns := game.DEFAULT_COLUMNS
// 	players := game.DEFAULT_PLAYERS_NUMBER

// 	pL := game.NewPlayerList()
// 	gs := game.NewGroups()

// 	p1 := game.NewPlayer("Lucas")
// 	pL.AddPlayer(p1.GetPlayer())

// 	p2 := game.NewPlayer("P2")
// 	pL.AddPlayer(p2.GetPlayer())

// 	p3 := game.NewPlayer("P3")
// 	pL.AddPlayer(p3.GetPlayer())

// 	p4 := game.NewPlayer("P4")
// 	pL.AddPlayer(p4.GetPlayer())

// 	_, err := pL.GroupFourPlayers(gs.GetGroups())
// 	if err != nil {
// 		fmt.Println(err)
// 		os.Exit(1)
// 	}
// 	gs.GetGroups()

// 	board := game.NewBoard(rows, columns, players)

// 	// board.GeneratePreviewRow()
// 	// board.RoundRows()

// 	// board.GeneratePreviewRow()

// 	// board.RoundRows()
// 	g, err := game.NewGame(board.GetBoard(), gs.GetFirstGroup())
// 	if err != nil {
// 		fmt.Println(err)
// 		os.Exit(1)
// 	}

// 	g.StartGame()

// 	time.Sleep(2 * time.Second)

// 	gm := g.GetGame()

// 	gm.Group.Players[0].NextJump(game.Position{
// 		Row: 2,
// 		Column: gm.Group.Players[0].Position.Column,
// 	})

// 	gm.Group.Players[1].NextJump(game.Position{
// 		Row: 2,
// 		Column: gm.Group.Players[1].Position.Column,
// 	})

// 	gm.Group.Players[2].NextJump(game.Position{
// 		Row: 2,
// 		Column: gm.Group.Players[2].Position.Column,
// 	})

// 	gm.Group.Players[3].NextJump(game.Position{
// 		Row: 2,
// 		Column: gm.Group.Players[3].Position.Column,
// 	})

// 	time.Sleep(4 * time.Second)

// 	gm.Group.Players[0].NextJump(game.Position{
// 		Row: 3,
// 		Column: gm.Group.Players[0].Position.Column,
// 	})

// 	gm.Group.Players[1].NextJump(game.Position{
// 		Row: 3,
// 		Column: gm.Group.Players[1].Position.Column,
// 	})

// 	gm.Group.Players[2].NextJump(game.Position{
// 		Row: 3,
// 		Column: gm.Group.Players[2].Position.Column,
// 	})

// 	gm.Group.Players[3].NextJump(game.Position{
// 		Row: 3,
// 		Column: gm.Group.Players[3].Position.Column,
// 	})

// 	time.Sleep(5 * time.Second)

// 	gm.Group.Players[0].NextJump(game.Position{
// 		Row: 4,
// 		Column: gm.Group.Players[0].Position.Column,
// 	})

// 	gm.Group.Players[1].NextJump(game.Position{
// 		Row: 4,
// 		Column: gm.Group.Players[1].Position.Column,
// 	})

// 	gm.Group.Players[2].NextJump(game.Position{
// 		Row: 4,
// 		Column: gm.Group.Players[2].Position.Column,
// 	})

// 	gm.Group.Players[3].NextJump(game.Position{
// 		Row: 4,
// 		Column: gm.Group.Players[3].Position.Column,
// 	})

// 	time.Sleep(10 * time.Second)

// 	// p1.NextJump(game.Position{
// 	// 	Row: 3,
// 	// 	Column: p1.GetPlayer().Position.Column,
// 	// })

// 	// p2.NextJump(game.Position{
// 	// 	Row: 3,
// 	// 	Column: p2.GetPlayer().Position.Column,
// 	// })

// 	// p3.NextJump(game.Position{
// 	// 	Row: 3,
// 	// 	Column: p3.GetPlayer().Position.Column,
// 	// })

// 	// p4.NextJump(game.Position{
// 	// 	Row: 3,
// 	// 	Column: p4.GetPlayer().Position.Column,
// 	// })

// 	g.FinishGame()
// }
