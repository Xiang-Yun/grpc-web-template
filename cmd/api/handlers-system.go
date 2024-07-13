package main

import "net/http"

func (app *application) PingApi(w http.ResponseWriter, r *http.Request) {
	// curl -X GET http://127.0.0.1:3500/api/ping -H "x-auth:0eafbf60-4374-4c79-b137-f9a1554c034e"
	app.writeJSON(w, http.StatusOK, StatusOK{Status: true, Data: "Pong"})
}
