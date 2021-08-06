package main

import (
	"log"
	"net/http"

	"github.com/shaheen-uddin/go-practice/pkg/config"
	"github.com/shaheen-uddin/go-practice/pkg/handlers"
	"github.com/shaheen-uddin/go-practice/pkg/render"
)

const portNumber = ":8080"

func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()

	if err != nil {
		log.Fatal("Error creating template cache")
	}

	app.TemplateCache = tc

	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	//fmt.Println(fmt.Fprintf("starting application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
