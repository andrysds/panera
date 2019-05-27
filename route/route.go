package route

import (
	"net/http"
	"os"
	"time"

	"github.com/andrysds/panera/handler"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

// Router return application http router
func Router() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	username = os.Getenv("USERNAME")
	password = os.Getenv("PASSWORD")
	r.Use(authorize)

	r.NotFound(handler.NotFound)

	r.Get("/", handler.Index)
	r.Get("/healthz", handler.Healthz)
	r.Get("/emulator", handler.Emulator)

	r.Route("/users", func(r chi.Router) {
		r.Get("/", handler.Users)
		r.Get("/new", handler.NewUser)
		r.Post("/", handler.CreateUser)
		r.Route("/{id}", func(r chi.Router) {
			r.Get("/add-to-standups", handler.AddToStandups)
			r.Get("/edit", handler.EditUser)
			r.Post("/", handler.UpdateUser)
			r.Get("/delete", handler.DeleteUser)
		})
	})

	r.Route("/standups", func(r chi.Router) {
		r.Get("/", handler.Standups)
		r.Route("/{id}", func(r chi.Router) {
			r.Get("/edit", handler.EditStandup)
			r.Post("/", handler.UpdateStandup)
			r.Get("/delete", handler.DeleteStandup)
		})
	})

	return r
}

var username, password string

func authorize(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user, pass, _ := r.BasicAuth()
		if username != user || password != pass {
			w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
			http.Error(w, "Unauthorized.", http.StatusUnauthorized)
			return
		}
		h.ServeHTTP(w, r)
	})
}
