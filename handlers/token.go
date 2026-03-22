package handlers

import (
	"net/http"
)

func (a *App) Token(w http.ResponseWriter, r *http.Request) {
	clientID := r.FormValue("client_id")
	clientSecret := r.FormValue("client_secret")

	grantType := r.FormValue("grant_type")

	if client, ok := a.Clients[clientID]; ok {
		if client.Secret != clientSecret {
			http.Error(w, "not matching secret", http.StatusUnauthorized)
			return
		}
	}

	switch grantType {
	case "client_credentials":
		w.Write([]byte(`{
			"access_token": "abc123",
			"token_type": "Bearer",
			"expires_in": 3600
		}`))
	case "authorization_code":
		isOIDC := true

		if isOIDC {
			w.Write([]byte(`{
				"access_token": "abc123",
				"id_token": "fake-jwt-token",
				"token_type": "Bearer",
				"expires_in": 3600
			}`))
			return
		}

		w.Write([]byte(`{
			"access_token": "abc123",
			"token_type": "Bearer",
			"expires_in": 3600
		}`))
		return
	case "refresh_token":
	default:
		http.Error(w, "unsupported grant", http.StatusBadRequest)
		return
	}
}
