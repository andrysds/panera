package handler

import (
	"net/http"

	"github.com/andrysds/panera/entity"
	"github.com/andrysds/panera/template"
	"github.com/globalsign/mgo/bson"
	"github.com/go-chi/chi"
)

// Standups is handler function for GET /standups
func Standups(w http.ResponseWriter, r *http.Request) {
	var standups []entity.Standup
	err := entity.Standups.All("state", &standups)
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

// EditStandup is handler function for GET /standups/:id
func EditStandup(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var standup entity.Standup
	entity.Standups.FindOne(id, &standup)
	user, _ := standup.User()
	data := struct {
		templateData
		Standup entity.Standup
		User    *entity.User
	}{
		templateData: templateData{
			PageTitle:  "Edit Standup",
			FormAction: "/standups/" + id,
		},
		Standup: standup,
		User:    user,
	}
	template.Execute(w, "standup-form.html", data)
}

// UpdateStandup is handler function for POST /standups/:id
func UpdateStandup(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	id := chi.URLParam(r, "id")
	standup := entity.Standup{
		ID:     bson.ObjectIdHex(id),
		UserID: bson.ObjectIdHex(r.Form.Get("user_id")),
		State:  r.Form.Get("state"),
	}
	err := entity.Standups.UpdateOne(id, standup)
	afterStandupAction(w, r, "update-standup", err)
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
