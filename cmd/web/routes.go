package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
)

func (app *application) routes() http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(SessionLoad)

	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*", "http://*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "X-AUTH", "Authorization", "Content-Type", "X-CSRF-Token", "Access-Control-Allow-Headers", "Origin", "Accept", "X-Requested-With"},
		ExposedHeaders: []string{"Access-Control-Allow-Origin"},
	}))

	// root
	mux.Get("/", app.Home)

	// websocket for refresh web component
	mux.Route("/ws", func(mux chi.Router) {
		mux.Get("/endPoint", app.WsEndPoint)
	})

	mux.Route("/admin", func(mux chi.Router) {
		mux.Get("/system/maintain", app.AdminSystemMaintain)
		mux.Get("/dashboard", app.AdminDashboard)

		mux.Get("/all-users", app.AllUsers)
		mux.Get("/all-users/{id}", app.OneUser)
	})

	// auth routes
	mux.Get("/login", app.LoginPage)
	mux.Post("/login", app.PostLoginPage)
	mux.Get("/logout", app.Logout)

	// file server
	fileServer := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
