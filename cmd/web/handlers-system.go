package main

import (
	"grpc-web-template/internal/common"
	"net/http"
	"time"
)

// Home displays the home page
func (app *application) Home(w http.ResponseWriter, r *http.Request) {
	if err := app.renderTemplate(w, r, "home", &templateData{}); err != nil {
		app.errorLog.Println(err)
	}
}

// AdminSystemMaintain shows system maintain
func (app *application) AdminSystemMaintain(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})
	data["buildTime"] = time.Now().Local().Format(common.TimeFormat)
	data["version"] = app.version

	if err := app.renderTemplate(w, r, "admin-system-maintain", &templateData{
		Data: data,
	}); err != nil {
		app.errorLog.Println("AdminSystemMaintain error:", err)
	}
}

// AllUsers shows the all users page
func (app *application) AllUsers(w http.ResponseWriter, r *http.Request) {
	if err := app.renderTemplate(w, r, "admin-all-users", &templateData{}); err != nil {
		app.errorLog.Println("AllUsers error:", err)
	}
}

// OneUser shows one admin user for add/edit/delete
func (app *application) OneUser(w http.ResponseWriter, r *http.Request) {
	if err := app.renderTemplate(w, r, "admin-one-user", &templateData{}); err != nil {
		app.errorLog.Println("OneUser error:", err)
	}
}

// AdminDashboard shows all dashboard in admin tool
func (app *application) AdminDashboard(w http.ResponseWriter, r *http.Request) {
	if err := app.renderTemplate(w, r, "admin-dashboard", &templateData{}); err != nil {
		app.errorLog.Println("AdminDashboard error:", err)
	}
}

// LoginPage displays the login page
func (app *application) LoginPage(w http.ResponseWriter, r *http.Request) {
	if err := app.renderTemplate(w, r, "login", &templateData{}); err != nil {
		app.errorLog.Println("LoginPage error:", err)
	}
}

// PostLoginPage handles the posted login form
func (app *application) PostLoginPage(w http.ResponseWriter, r *http.Request) {
	err := app.Session.RenewToken(r.Context())
	if err != nil {
		app.errorLog.Println("PostLoginPage RenewToken error:", err)
		return
	}

	err = r.ParseForm()
	if err != nil {
		app.errorLog.Println("PostLoginPage ParseForm error:", err)
		return
	}
	user := r.Form.Get("user")
	password := r.Form.Get("password")

	id, err := app.DB.Authenticate(user, password)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	app.Session.Put(r.Context(), "userID", id)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// Logout logs the user out
func (app *application) Logout(w http.ResponseWriter, r *http.Request) {
	app.Session.Destroy(r.Context())
	app.Session.RenewToken(r.Context())

	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
