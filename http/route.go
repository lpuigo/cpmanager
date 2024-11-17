package http

import (
	"github.com/lpuig/cpmanager/html"
	"log/slog"
	"net/http"
)

func (s *Server) setupRoutes() {

	//TODO use compress middleware github.com/klauspost/compress
	s.mux.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir(s.config.Dir_Images))))
	s.mux.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir(s.config.Dir_Css))))
	s.mux.Handle("/script/", http.StripPrefix("/script/", http.FileServer(http.Dir(s.config.Dir_Script))))

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
