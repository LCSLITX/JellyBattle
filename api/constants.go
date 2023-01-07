package api

import (
	"os"
	"time"
)

func init() {
	os.Setenv("PORT", ":3333")
}

const (
	DEFAULT_INTERVAL = time.Second * 15
)
