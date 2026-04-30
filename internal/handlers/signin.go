package handlers

import (
	"log/slog"
	"net/http"
	"text/template"
)

func (a *App) Signin(w http.ResponseWriter, r *http.Request) {
	slog.Info("")
	slog.Info("===== Signin Handler =====")

	slog.Info("Getting uri params")
	responseType := r.URL.Query().Get("response_type")
	redirectURI := r.URL.Query().Get("redirect_uri")
	clientID := r.URL.Query().Get("client_id")
	scope := r.URL.Query().Get("scope")
	state := r.URL.Query().Get("state")

	if r.Method == "GET" {
		loginURI := CreateURI("/login", clientID, responseType, redirectURI, scope, state)

		template, _ := template.ParseFiles("templates/login.html")

		template.Execute(w, struct {
			SubmitURI string
		}{
			SubmitURI: loginURI,
		})

		return
	}

	if r.Method != "POST" {
		http.Error(w, "wrong method call", http.StatusBadRequest)
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")

	if _, ok := a.Users[username]; !ok {
		http.Error(w, "wrong username", http.StatusUnauthorized)
		return
	}

	if a.Users[username] != password {
		http.Error(w, "wrong password", http.StatusUnauthorized)
		return
	}

	sessionID, err := newSessionID()
	if err != nil {
		http.Error(w, "could not create session", http.StatusInternalServerError)
		return
	}

	a.Sessions[sessionID] = username

	http.SetCookie(w, &http.Cookie{
		Name:     "session_id",
		Value:    sessionID,
		Path:     "/",
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
	})

	l := CreateURI("/oauth/authorize", clientID, responseType, redirectURI, scope, state)

	http.Redirect(w, r, l, http.StatusFound)
}
