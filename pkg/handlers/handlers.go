package handlers

import (
	"net/http"

	"github.com/shaheen-uddin/go-practice/pkg/render"
)

func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.html")
	//fmt.Fprintf(w, "hellow from home")

}

func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.html")
	//fmt.Fprintf(w, "hello from about")
}
