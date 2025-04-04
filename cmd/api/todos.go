package main

import (
	"net/http"
)

func (app *application) createTodoHandler(w http.ResponseWriter, r *http.Request) {
	data := map[any]any{
		"id":          1,
		"title":       "Buy groceries",
		"description": "Buy milk, eggs, and bread",
	}

	err := app.writeJSON(w, nil, http.StatusOK, data)

	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

}
