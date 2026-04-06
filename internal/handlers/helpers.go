package handlers

import (
	"fmt"
	"net/http"
	"net/url"
)

func GetUserFromRequest(r *http.Request, sessions map[string]string) (string, error) {
	cookie, err := r.Cookie("session_id")
	if err != nil || cookie.Value == "" {
		return "", fmt.Errorf("no session")
	}

	userID, ok := sessions[cookie.Value]
	if !ok {
		return "", fmt.Errorf("invalid session")
	}

	return userID, nil
}

func CreateURI(base, clientID, responseType, redirectURI, scope, state string) string {
	u, _ := url.Parse(base)
	q := u.Query()
	q.Set("client_id", clientID)
	q.Set("response_type", responseType)
	q.Set("redirect_uri", redirectURI)
	q.Set("scope", scope)
	q.Set("state", state)

	u.RawQuery = q.Encode()

	return u.String()
}
