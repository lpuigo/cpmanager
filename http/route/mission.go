package route

import (
	"fmt"
	"github.com/lpuig/cpmanager/html/comp"
	"github.com/lpuig/cpmanager/model/consultant"
	"github.com/lpuig/cpmanager/model/manager"
	"net/http"
)

// Open New Mission Modal
func GetShowAddMissionModal(m manager.Manager, w http.ResponseWriter, r *http.Request) {
	consultId := r.PathValue("id")
	consult, found := m.Consultants.Get(consultId)
	if !found {
		w.WriteHeader(http.StatusNotFound)
		m.Log.ErrorContextWithTime(r.Context(), fmt.Sprintf("open add mission modal failed: unknown consultant id %s", consultId))
		return
	}

	comp.AddMissionModal(consult).Render(w)
	m.Log.InfoContextWithTime(r.Context(), fmt.Sprintf("open add mission modal modal for consultant (%s)", consult.Name()))
}

// Retrieve New Mission info from front Form. Return updated consultant row and trigger modal closing
func PostAddMissionFromModal(m manager.Manager, w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	consultId := r.PathValue("id")
	consult, found := m.Consultants.Get(consultId)
	if !found {
		w.WriteHeader(http.StatusNotFound)
		m.Log.ErrorContextWithTime(r.Context(), fmt.Sprintf("add mission failed: unknown consultant id %s", consultId))
		return
	}

	// Trig modal closing
	w.Header().Add("HX-Trigger-After-Swap", "closeModal")

	// create mission
	mission := consultant.NewCleanMission(
		r.FormValue("Title"),
		r.FormValue("Company"),
		r.FormValue("Manager"),
		r.FormValue("StartDay"),
		r.FormValue("EndDay"),
		r.FormValue("DailyCost"),
		r.FormValue("DailyRate"),
	)

	// add it to pertaining consultant
	consult.AddMission(mission)
	// update consultant
	m.Consultants.Update(consult)
	// return update consultant row
	comp.ConsultantTableRow(consult).Render(w)
	m.Log.InfoContextWithTime(r.Context(), fmt.Sprintf("add mission for consultant (%s)", consult.Name()))
}
