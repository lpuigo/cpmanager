package http

import (
	"github.com/lpuig/cpmanager/html"
	"log/slog"
	"net/http"
)

func (s *Server) setupRoutes() {

	s.mux.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir(s.config.Dir_Images))))

	s.mux.HandleFunc("GET /{$}", func(w http.ResponseWriter, r *http.Request) {
		s.log.Log(r.Context(), slog.LevelInfo, "Get root")
		html.MainPage(w)
	})

	s.mux.HandleFunc("GET /path", func(w http.ResponseWriter, r *http.Request) {
		s.log.Log(r.Context(), slog.LevelInfo, "Get Path")
	})

	s.mux.HandleFunc("GET /value/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		s.log.Log(r.Context(), slog.LevelInfo, "Get Value", "value", id)
	})
}
