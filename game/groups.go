package game

// import (
// 	"fmt"

// 	"github.com/lucassauro/J.B.Remake/game/utils"
// )

// func NewGroups() IGroups {
// 	g := &Groups{
// 		Groups: []Group{},
// 	}

// 	if utils.DebugMode() {
// 		fmt.Printf("%v: %+v\n\n", utils.Trace(), g)
// 	}
// 	return g
// }

// func (groups *Groups) GetGroups() Groups {
// 	if utils.DebugMode() {
// 		fmt.Printf("%v: %+v\n\n", utils.Trace(), *groups)
// 	}
// 	return *groups
// }

// func (groups *Groups) AddGroup(group Group) {
// 	groups.Groups = append(groups.Groups, group)
// }

// // func (groups *Groups) RemoveGroup(Index) {
// // 	groups.Groups = groups.Groups[1:]
// // }