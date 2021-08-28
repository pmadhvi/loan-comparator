package router

import (
	"github.com/gorilla/mux"
)

//Router define handlers based on path
func Router() *mux.Router {
	router := mux.NewRouter()
	//router.HandleFunc("/api/health", handler.HealthHandler)
	router.HandleFunc("/api/application/{id}", handler.GetApplicationById)
	return router
}
