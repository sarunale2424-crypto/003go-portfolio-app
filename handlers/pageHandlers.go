package handlers

import (
	"html/template"
	"net/http"
)

var tmpl template.Template

type PageHandler struct {
}

func (p *PageHandler) IndexPageHandler(w http.ResponseWriter, r *http.Request) {
	tmpl = *template.Must(template.ParseFiles("templates/pages/index_page.html"))
	tmpl.Execute(w, nil)
	// w.Write([]byte("Welcome to Index Page"))
}

func (p *PageHandler) AboutPageHandler(w http.ResponseWriter, r *http.Request) {
	tmpl = *template.Must(template.ParseFiles("templates/pages/about_page.html"))
	tmpl.Execute(w, nil)
	// w.Write([]byte("Welcome to About Page"))
}
func (p *PageHandler) ContactPageHandler(w http.ResponseWriter, r *http.Request) {
	tmpl = *template.Must(template.ParseFiles("templates/pages/contact_page.html"))
	tmpl.Execute(w, nil)
	// w.Write([]byte("Welcome to Contact Page"))
}

func (p *PageHandler) ServicePageHandler(w http.ResponseWriter, r *http.Request) {
	tmpl = *template.Must(template.ParseFiles("templates/pages/service_page.html"))
	tmpl.Execute(w, nil)
	// w.Write([]byte("Welcome to Service Page"))
}

func (p *PageHandler) ProductPageHandler(w http.ResponseWriter, r *http.Request) {
	tmpl = *template.Must(template.ParseFiles("templates/pages/product_page.html"))
	tmpl.Execute(w, nil)
	// w.Write([]byte("Welcome to Product Page"))
}
