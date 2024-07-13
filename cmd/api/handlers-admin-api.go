package main

import (
	"errors"
	"fmt"
	"grpc-web-template/internal/models"
	"grpc-web-template/internal/repository/sqliteRepo"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
	"golang.org/x/crypto/bcrypt"
)

// CreateAuthToken creates and sends an auth token, if user supplies valid information
func (app *application) CreateAuthToken(w http.ResponseWriter, r *http.Request) {
	var userInput struct {
		User     string `json:"user"`
		Password string `json:"password"`
	}

	err := app.readJSON(w, r, &userInput)
	if err != nil {
		app.badRequest(w, err)
		return
	}

	// get the user from the database by email; send error if invalid email
	user, err := app.DB.GetUserByUser(userInput.User)
	if err != nil {
		app.invalidCredentials(w)
		return
	}

	// validate the password; send error if invalid password
	validPassword, err := app.passwordMatches(user.Password, userInput.Password)
	if err != nil {
		app.invalidCredentials(w)
		return
	}

	if !validPassword {
		app.invalidCredentials(w)
		return
	}

	// generate the token
	token, err := sqliteRepo.GenerateToken(user.ID, 24*time.Hour, sqliteRepo.ScopeAuthentication)
	if err != nil {
		app.badRequest(w, err)
		return
	}

	// save to database
	err = app.DB.AddToken(token, user)
	if err != nil {
		app.badRequest(w, err)
		return
	}

	// send response
	var payload struct {
		Status  bool          `json:"status"`
		Message string        `json:"message"`
		Token   *models.Token `json:"authentication_token"`
	}
	payload.Status = true
	payload.Message = fmt.Sprintf("token for %s created", userInput.User)
	payload.Token = token

	_ = app.writeJSON(w, http.StatusOK, payload)
}

// authenticateToken checks an auth token for validity
func (app *application) authenticateToken(r *http.Request) (*models.User, error) {
	authorizationHeader := r.Header.Get("Authorization")
	if authorizationHeader == "" {
		return nil, errors.New("no authorization header received")
	}

	headerParts := strings.Split(authorizationHeader, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		return nil, errors.New("no authorization header received")
	}

	token := headerParts[1]
	if len(token) != 26 {
		return nil, errors.New("authentication token wrong size")
	}

	// get the user from the tokens table
	user, err := app.DB.GetUserForToken(token)
	if err != nil {
		return nil, errors.New("no matching user found")
	}

	return user, nil
}

// CheckAuthentication checks auth status
func (app *application) CheckAuthentication(w http.ResponseWriter, r *http.Request) {
	// validate the token, and get associated user
	user, err := app.authenticateToken(r)
	if err != nil {
		app.invalidCredentials(w)
		return
	}

	// valid user
	message := fmt.Sprintf("authenticated user %s", user.User)
	app.writeJSON(w, http.StatusOK, StatusOK{Status: true, Data: message})
}

// AllUsers returns a JSON file listing all admin users
func (app *application) AllUsers(w http.ResponseWriter, r *http.Request) {
	allUsers, err := app.DB.GetAllUsers()
	if err != nil {
		app.badRequest(w, err)
		return
	}
	app.writeJSON(w, http.StatusOK, StatusOK{Status: true, Data: allUsers})
}

// OneUser gets one user by id (from the url) and returns it as JSON
func (app *application) OneUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	userID, _ := strconv.Atoi(id)

	user, err := app.DB.GetOneUser(userID)
	if err != nil {
		app.badRequest(w, err)
		return
	}
	app.writeJSON(w, http.StatusOK, StatusOK{Status: true, Data: user})
}

// EditUser is the handler for adding or editing an existing user
func (app *application) EditUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	userID, _ := strconv.Atoi(id)

	var user models.User

	err := app.readJSON(w, r, &user)
	if err != nil {
		app.badRequest(w, err)
		return
	}

	if userID > 0 {
		err = app.DB.EditUser(user)
		if err != nil {
			app.badRequest(w, err)
			return
		}

		if user.Password != "" {
			newHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 12)
			if err != nil {
				app.badRequest(w, err)
				return
			}

			err = app.DB.UpdatePasswordForUser(user, string(newHash))
			if err != nil {
				app.badRequest(w, err)
				return
			}
		}
	} else {
		newHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 12)
		if err != nil {
			app.badRequest(w, err)
			return
		}
		err = app.DB.AddUser(user, string(newHash))
		if err != nil {
			app.badRequest(w, err)
			return
		}
	}
	app.writeJSON(w, http.StatusOK, StatusOK{Status: true})
}

// DeleteUser deletes a user, and all associated tokens, from the database
func (app *application) DeleteUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "0" {
		app.badRequest(w, errors.New("error delete admin user"))
		return
	}
	userID, _ := strconv.Atoi(id)

	err := app.DB.DeleteUser(userID)
	if err != nil {
		app.badRequest(w, err)
		return
	}
	app.writeJSON(w, http.StatusOK, StatusOK{Status: true})
}
