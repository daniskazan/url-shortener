package application

import (
	"github.com/daniskazan/url-shortener/cmd/url-shortener/routing"
	"github.com/daniskazan/url-shortener/configs"
	"github.com/go-chi/chi/v5"
	"net/http"
)

type Application struct {
	router chi.Router
	config configs.ApplicationConfig
}

func NewApplication() *Application {
	app := &Application{router: routing.ConfigureRouting(), config: configs.ApplicationConfig{Addr: ":8000"}}
	return app
}

func (a *Application) Start() error {
	err := http.ListenAndServe(a.config.Addr, a.router)
	if err != nil {
		panic("failed")
	}
	return nil
}

func (a *Application) ShutDown() error {
	return nil
}
