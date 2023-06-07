package app

import (
	"allfather/internal/config"
	"allfather/internal/transport/handler"
	"net/http"
)

type App struct {
	Config config.Config
}

func (app *App) Run() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler.RootHandler)
	mux.HandleFunc("/user", handler.UserHandler)

	http.ListenAndServe(":9000", mux)
}
