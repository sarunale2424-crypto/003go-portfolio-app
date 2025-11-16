package utils

import (
	"fmt"
	"html/template"
	"net/http"
)

var tmpl *template.Template

func RenderTemplate(w http.ResponseWriter, name string, data interface{}) {
	tmpl = template.Must(template.ParseFiles(
		fmt.Sprintf("templates/pages/%s", name),
		"templates/layout/layout.html"))
	tmpl.ExecuteTemplate(w, "layout.html", data)

}
