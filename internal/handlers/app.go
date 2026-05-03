package handlers

import (
	"time"

	"github.com/fmo/oauth/internal"
)

type App struct {
	Sessions     map[string]string
	Clients      map[string]Client
	Codes        map[string]AuthCode
	Users        map[string]string
	Consents     map[string]Consent
	AccessTokens map[string]AccessToken
	Logger       *internal.Logger
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

type AccessToken struct {
	UserID    string
	ClientID  string
	Scope     string
	ExpiresAt time.Time
}

func NewApp(logger *internal.Logger) *App {
	users := map[string]string{
		"fmo": "123123",
	}

	clients := map[string]Client{
		"web_client": Client{
			Secret:      "demo-client-secret",
			RedirectURI: "http://localhost:8081/callback",
		},
	}

	consents := map[string]Consent{}

	accessTokens := map[string]AccessToken{}

	return &App{
		Sessions:     make(map[string]string),
		Clients:      clients,
		Codes:        make(map[string]AuthCode),
		Users:        users,
		Consents:     consents,
		AccessTokens: accessTokens,
		Logger:       logger,
	}
}
