package server

import (
	"html/template"
	"net/http"
)

const (
	LayoutTemplateName = "layout.html"
	HomeTemplate       = "home"
)

type Server struct {
	Mux       *http.ServeMux
	Templates map[string]*template.Template
}

func NewServer() *Server {
	server := &Server{
		Mux:       http.NewServeMux(),
		Templates: parseTemplates(),
	}

	server.registerRoutes()
	return server
}

func (s Server) registerRoutes() {
	s.Mux.HandleFunc("GET /{$}", s.homeView)
}

func (s Server) ListenAndServe(addr string) error {
	return http.ListenAndServe(addr, s.Mux)
}

func parseTemplates() map[string]*template.Template {
	parsedTemplates := make(map[string]*template.Template)

	parsedTemplates[HomeTemplate] = template.Must(parseLayout().ParseFiles("./templates/home.html"))

	return parsedTemplates
}

func parseLayout() *template.Template {
	return template.Must(template.ParseFiles("./templates/header.html", "./templates/layout.html"))
}
