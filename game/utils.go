package game

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

func DebugModeBoard() bool {
	return os.Getenv("DEBUG_MODE_BOARD") == "true"
}

func DebugModeGroups() bool {
	return os.Getenv("DEBUG_MODE_GROUPS") == "true"
}

func DebugModeGame() bool {
	return os.Getenv("DEBUG_MODE_GAME") == "true"
}

func DebugModePlayer() bool {
	return os.Getenv("DEBUG_MODE_PLAYER") == "true"
}

func DebugModePlayerList() bool {
	return os.Getenv("DEBUG_MODE_PLAYERLIST") == "true"
}

func DebugModeSpecial() bool {
	return os.Getenv("DEBUG_MODE_SPECIAL") == "true"
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

// TODO: Implement GenerateID
// GenerateID creates a random ID
func GenerateID() string {
	return ""
}

// VerifyDuplicateID returns true if ID is already used.
func VerifyDuplicateID(g *Groups, id string) bool {
	for _, v := range (*g) {
		if v.ID == id {
			return true
		}
	}
	return false
}