package handlers

import (
	"time"
)

type App struct {
	Sessions map[string]string
	Clients  map[string]Client
	Codes    map[string]AuthCode
	Users    map[string]string
	Consents map[string]Consent
}

type Client struct {
	Secret      string
	RedirectURI string
}

type AuthCode struct {
	UserID      string
	ClientID    string
	RedirectURI string
	ExpiresAt   time.Time
	Scope       string
}

type Consent struct {
	ClientID string
	Scope    string
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

	consents := map[string]Consent{}

	return &App{
		Sessions: make(map[string]string),
		Clients:  clients,
		Codes:    make(map[string]AuthCode),
		Users:    users,
		Consents: consents,
	}
}
