package handlers

import (
	"encoding/json"
	"net/http"
	"strings"
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
		authCode, err := a.ConsumeCode(code, clientID, redirectURI)
		if err != nil {
			http.Error(w, "code is wrong", http.StatusUnauthorized)
			return
		}

		token, err := a.GenerateCode()
		if err != nil {
			http.Error(w, "cant generate token", http.StatusInternalServerError)
			return
		}

		a.StoreToken(token, authCode.UserID, authCode.ClientID, authCode.Scope)

		var isOIDC bool
		scopes := strings.Fields(authCode.Scope)
		for _, scope := range scopes {
			if scope == "openid" {
				isOIDC = true
			}
		}

		if isOIDC {
			resp := map[string]any{
				"access_token": token,
				"id_token":     "blabla",
				"token_type":   "Bearer",
				"expires_in":   3600,
			}

			w.Header().Set("Content-Type", "application/json")
			_ = json.NewEncoder(w).Encode(resp)

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
