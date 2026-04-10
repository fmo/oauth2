package handlers

import (
	"net/http"
	"text/template"
)

func (a *App) Consent(w http.ResponseWriter, r *http.Request) {
	clientID := r.URL.Query().Get("client_id")
	responseType := r.URL.Query().Get("response_type")
	redirectURI := r.URL.Query().Get("redirect_uri")
	scope := r.URL.Query().Get("scope")
	state := r.URL.Query().Get("state")

	if r.Method == "POST" {
		userID, err := GetUserFromRequest(r, a.Sessions)
		if err != nil {
			loginURI := CreateURI("/login", clientID, responseType, redirectURI, scope, state)
			http.Redirect(w, r, loginURI, http.StatusFound)
			return
		}

		if r.FormValue("scopes") == scope {
			a.Consents[userID] = Consent{
				ClientID: clientID,
				Scope:    scope,
			}

			authorizeURI := CreateURI("/oauth/authorize", clientID, responseType, redirectURI, scope, state)

			http.Redirect(w, r, authorizeURI, http.StatusFound)
			return
		}

	}

	consentURI := CreateURI("/consent", clientID, responseType, redirectURI, scope, state)

	template, err := template.ParseFiles("templates/consent.html")
	if err != nil {
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	vars := struct {
		FormAction string
		Scopes     string
	}{consentURI, scope}

	template.Execute(w, vars)
}
