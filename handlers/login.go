package handlers

import (
	"net/http"
	"text/template"
)

func (a *App) Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		template, _ := template.ParseFiles("templates/login.html")

		template.Execute(w, nil)
		return
	}

	sessionID, err := newSessionID()
	if err != nil {
		http.Error(w, "could not create session", http.StatusInternalServerError)
		return
	}

	a.Sessions[sessionID] = "user_1"

	http.SetCookie(w, &http.Cookie{
		Name:     "session_id",
		Value:    sessionID,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	})
}
