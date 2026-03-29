package handlers

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/fmo/oauth/internal"
)

func (a *App) Authorize(w http.ResponseWriter, r *http.Request) {
	clientID := r.URL.Query().Get("client_id")
	redirectURI := r.URL.Query().Get("redirect_uri")
	responseType := r.URL.Query().Get("response_type")
	scope := r.URL.Query().Get("scope")

	if _, ok := a.Clients[clientID]; !ok {
		http.Error(w, "client is not defined", http.StatusBadRequest)
		return
	}

	if a.Clients[clientID].RedirectURI != redirectURI {
		http.Error(w, "redirect url is not matching", http.StatusBadRequest)
		return
	}

	if responseType != "code" {
		http.Error(w, "response type is not valid", http.StatusBadRequest)
		return
	}

	// get user
	userID, err := a.getUserFromRequest(r)
	if err != nil {
		u, _ := url.Parse("/login")
		q := u.Query()
		q.Set("client_id", clientID)
		q.Set("response_type", responseType)
		q.Set("redirect_uri", redirectURI)
		q.Set("scope", scope)

		u.RawQuery = q.Encode()

		http.Redirect(w, r, u.String(), http.StatusFound)
		return
	}

	code, _ := internal.GenerateCode()
	a.StoreCode(code, userID, clientID, redirectURI, scope)

	redirect := redirectURI + "?code=" + code
	http.Redirect(w, r, redirect, http.StatusFound)
}

func (a *App) getUserFromRequest(r *http.Request) (string, error) {
	cookie, err := r.Cookie("session_id")
	if err != nil || cookie.Value == "" {
		return "", fmt.Errorf("no session")
	}

	userID, ok := a.Sessions[cookie.Value]
	if !ok {
		return "", fmt.Errorf("invalid session")
	}

	return userID, nil
}
