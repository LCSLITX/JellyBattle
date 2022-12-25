package utils

import (
	"fmt"
	"math/rand"
	"os"
	"runtime"
	"strings"
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

func DebugMode() bool {
	return os.Getenv("DEBUG_MODE") == "true"
}

func Trace() string {
	pc := make([]uintptr, 15)
	n := runtime.Callers(2, pc)
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()
	file := strings.Split(frame.File, "/")
	function := strings.Split(frame.Function, "/")
	return fmt.Sprintf("[%s:%d - %s]", file[7], frame.Line, function[3])
}
