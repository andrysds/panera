package route

import (
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
			r.Get("/set-done", handler.SetDone)
			r.Get("/delete", handler.DeleteStandup)
		})
	})

	return r
}
