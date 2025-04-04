package main

import (
	"net/http"
	"time"
)

type ErrorResponse struct {
	Error     string `json:"error"`
	Message   string `json:"message"`
	Timestamp string `json:"timestamp"`
}

func (app *application) logError(r *http.Request, err error) {
	app.logger.Error(err.Error(), "method", r.Method, "uri", r.URL.RequestURI())
}

func (app *application) errorResponse(w http.ResponseWriter, r *http.Request, status int, res ErrorResponse) {
	err := app.writeJSON(w, nil, status, res)
	if err != nil {
		app.logError(r, err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (app *application) serverErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.logError(r, err)
	res := ErrorResponse{
		Error:     "Internal Server Error",
		Message:   "An unexpected error occurred. Please try again later.",
		Timestamp: time.Now().UTC().Format(time.RFC3339),
	}
	app.errorResponse(w, r, http.StatusInternalServerError, res)
}
