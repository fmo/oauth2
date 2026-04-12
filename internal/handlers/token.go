package handlers

import (
	"net/http"
)

func (a *App) Token(w http.ResponseWriter, r *http.Request) {
	clientID := r.FormValue("client_id")
	clientSecret := r.FormValue("client_secret")
	redirectURI := r.FormValue("redirect_uri")
	code := r.FormValue("code")

	if _, ok := a.Clients[clientID]; !ok {
		http.Error(w, "client does not exist", http.StatusBadRequest)
		return
	}

	client := a.Clients[clientID]

	if client.Secret != clientSecret {
		http.Error(w, "wrong client secret", http.StatusUnauthorized)
		return
	}

	grantType := r.FormValue("grant_type")

	switch grantType {
	case "client_credentials":
		w.Write([]byte(`{
			"access_token": "abc123",
			"token_type": "Bearer",
			"expires_in": 3600
		}`))
	case "authorization_code":
		_, err := a.ConsumeCode(code, clientID, redirectURI)
		if err != nil {
			http.Error(w, "code is wrong", http.StatusUnauthorized)
			return
		}

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
