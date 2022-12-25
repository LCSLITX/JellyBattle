package players

// import "github.com/lucassauro/J.B.Remake/game/games"

func NewPlayer(name string) IPLayer {
	return &Player{
		Name: name,
		Rank: uint8(0),
	}
}

// TODO: Implement CreateGame
// CreateGame 
// func (player *Player) CreateGame() games.Game {
// 	g := games.Game{}
// 	return g
// }