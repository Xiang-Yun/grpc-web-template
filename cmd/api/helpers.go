package main

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

type StatusOK struct {
	Status bool        `json:"status"`
	Data   interface{} `json:"data"`
}

type StatusFailed struct {
	Status bool              `json:"status"`
	Code   string            `json:"code"`
	Data   interface{}       `json:"data"`
	Errors map[string]string `json:"errors"`
}

// writeJSON writes aribtrary data out as JSON
func (app *application) writeJSON(w http.ResponseWriter, status int, data interface{}, headers ...http.Header) error {
	out, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return err
	}

	if len(headers) > 0 {
		for k, v := range headers[0] {
			w.Header()[k] = v
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(out)
	return nil
}

// readJSON reads json from request body into data. We only accept a single json value in the body
func (app *application) readJSON(w http.ResponseWriter, r *http.Request, data interface{}) error {
	maxBytes := 4194304 // max 4 megabyte in request body
	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))

	dec := json.NewDecoder(r.Body)
	err := dec.Decode(data)
	if err != nil {
		return err
	}

	// we only allow one entry in the json file
	err = dec.Decode(&struct{}{})
	if err != io.EOF {
		return errors.New("body must only have a single JSON value")
	}
	return nil
}

// badRequest sends a JSON response with status http.StatusBadRequest, describing the error
func (app *application) badRequest(w http.ResponseWriter, err error) error {
	payload := StatusFailed{
		Status: false,
		Data:   err.Error(),
		Code:   "ER_UNKNOWN",
	}

	out, err := json.MarshalIndent(payload, "", "\t")
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	w.Write(out)
	return nil
}

func (app *application) invalidCredentials(w http.ResponseWriter) error {
	payload := StatusFailed{
		Status: false,
		Data:   "invalid authentication credentials",
		Code:   "AUTH_ERROR",
	}

	err := app.writeJSON(w, http.StatusUnauthorized, payload)
	if err != nil {
		return err
	}
	return nil
}

func (app *application) passwordMatches(hash, password string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		switch {
		case errors.Is(err, bcrypt.ErrMismatchedHashAndPassword):
			return false, nil
		default:
			return false, err
		}
	}
	return true, nil
}
