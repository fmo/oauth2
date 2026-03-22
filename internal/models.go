package internal

import "time"

type AuthCode struct {
	UserID      string
	ClientID    string
	RedirectURI string
	ExpiresAt   time.Time
}

type Client struct {
	Secret      string
	RedirectURI string
}
