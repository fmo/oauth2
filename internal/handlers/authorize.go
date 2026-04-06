package handlers

import (
	"net/http"
	"net/url"
)

func (a *App) Authorize(w http.ResponseWriter, r *http.Request) {
	clientID := r.URL.Query().Get("client_id")
	redirectURI := r.URL.Query().Get("redirect_uri")
	responseType := r.URL.Query().Get("response_type")
	scope := r.URL.Query().Get("scope")
	state := r.URL.Query().Get("state")

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
	userID, err := GetUserFromRequest(r, a.Sessions)
	if err != nil {
		loginURI := CreateURI("/login", clientID, responseType, redirectURI, scope, state)
		http.Redirect(w, r, loginURI, http.StatusFound)
		return
	}

	code, err := a.GenerateCode()
	if err != nil {
		http.Error(w, "cant generate code", http.StatusInternalServerError)
		return
	}
	a.StoreCode(code, userID, clientID, redirectURI, scope)

	u, _ := url.Parse(redirectURI)
	q := u.Query()
	q.Add("code", code)
	q.Add("state", state)

	u.RawQuery = q.Encode()

	http.Redirect(w, r, u.String(), http.StatusFound)
}
