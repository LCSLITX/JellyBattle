package game

import (
	"crypto/md5"
	"fmt"
	"math/rand"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/google/uuid"
)

// RandomNumber returns a random uint8 between 0 and n.
func RandomNumber(n int) uint8 {
	seed := rand.NewSource(randomSeed())
	random := rand.New(seed)
	return uint8(random.Intn(n))
}

// randomSeed returns a random int64 to be used as a seed. Its summing time with random int because it is called a lot of times per second.
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

func DebugModeStructs() bool {
	return os.Getenv("DEBUG_MODE_STRUCTS") == "true"
}

// Trace("") used for debug purpose, to show file, line and function on printing information.
// if receives "function" return only the function name.
// Otherwise returns file, line and function.
func Trace(object string) string {
	pc := make([]uintptr, 15)
	n := runtime.Callers(2, pc)
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()
	file := strings.Split(frame.File, "/")
	function := strings.Split(frame.Function, "/")
	if object == "function" {
		return function[len(function)-1]
	} else {
		return fmt.Sprintf("[%s:%d - %s]", file[len(file)-1], frame.Line, function[len(function)-1])
	}
}

// TODO: Implement GenerateID
// GenerateID creates a random ID. Well, not yet.
func GenerateID() (string, error) {
	id, err := uuid.New().MarshalBinary()
	if err != nil {
		return "", err
	}
	sum := md5.Sum(id)

	s := fmt.Sprintf("%x", sum)

	return s, nil
}

// VerifyDuplicateID returns true if ID is already used.
func VerifyDuplicateID(g *Groups, id string) bool {
	for _, v := range *g {
		if v.GID == id {
			return true
		}
	}
	return false
}
