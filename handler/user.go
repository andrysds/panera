package handler

import (
	"net/http"

	"github.com/andrysds/panera/entity"
	"github.com/andrysds/panera/template"
	"github.com/globalsign/mgo/bson"
	"github.com/go-chi/chi"
)

// Users is handler function for GET /users
func Users(w http.ResponseWriter, r *http.Request) {
	users, err := entity.AllUsers()
	if err != nil {
		internalServerError(w, err)
	} else {
		tpltData := templateData{PageTitle: "Users"}
		tpltData.setPastActionInfo(r)
		data := struct {
			templateData
			Users []*entity.User
		}{
			templateData: tpltData,
			Users:        users,
		}
		template.Execute(w, "user.html", data)
	}
}

// NewUser is handler function for GET /users/new
func NewUser(w http.ResponseWriter, r *http.Request) {
	data := struct {
		templateData
		User entity.User
	}{
		templateData: templateData{
			PageTitle:  "Create User",
			FormAction: "/users",
		},
	}
	template.Execute(w, "user-form.html", data)
}

// CreateUser is handler function for POST /users
func CreateUser(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	newUser := entity.User{
		ID:       bson.NewObjectId(),
		Name:     r.Form.Get("name"),
		Username: r.Form.Get("username"),
		Birthday: r.Form.Get("birthday"),
	}
	err := entity.Users.InsertOne(newUser)
	afterUserAction(w, r, "create-user", err)
}

// EditUser is handler function for GET /users/:id
func EditUser(w http.ResponseWriter, r *http.Request) {
	var user entity.User
	id := chi.URLParam(r, "id")
	entity.Users.FindOne(id, &user)
	data := struct {
		templateData
		User entity.User
	}{
		templateData: templateData{
			PageTitle:  "Edit User",
			FormAction: "/users/" + id,
		},
		User: user,
	}
	template.Execute(w, "user-form.html", data)
}

// UpdateUser is handler function for POST /users/:id
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	id := chi.URLParam(r, "id")
	user := entity.User{
		ID:       bson.ObjectIdHex(id),
		Name:     r.Form.Get("name"),
		Username: r.Form.Get("username"),
		Birthday: r.Form.Get("birthday"),
	}
	err := entity.Users.UpdateOne(id, user)
	afterUserAction(w, r, "update-user", err)
}

// DeleteUser is handler function for GET /users/:id/delete
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	err := entity.Users.RemoveOne(id)
	if err == nil {
		_, err = entity.Standups.RemoveAll(bson.M{"user_id": bson.ObjectIdHex(id)})
	}
	afterUserAction(w, r, "delete-user", err)
}

// AddToStandups adds new standups with given user ID
func AddToStandups(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	err := entity.AddUserToStandups(id)
	afterUserAction(w, r, "add-to-standups", err)
}

func afterUserAction(w http.ResponseWriter, r *http.Request, action string, err error) {
	afterAction(w, r, afterActionOpts{action, err, "/users"})
}
