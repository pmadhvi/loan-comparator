package handler

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

type Router struct {
	Log                log.Logger
	ApplicationService applicationService
	JobService         jobService
}

//Router define handlers based on path
func (r Router) Router() *mux.Router {
	router := mux.NewRouter()
	//Health routes
	router.HandleFunc("/api/health", func(rw http.ResponseWriter, r *http.Request) {
		rw.WriteHeader(http.StatusOK)
		fmt.Fprintf(rw, "APi is alive")
	})
	// Application routes
	router.HandleFunc("/api/applications/{id}", r.getApplicationById)
	router.HandleFunc("/api/applications", r.getApplications)
	router.HandleFunc("/api/application", r.createApplication)
	// Job routes
	router.HandleFunc("/api/job/{id}", r.getJobByApplicationId)
	return router
}
