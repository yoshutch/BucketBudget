package server

import (
	"net/http"
)

func (s *Server) homeView(w http.ResponseWriter, req *http.Request) {
	err := s.Templates[HomeTemplate].ExecuteTemplate(w, LayoutTemplateName, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
