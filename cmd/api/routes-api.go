package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func (app *application) routes() http.Handler {
	mux := chi.NewRouter()

	// mux.Use(middleware.Logger)

	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*", "http://*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{
			"*",
		},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	mux.Route("/api", func(mux chi.Router) {
		mux.Get("/ping", app.PingApi)
		mux.Post("/authenticate", app.CreateAuthToken)

		mux.Route("/admin", func(mux chi.Router) {
			mux.Post("/all-users", app.AllUsers)
			mux.Post("/all-users/{id}", app.OneUser)
			mux.Post("/all-users/edit/{id}", app.EditUser)
			mux.Post("/all-users/delete/{id}", app.DeleteUser)
		})
	})

	return mux
}
