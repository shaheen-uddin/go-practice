package main

import (
	"net/http"

	"github.com/shaheen-uddin/go-practice/pkg/handlers"
)

const portNumber = ":8080"

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	//fmt.Println(fmt.Fprintf("starting application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
