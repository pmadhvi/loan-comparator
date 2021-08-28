package main

import "github.com/go-chi/chi/v5"

type usersResource struct{}

func (u usersResource) Routes() chi.Router {

	r := chi.NewRouter()
	r.Get("/", u.listUsers)
	return r
}
