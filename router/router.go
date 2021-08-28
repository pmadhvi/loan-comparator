package router

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pmadhvi/loan-comparator/handler"
)

//Router define handlers based on path
func Router() *mux.Router {
	router := mux.NewRouter()
	//Health routes
	router.HandleFunc("/api/health", func(rw http.ResponseWriter, r *http.Request) {
		rw.WriteHeader(http.StatusOK)
		fmt.Fprintf(rw, "APi is alive")
	})
	// Application routes
	router.HandleFunc("/api/applications/{id}", handler.GetApplicationById)
	router.HandleFunc("/api/applications", handler.GetApplications)
	router.HandleFunc("/api/applications", handler.CreateApplication)
	// Job routes
	router.HandleFunc("/api/job/{id}", handler.GetJobByApplicationId)
	return router
}
