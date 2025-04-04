package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	router.HandlerFunc(http.MethodPost, "/v1/todos", app.createTodoHandler)
	router.HandlerFunc(http.MethodPut, "/v1/todos/:id", app.updateTodoHandler)
	router.HandlerFunc(http.MethodDelete, "/v1/todos/:id", app.deleteTodoHandler)
	router.HandlerFunc(http.MethodGet, "/v1/todos", app.getTodosHandler)

	return router
}
