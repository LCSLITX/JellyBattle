package game

import (
	"fmt"
)

// NewGroups() constructor returns a groups instance.
func NewGroups() IGroups {
	g := &Groups{}

	if DebugModeGroups() {
		fmt.Printf("%v: %+v\n\n", Trace(""), g)
	}
	return g
}

// GetGroups() returns groups.
func (groups *Groups) GetGroups() *Groups {
	if DebugModeGroups() {
		fmt.Printf("%v: %+v\n\n", Trace(""), groups)
	}
	return groups
}

// AddGroup() adds a group to groups.
func (groups *Groups) AddGroup(group Group) {
	*groups = append(*groups, group)

	if DebugModeGroups() {
		fmt.Printf("%v: %+v\n\n", Trace(""), *groups)
	}
}

// RemoveGroup() removes a group of groups.
func (groups *Groups) RemoveGroup(g Group) {
	i := groups.FindGroup(g)
	(*groups)[i] = (*groups)[len(*groups)-1]
	(*groups) = (*groups)[:len(*groups)-1]

	if DebugModeGroups() {
		fmt.Printf("%v: %+v\n\n", Trace(""), *groups)
	}
}

// FindGroup() returns the index of a given group in the Groups array. If not present returns -1.
func (groups *Groups) FindGroup(g Group) (i int) {
	i = -1
	for k, v := range *groups {
		if v.GID == g.GID {
			i = k
			break
		}
	}
	return
}

// GetFirstGroup() returns the first group in the Groups array.
func (groups *Groups) GetFirstGroup() Group {
	if len(*groups) == 0 {
		return Group{}
	}
	return (*groups)[0]
}
