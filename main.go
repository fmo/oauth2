package main

import (
	"fmt"
	"net/http"
)

var sessions = make(map[string]string)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/login", http.HandlerFunc(login))

	mux.HandleFunc("/oauth/authorize", http.HandlerFunc(AuthorizeHandler))

	mux.HandleFunc("/oauth/token", http.HandlerFunc(TokenHandler))

	fmt.Println("server runs on 8080")
	http.ListenAndServe(":8080", mux)
}
