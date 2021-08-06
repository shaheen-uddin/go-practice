package render

import (
	"fmt"
	"html/template"
	"net/http"
)

func RenderTemplate(w http.ResponseWriter, tmpl string) {
	parseTemplate, _ := template.ParseFiles("./teplates" + tmpl)
	if err := parseTemplate.Execute(w, nil); err != nil {
		fmt.Println("error parsing template", err)
		return
	}
}
