package route

import (
	"fmt"
	"github.com/lpuig/cpmanager/html/comp"
	"github.com/lpuig/cpmanager/model/consultant"
	"github.com/lpuig/cpmanager/model/manager"
	"log/slog"
	"net/http"
)

// Open New Consultant Modal
func GetShowNewConsultantModal(m *manager.Manager, w http.ResponseWriter, r *http.Request) {
	m.Log.InfoContext(r.Context(), "open new consult modal")
	comp.AddConsultantModal().Render(w)
}

// Retrieve New Consultant info from front Form. Return updated consultant list and trigger modal closing
func PostAddNewConsultantFromModal(m *manager.Manager, w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	// Trig modal closing
	w.Header().Add("HX-Trigger-After-Swap", "closeModal")

	// Create new consultant
	newConsultant := consultant.Consultant{
		FirstName: r.FormValue("FirstName"),
		LastName:  r.FormValue("LastName"),
	}
	id := m.Consultants.Add(newConsultant)
	m.Log.InfoContext(r.Context(), fmt.Sprintf("add new consult with id %s (%s)", id, newConsultant.Name()))
	comp.ConsultantsList(m.Consultants).Render(w)
}

// Open Update Consultant Modal for consultant with url given id
func GetShowUpdateConsultantModal(m *manager.Manager, w http.ResponseWriter, r *http.Request) {
	consultId := r.PathValue("id")
	consult, found := m.Consultants.Get(consultId)
	if !found {
		m.Log.ErrorContext(r.Context(), fmt.Sprintf("open update consult modal failed: unknown id %s", consultId))
		w.WriteHeader(http.StatusNotFound)
		return
	}

	m.Log.Log(r.Context(), slog.LevelInfo, fmt.Sprintf("open update consult modal for id %s (%s)", consultId, consult.Name()))
	comp.UpdateConsultantModal(consult).Render(w)
}

// Retrieve New Consultant info from front Form. Return updated consultant list and trigger modal closing
func PostUpdateConsultantFromModal(m *manager.Manager, w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	consultId := r.PathValue("id")
	consult, found := m.Consultants.Get(consultId)
	if !found {
		m.Log.ErrorContext(r.Context(), fmt.Sprintf("update consult failed: unknown id %s", consultId))
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// Trig modal closing
	w.Header().Add("HX-Trigger-After-Swap", "closeModal")

	// update consultant
	consult.FirstName = r.FormValue("FirstName")
	consult.LastName = r.FormValue("LastName")
	m.Consultants.Update(consult)
	m.Log.InfoContext(r.Context(), fmt.Sprintf("update consult with id %s (%s)", consultId, consult.Name()))
	comp.ConsultantLine(consult).Render(w)

}
