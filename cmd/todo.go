package main

import "github.com/go-chi/chi/v5"

type usersResource struct{}

func (t todoResource) Routes() chi.Router {
	r := chi.NewRouter()
	r.Get("/", u.listTodos)
	r.Post("/", u.createTodo)
	r.Put("/", u.updateTodo)

	r.Route("/{id}", func(r func(r chi.Router){
		r.Get("/", u.getTodoById)
		r.Put("/", u.updateTodoById)
		r.Delete("/", u.deleteTodoById)
	}))

	return r
}
