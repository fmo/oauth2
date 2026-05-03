package main

import (
	"fmt"
	"net/http"

	"github.com/fmo/oauth/internal"
	"github.com/fmo/oauth/internal/handlers"
)

func main() {
	mux := http.NewServeMux()

	logger := internal.NewLogger()

	app := handlers.NewApp(logger)

	mux.HandleFunc("/signin", app.Signin)
	mux.HandleFunc("/consent", app.Consent)
	mux.HandleFunc("/oauth/authorize", app.Authorize)
	mux.HandleFunc("/oauth/token", app.Token)
	mux.HandleFunc("/oauth/sessions", app.ListSessions)

	fmt.Println("server runs on 8080")
	http.ListenAndServe(":8080", mux)
}
