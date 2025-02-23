package route

import (
	"fmt"
	"github.com/lpuig/cpmanager/html/comp"
	"github.com/lpuig/cpmanager/model/consultant"
	"github.com/lpuig/cpmanager/model/manager"
	"net/http"
)

// Open New Consultant Modal
func GetShowNewConsultantModal(m manager.Manager, w http.ResponseWriter, r *http.Request) {
	comp.AddConsultantModal().Render(w)
	m.Log.InfoContextWithTime(r.Context(), "open new consult modal")
}

// Retrieve New Consultant info from front Form. Return updated consultant list and trigger modal closing
func PostAddNewConsultantFromModal(m manager.Manager, w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	// Trig modal closing
	w.Header().Add("HX-Trigger-After-Swap", "closeModal")

	// Create new consultant
	newConsultant := &consultant.Consultant{}
	newConsultant.Clean(
		r.FormValue("FirstName"),
		r.FormValue("LastName"),
		r.FormValue("Profile"),
		r.FormValue("CrmrId"),
	)
	m.Consultants.Add(newConsultant)
	comp.ConsultantTable(m.Consultants).Render(w)
	m.Log.InfoContextWithTime(r.Context(), fmt.Sprintf("add new consult (%s)", newConsultant.Name()))
}

// Open Update Consultant Modal for consultant with url given id
func GetShowUpdateConsultantModal(m manager.Manager, w http.ResponseWriter, r *http.Request) {
	consultId := r.PathValue("id")
	consult, found := m.Consultants.Get(consultId)
	if !found {
		w.WriteHeader(http.StatusNotFound)
		m.Log.ErrorContextWithTime(r.Context(), "open update consult modal failed: unknown id")
		return
	}

	comp.UpdateConsultantModal(consult).Render(w)
	m.Log.InfoContextWithTime(r.Context(), fmt.Sprintf("open update consult modal for (%s)", consult.Name()))
}

// Retrieve New Consultant info from front Form. Return updated consultant row and trigger modal closing
func PostUpdateConsultantFromModal(m manager.Manager, w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	consultId := r.PathValue("id")
	consult, found := m.Consultants.Get(consultId)
	if !found {
		w.WriteHeader(http.StatusNotFound)
		m.Log.ErrorContextWithTime(r.Context(), "update consult failed: unknown id")
		return
	}

	// Trig modal closing
	w.Header().Add("HX-Trigger-After-Swap", "closeModal")

	// update consultant
	consult.Clean(
		r.FormValue("FirstName"),
		r.FormValue("LastName"),
		r.FormValue("Profile"),
		r.FormValue("CrmrId"),
	)
	m.Consultants.Update(consult)
	comp.ConsultantTableRow(consult).Render(w)
	m.Log.InfoContextWithTime(r.Context(), fmt.Sprintf("update consult (%s)", consult.Name()))
}

// Delete Consultant
func DeleteConsultant(m manager.Manager, w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	consultId := r.PathValue("id")
	consult, found := m.Consultants.Get(consultId)
	if !found {
		w.WriteHeader(http.StatusNotFound)
		m.Log.ErrorContextWithTime(r.Context(), "delete consult failed: unknown id")
		return
	}

	// delete consultant
	m.Consultants.Remove(consult)
	m.Log.InfoContextWithTime(r.Context(), fmt.Sprintf("delete consult (%s)", consult.Name()))
}
