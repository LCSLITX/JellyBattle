package utils

import (
	"math/rand"
)
// RandomNumber returns a random uint8 between 0-n
func RandomNumber(n int) uint8 {
	i := int64(rand.Int())
	seed := rand.NewSource(i)
	random := rand.New(seed)
	return uint8(random.Intn(n))
}