package main

import (
	"fmt"
	"net/http"

	"github.com/fmo/oauth/handlers"
)

func main() {
	mux := http.NewServeMux()

	app := handlers.NewApp()

	mux.HandleFunc("/login", app.Login)
	mux.HandleFunc("/oauth/authorize", app.Authorize)
	mux.HandleFunc("/oauth/token", app.Token)

	fmt.Println("server runs on 8080")
	http.ListenAndServe(":8080", mux)
}
