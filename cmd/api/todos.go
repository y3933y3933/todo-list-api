package main

import (
	"encoding/json"
	"net/http"
)

func (app *application) createTodoHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]any{
		"id":          1,
		"title":       "Buy groceries",
		"description": "Buy milk, eggs, and bread",
	}

	js, err := json.Marshal(data)
	if err != nil {
		app.logger.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("something went wrong"))
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
