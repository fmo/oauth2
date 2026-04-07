package handlers

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"time"
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
	q.Add("client_id", clientID)
	q.Add("response_type", responseType)
	q.Add("redirect_uri", redirectURI)
	q.Add("scope", scope)
	q.Add("state", state)

	u.RawQuery = q.Encode()

	return u.String()
}

func CreateRedirectURI(redirectURI, code, state string) string {
	u, _ := url.Parse(redirectURI)

	q := u.Query()
	q.Add("code", code)
	q.Add("state", state)

	u.RawQuery = q.Encode()

	return u.String()
}

func (a *App) GenerateCode() (string, error) {
	b := make([]byte, 32) // 256-bit

	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	// URL-safe (no + / =)
	code := base64.RawURLEncoding.EncodeToString(b)

	return code, nil
}

func (a *App) StoreCode(code, userID, clientID, redirectURI, scope string) {
	a.Codes[code] = AuthCode{
		UserID:      userID,
		ClientID:    clientID,
		RedirectURI: redirectURI,
		ExpiresAt:   time.Now().Add(2 * time.Minute),
		Scope:       scope,
	}
}

func (a *App) ConsumeCode(code, clientID, redirectURI string) (*AuthCode, error) {
	data, ok := a.Codes[code]
	if !ok {
		return nil, fmt.Errorf("invalid code")
	}

	// check expiration
	if time.Now().After(data.ExpiresAt) {
		delete(a.Codes, code)
		return nil, fmt.Errorf("code expired")
	}

	// check client binding
	if data.ClientID != clientID {
		return nil, fmt.Errorf("client mismatch")
	}

	// check redirect URI (important!)
	if data.RedirectURI != redirectURI {
		return nil, fmt.Errorf("redirect mismatch")
	}

	// one-time use → delete immediately
	delete(a.Codes, code)

	return &data, nil
}
