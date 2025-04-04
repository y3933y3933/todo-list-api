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

func (app *application) updateTodoHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]any{
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

func (app *application) deleteTodoHandler(w http.ResponseWriter, r *http.Request) {

	err := app.writeJSON(w, nil, http.StatusNoContent, nil)

	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
}

func (app *application) getTodosHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]any{
		"data": []map[string]any{
			{
				"id":          1,
				"title":       "Buy groceries",
				"description": "Buy milk, eggs, bread",
			},
			{
				"id":          2,
				"title":       "Pay bills",
				"description": "Pay electricity and water bills",
			},
		},
		"page":  1,
		"limit": 10,
		"total": 2,
	}

	err := app.writeJSON(w, nil, http.StatusOK, data)

	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
}
