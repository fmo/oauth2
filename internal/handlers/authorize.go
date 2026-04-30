// Package handlers take care of requests
package handlers

import (
	"log/slog"
	"net/http"
)

func (a *App) Authorize(w http.ResponseWriter, r *http.Request) {
	slog.Info("")
	slog.Info("===== Authorize Handler =====")

	slog.Info("Getting uri parameters")
	clientID := r.URL.Query().Get("client_id")
	redirectURI := r.URL.Query().Get("redirect_uri")
	responseType := r.URL.Query().Get("response_type")
	scope := r.URL.Query().Get("scope")
	state := r.URL.Query().Get("state")

	slog.Info("Checking client id if it is defined")
	if _, ok := a.Clients[clientID]; !ok {
		http.Error(w, "client is not defined", http.StatusBadRequest)
		return
	}

	slog.Info("Checking redirect uri if it is matching")
	if a.Clients[clientID].RedirectURI != redirectURI {
		http.Error(w, "redirect url is not matching", http.StatusBadRequest)
		return
	}

	slog.Info("Checking if the response type is code")
	if responseType != "code" {
		http.Error(w, "response type is not valid", http.StatusBadRequest)
		return
	}

	slog.Info("Check if session cookie exists")
	userID, err := GetUserFromRequest(r, a.Sessions)
	if err != nil {
		slog.Info("Session has not started, redirecting to signin page")
		loginURI := CreateURI("/signin", clientID, responseType, redirectURI, scope, state)
		http.Redirect(w, r, loginURI, http.StatusFound)
		return
	}

	if _, ok := a.Consents[userID]; !ok {
		consentURI := CreateURI("/consent", clientID, responseType, redirectURI, scope, state)
		http.Redirect(w, r, consentURI, http.StatusFound)
		return
	}

	code, err := a.GenerateCode()
	if err != nil {
		http.Error(w, "cant generate code", http.StatusInternalServerError)
		return
	}
	a.StoreCode(code, userID, clientID, redirectURI, scope)

	rduri := CreateRedirectURI(redirectURI, code, state)

	http.Redirect(w, r, rduri, http.StatusFound)
}
