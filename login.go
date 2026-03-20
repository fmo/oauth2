package main

import "net/http"

func login(w http.ResponseWriter, r *http.Request) {
	sessionID := "some_cookie"

	sessions[sessionID] = "user_1"

	http.SetCookie(w, &http.Cookie{
		Name:  "session_id",
		Value: sessionID,
		Path:  "/",
	})
}
