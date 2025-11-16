package handlers

import (
	"003go-portfolio-app/models"
	"003go-portfolio-app/utils"
	"net/http"
)

type PageHandler struct {
}

func (p *PageHandler) IndexPageHandler(w http.ResponseWriter, r *http.Request) {
	page := models.Page{
		Title:   "Index Page",
		Heading: "Welcome to Index Page",
		Content: "This is content of index page",
	}

	utils.RenderTemplate(w, "index_page.html", page)
}

func (p *PageHandler) AboutPageHandler(w http.ResponseWriter, r *http.Request) {
	page := models.Page{
		Title:   "About Page",
		Heading: "Welcome to About Page",
		Content: "This is content of about page",
	}

	utils.RenderTemplate(w, "about_page.html", page)
}
func (p *PageHandler) ContactPageHandler(w http.ResponseWriter, r *http.Request) {
	page := models.Page{
		Title:   "Contact Page",
		Heading: "Welcome to Contact Page",
		Content: "This is content of contact page",
	}

	utils.RenderTemplate(w, "contact_page.html", page)
}

func (p *PageHandler) ServicePageHandler(w http.ResponseWriter, r *http.Request) {
	page := models.Page{
		Title:   "Service Page",
		Heading: "Welcome to Service Page",
		Content: "This is content of service page",
	}

	utils.RenderTemplate(w, "service_page.html", page)
}

func (p *PageHandler) ProductPageHandler(w http.ResponseWriter, r *http.Request) {
	page := models.Page{
		Title:   "Product Page",
		Heading: "Welcome to Product Page",
		Content: "This is content of product page",
	}

	utils.RenderTemplate(w, "product_page.html", page)
}
