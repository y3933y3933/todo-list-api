package main

import "net/http"

func (app *application) registerHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]any{
		"token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9",
	}

	err := app.writeJSON(w, nil, http.StatusOK, data)

	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
}

func (app *application) loginHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]any{
		"token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9",
	}

	err := app.writeJSON(w, nil, http.StatusOK, data)

	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
}
