package api

import (
	"os"
	"time"
)

var PORT string = os.Getenv("PORT")

const (
	DEFAULT_INTERVAL = time.Second * 15
)
