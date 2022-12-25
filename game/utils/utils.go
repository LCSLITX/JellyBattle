package utils

import (
	"math/rand"
	"time"
)

// RandomNumber returns a random uint8 between 0-n
func RandomNumber(n int) uint8 {
	seed := rand.NewSource(randomSeed())
	random := rand.New(seed)
	return uint8(random.Intn(n))
}

func randomSeed() (n int64) {
	i1 := time.Now().UnixNano()
	i2 := int64(rand.Int())
	i3 := i1 + i2
	return i3
}
