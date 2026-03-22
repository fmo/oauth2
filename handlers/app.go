package handlers

import "github.com/fmo/oauth/internal"

type App struct {
	Sessions map[string]string
	Clients  map[string]internal.Client
}

func NewApp() *App {
	return &App{
		Sessions: make(map[string]string),
		Clients:  internal.GetClients(),
	}
}
