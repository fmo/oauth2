package handlers

import (
	"fmt"
	"time"

	"github.com/fmo/oauth/internal"
)

type App struct {
	Sessions map[string]string
	Clients  map[string]Client
	Codes    map[string]internal.AuthCode
	Users    map[string]string
}

type Client struct {
	Secret      string
	RedirectURI string
}

func NewApp() *App {
	users := map[string]string{
		"fmo": "123123",
	}

	clients := map[string]Client{
		"web_client": Client{
			Secret:      "axaa",
			RedirectURI: "http://localhost:8081/callback",
		},
	}

	return &App{
		Sessions: make(map[string]string),
		Clients:  clients,
		Codes:    make(map[string]internal.AuthCode),
		Users:    users,
	}
}

func (a *App) StoreCode(code, userID, clientID, redirectURI, scope string) {
	a.Codes[code] = internal.AuthCode{
		UserID:      userID,
		ClientID:    clientID,
		RedirectURI: redirectURI,
		ExpiresAt:   time.Now().Add(2 * time.Minute),
		Scope:       scope,
	}
}

func (a *App) ConsumeCode(code, clientID, redirectURI string) (*internal.AuthCode, error) {
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
