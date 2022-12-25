package game

func NewPlayer(name string) IPLayer {
	return &Player{
		Name: name,
		Rank: uint8(0),
	}
}

// TODO: Implement CreateGame
// CreateGame 
func (player *Player) CreateGame() Game {
	g := Game{}
	return g
}