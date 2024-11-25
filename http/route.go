package http

import (
	"github.com/lpuig/cpmanager/html"
	"github.com/lpuig/cpmanager/html/comp"
	"github.com/lpuig/cpmanager/model/consultant"
	"log/slog"
	"net/http"
)

func (s *Server) setupRoutes() {

	//TODO use compress middleware github.com/klauspost/compress
	s.mux.Handle("/Assets/", http.StripPrefix("/Assets/", http.FileServer(http.Dir(s.config.Dir_Asset))))

	s.mux.HandleFunc("GET /{$}", func(w http.ResponseWriter, r *http.Request) {
		s.log.Log(r.Context(), slog.LevelInfo, "Get root")
		html.MainPage(w, s.manager)
	})

	s.mux.HandleFunc("GET /action/consult/add", func(w http.ResponseWriter, r *http.Request) {
		s.log.Log(r.Context(), slog.LevelInfo, "Action Add consult")

		s.manager.Consultants.AddNewConsultant()
		comp.ConsultantsList(s.manager.Consultants).Render(w)
	})

	s.mux.HandleFunc("POST /action/consult/add", func(w http.ResponseWriter, r *http.Request) {
		s.log.Log(r.Context(), slog.LevelInfo, "Action Add consult")
		defer r.Body.Close()

		newConsultant := consultant.Consultant{
			FistName: r.FormValue("FirstName"),
			LastName: r.FormValue("LastName"),
		}
		s.manager.Consultants.Add(newConsultant)
		w.Header().Add("HX-Trigger-After-Swap", "closeModal")
		comp.ConsultantsList(s.manager.Consultants).Render(w)
	})

	s.mux.HandleFunc("GET /action/consult/addmodal", func(w http.ResponseWriter, r *http.Request) {
		s.log.Log(r.Context(), slog.LevelInfo, "Action Add consult")

		comp.AddConsultantModal().Render(w)
	})

	s.mux.HandleFunc("GET /action/closemodal", func(w http.ResponseWriter, r *http.Request) {
		s.log.Log(r.Context(), slog.LevelInfo, "Action Close")
		comp.ModalHook().Render(w)
	})

}
