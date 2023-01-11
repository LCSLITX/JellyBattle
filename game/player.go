package game

import (
	"fmt"
)

// NewPlayer() constructor returns a player instance.
func NewPlayer(name string) IPLayer {
	id, err := GenerateID()
	if err != nil {
		fmt.Println(err)
	}

	player := &Player{
		ID:   id,
		Name: name,
		Rank: uint8(0),
	}

	if DebugModePlayer() {
		fmt.Printf("%v: %+v\n\n", Trace(""), player)
	}

	return player
}

// GetPlayer() returns a player.
func (player *Player) GetPlayer() Player {
	if DebugModePlayer() {
		fmt.Printf("%v: %+v\n\n", Trace(""), *player)
	}

	return *player
}

// GetJumpDistance() returns a player jump distance which could be:
//
// - 2 on a normal basis;
//
// - 3 with Triple Jump buff activated.
func (player *Player) GetJumpDistance() {
	player.JumpDistance = 2
	for _, v := range player.Buffs {
		if v.SpecialName == specialName.String(11) {
			player.JumpDistance = 3
		}
	}
	if DebugModePlayer() {
		fmt.Printf("%v: %+v\n\n", Trace(""), player.JumpDistance)
	}
}

// TODO: Think about how its gonna work. Return a matrix maybe, or just return 4 limits to construct the matrix in front end. Not ideal in my opinion.
// The horizontal order of buttons (rows) is inverse to the cartesian plane.
func (player *Player) GetJumpArea() {
	// jd := player.JumpDistance
	// row, column := player.Position.Row, player.Position.Column

	// superiorRowLimit := row - jd
	// inferiorRowLimit := row + jd
	// leftColumnLimit := column - jd
	// rightColumnLimit := column + jd
}

// JumpTo() is the function that changes player's Position to its JumpPosition.
func (player *Player) JumpTo() {
	player.Position.Row, player.Position.Column = player.JumpPosition.Row, player.JumpPosition.Column

	if DebugModePlayer() {
		fmt.Printf("%v: %+v\n\n", Trace(""), *player)
	}
}

// NextJump() is the function that receives the next JumpPosition. Where the player intent to move on next round.
func (player *Player) NextJump(p Position) {
	player.JumpPosition.Row, player.JumpPosition.Column = p.Row, p.Column

	if DebugModePlayer() {
		fmt.Printf("%v: %+v\n\n", Trace(""), *player)
	}
}

// SufferDamage() decreases Health Points of a player.
func (player *Player) SufferDamage(s Special) {
	if s.Type == "Weapon" {
		player.Life -= s.Damage
	} else if s.Type == "Healing" {
		player.Life += s.Damage
	}
}

// CauseDamage() decreases Health Points of adversary players.
// TODO: think about how CauseDamage is gonna work. Apparently it will need Board and special.
// Because every weapon has one type of calculating the attacked player:
// Lasers do damage horizontally and vertically;
// Nuke do damage to all players even if invisible;
// Air Strike seems to have a percentage to do damage to every visible player in the board;
// Random Drop do damage to one visible random player. If the only visible player is the one who used it, he will receive the damage;
// Handbag do damage to all players within one square;
// Blast do damage to all players, while it increases its ratio and decreases blast damage;
// Players with Shield buff do not receive damage with any weapon.
// Maybe a function called CalculateAttackablePlayers that receives Special and Player (who used the special) and returns a list of Players who are supposed to be damaged and then call CauseDamage of these players.
func (player *Player) CauseDamage(s Special, b Board) {

}
