package api

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
)

func GetHome(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("[%s] %s%s\n", r.Method, r.Host, r.URL)

	homeTemplate := template.Must(template.New("Home").Parse(getHomePage()))

	homeTemplate.Execute(w, r.Host)
}

func GetPing(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("[%s] %s%s\n", r.Method, r.Host, r.URL)
	// wsConn, _ := upgrader.Upgrade(w, r, nil)
	// wsConn.SetPingHandler()
	io.WriteString(w, "Pong!\n")
}
