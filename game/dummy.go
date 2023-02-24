package game

// TODO: Decide future of Dummies

func NewDummy(name string) IDummy {
	return &Dummy{
		Name: name,
		Rank: uint8(0),
	}
}
