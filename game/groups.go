package game

import (
	"fmt"
)

func NewGroups() IGroups {
	g := &Groups{}

	if DebugModeGroups() {
		fmt.Printf("%v: %+v\n\n", Trace(), g)
	}
	return g
}

func (groups *Groups) GetGroups() *Groups {
	if DebugModeGroups() {
		fmt.Printf("%v: %+v\n\n", Trace(), groups)
	}
	return groups
}

func (groups *Groups) AddGroup(group Group) {
	*groups = append(*groups, group)

	if DebugModeGroups() {
		fmt.Printf("%v: %+v\n\n", Trace(), *groups)
	}
}

func (groups *Groups) RemoveGroup(g Group) {
	i := groups.FindGroup(g)
	(*groups)[i] = (*groups)[len(*groups)-1]
	(*groups) = (*groups)[:len(*groups)-1]

	if DebugModeGroups() {
		fmt.Printf("%v: %+v\n\n", Trace(), *groups)
	}
}

// FindGroup returns the index of a given group in the Groups. If not present returns -1.
func (groups *Groups) FindGroup(g Group) (i int) {
	i = -1
	for k, v := range *groups {
		if v.ID == g.ID {
			i = k
			break
		}
	}
	return
}

func (groups *Groups) GetFirstGroup() Group {
	if len(*groups) == 0 {
		return Group{}
	}
	return (*groups)[0]
}
