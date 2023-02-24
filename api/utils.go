package api

import (
	"fmt"
	"net/http"
	"os"
)

func DebugModeWebSocket() bool {
	return os.Getenv("DEBUG_MODE_WEBSOCKET") == "true"
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "http://localhost:3333")
}

func getHomePage() (s string) {
	file, err := os.ReadFile("frontend/page.html")
	if err != nil {
		fmt.Println(err)
	}

	s = string(file)
	
	return s
}
