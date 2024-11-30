package http

import (
	"fmt"
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
		s.log.InfoContext(r.Context(), "Get root")
		html.MainPage(w, s.manager)
	})

	// Add New Consultant ==============================================================================================
	//

	// Open New Consultant Modal
	s.mux.HandleFunc("GET /action/consult/addmodal", func(w http.ResponseWriter, r *http.Request) {
		s.log.InfoContext(r.Context(), "consult/addmodal: open new consult modal")

		comp.AddConsultantModal().Render(w)
	})

	// Retrieve New Consultant info from front Form. Return updated consultant list and trigger modal closing
	s.mux.HandleFunc("POST /action/consult/add", func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		// Trig modal closing
		w.Header().Add("HX-Trigger-After-Swap", "closeModal")

		// Create new consultant
		newConsultant := consultant.Consultant{
			FirstName: r.FormValue("FirstName"),
			LastName:  r.FormValue("LastName"),
		}
		id := s.manager.Consultants.Add(newConsultant)
		s.log.InfoContext(r.Context(), fmt.Sprintf("consult/add: add new consult with id %s (%s)", id, newConsultant.Name()))
		comp.ConsultantsList(s.manager.Consultants).Render(w)
	})

	// Update Consultant ===============================================================================================
	//

	// Open Update Consultant Modal for consultant with url given id
	s.mux.HandleFunc("GET /action/consult/updatemodal/{id}", func(w http.ResponseWriter, r *http.Request) {

		consultId := r.PathValue("id")
		consult, found := s.manager.Consultants.Get(consultId)
		if !found {
			s.log.ErrorContext(r.Context(), fmt.Sprintf("consult/updatemodal: open update consult modal failed: unknown id %s", consultId))
			w.WriteHeader(http.StatusNotFound)
			return
		}

		s.log.Log(r.Context(), slog.LevelInfo, fmt.Sprintf("consult/updatemodal: open update consult modal for id %s (%s)", consultId, consult.Name()))
		comp.UpdateConsultantModal(consult).Render(w)
	})

	// Retrieve New Consultant info from front Form. Return updated consultant list and trigger modal closing
	s.mux.HandleFunc("POST /action/consult/update/{id}", func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		consultId := r.PathValue("id")
		consult, found := s.manager.Consultants.Get(consultId)
		if !found {
			s.log.ErrorContext(r.Context(), fmt.Sprintf("consult/update: update consult failed: unknown id %s", consultId))
			w.WriteHeader(http.StatusNotFound)
			return
		}

		// Trig modal closing
		w.Header().Add("HX-Trigger-After-Swap", "closeModal")

		// update consultant
		consult.FirstName = r.FormValue("FirstName")
		consult.LastName = r.FormValue("LastName")
		s.manager.Consultants.Update(consult)
		s.log.InfoContext(r.Context(), fmt.Sprintf("consult/update: update consult with id %s (%s)", consultId, consult.Name()))
		comp.ConsultantLine(consult).Render(w)
	})

	// Return close modal elem
	s.mux.HandleFunc("GET /action/closemodal", func(w http.ResponseWriter, r *http.Request) {
		s.log.Log(r.Context(), slog.LevelInfo, "action/closemodal")
		comp.ModalHook().Render(w)
	})

}
