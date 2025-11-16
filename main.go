package main

import (
	"003go-portfolio-app/handlers"
	"log"
	"log/slog"
	"net/http"
)

func main() {

	//setup static files

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	//initialise handlers
	pageHandler := handlers.PageHandler{}

	//setup routes
	http.HandleFunc("/pages", pageHandler.IndexPageHandler)

	http.HandleFunc("/pages/about", pageHandler.AboutPageHandler)

	http.HandleFunc("/pages/contact", pageHandler.ContactPageHandler)

	http.HandleFunc("/pages/service", pageHandler.ServicePageHandler)

	http.HandleFunc("/pages/product", pageHandler.ProductPageHandler)

	// setup server
	slog.Info("Server is running at http://localhost:4000")
	err := http.ListenAndServe(":4000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
