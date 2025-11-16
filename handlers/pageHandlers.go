package handlers

import "net/http"

type PageHandler struct {
}

func (p *PageHandler) IndexPageHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to Index Page"))
}

func (p *PageHandler) AboutPageHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to About Page"))
}
func (p *PageHandler) ContactPageHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to Contact Page"))
}

func (p *PageHandler) ServicePageHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to Service Page"))
}

func (p *PageHandler) ProductPageHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to Product Page"))
}
