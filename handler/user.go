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
	users, err := entity.Users().All()
	if err != nil {
		internalServerError(w, err)
	} else {
		data := struct {
			templateData
			Users []entity.User
		}{
			templateData: templateData{PageTitle: "Users"},
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
	new := entity.User{
		ID:       bson.NewObjectId(),
		Name:     r.Form.Get("name"),
		Username: r.Form.Get("username"),
		Birthday: r.Form.Get("birthday"),
		Role:     r.Form.Get("role"),
	}
	err := entity.Users().InsertOne(new)
	afterUserAction(w, r, err)
}

// EditUser is handler function for GET /users/:id
func EditUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	user, _ := entity.Users().FindOne(id)
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
		Role:     r.Form.Get("role"),
	}
	err := entity.Users().UpdateOne(id, user)
	afterUserAction(w, r, err)
}

// DeleteUser is handler function for GET /users/:id/delete
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	err := entity.Users().RemoveOne(id)
	afterUserAction(w, r, err)
}

func afterUserAction(w http.ResponseWriter, r *http.Request, err error) {
	afterAction(w, r, err, "/users")
}
