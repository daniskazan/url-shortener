package main

import (
	"github.com/daniskazan/url-shortener/cmd/url-shortener/application"
)

func main() {
	app := application.NewApplication()
	app.ConfigureRouting()
	err := app.Start()
	if err != nil {
		panic("fail start app")
	}
}
