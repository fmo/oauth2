package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/oauth/authorize", func(w http.ResponseWriter, r *http.Request) {
		clients := getClients()

		clientID := r.URL.Query().Get("client_id")
		redirectURL := r.URL.Query().Get("redirect_uri")
		responseType := r.URL.Query().Get("response_type")

		if _, ok := clients[clientID]; !ok {
			http.Error(w, "client is not defined", http.StatusBadRequest)
			return
		}

		if clients[clientID] != redirectURL {
			http.Error(w, "redirect url is not matching", http.StatusBadRequest)
			return
		}

		if responseType != "code" {
			http.Error(w, "response type is not valid", http.StatusBadRequest)
			return
		}

		cookie, err := r.Cookie("session_id")
		if err != nil || cookie.Value == "" {
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}
	})

	fmt.Println("server runs on 8080")
	http.ListenAndServe(":8080", mux)
}

func getClients() map[string]string {
	return map[string]string{
		"web_client": "http://localhost:8081/callback",
	}
}
