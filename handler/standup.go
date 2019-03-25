package handler

import (
	"net/http"

	"github.com/andrysds/panera/entity"
	"github.com/andrysds/panera/template"
	"github.com/go-chi/chi"
)

// Standups is handler function for GET /standups
func Standups(w http.ResponseWriter, r *http.Request) {
	var standups []entity.Standup
	err := entity.Standups.All("timestamp", &standups)
	if err != nil {
		internalServerError(w, err)
	} else {
		tpltData := templateData{PageTitle: "Standups"}
		tpltData.setPastActionInfo(r)
		data := struct {
			templateData
			Standups []entity.Standup
		}{
			templateData: tpltData,
			Standups:     standups,
		}
		template.Execute(w, "standup.html", data)
	}
}

// SetDone is handler function for POST /standups/:id
func SetDone(w http.ResponseWriter, r *http.Request) {
	var standup entity.Standup
	id := chi.URLParam(r, "id")
	err := entity.Standups.FindOne(id, &standup)
	if err == nil {
		err = standup.SetDone()
	}
	afterStandupAction(w, r, "set-done", err)
}

// DeleteStandup is handler function for GET /standups/:id/delete
func DeleteStandup(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	err := entity.Standups.RemoveOne(id)
	afterStandupAction(w, r, "delete-standup", err)
}

func afterStandupAction(w http.ResponseWriter, r *http.Request, action string, err error) {
	afterAction(w, r, afterActionOpts{action, err, "/standups"})
}
