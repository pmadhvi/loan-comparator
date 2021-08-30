package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/google/uuid"
	"github.com/pmadhvi/loan-comparator/model"
)

type applicationService interface {
	CreateApplication(ctx context.Context, application model.ApplicationRequest) error
	GetApplicationById(ctx context.Context, id uuid.UUID) (model.Application, error)
	GetApplications(ctx context.Context) ([]model.Application, error)
}

func (r Router) getApplicationById(rw http.ResponseWriter, req *http.Request) {
	fmt.Println("Inside getApplicationById")
	ctx := req.Context()
	id, _ := uuid.Parse(req.URL.Query().Get("id"))
	application, err := r.ApplicationService.GetApplicationById(ctx, id)
	if err != nil {
		fmt.Printf("error while fetching application with id: %d %v", id, err)
	}
	fmt.Fprintf(rw, "Application with application_id %d: %v", application)
}

func (r Router) getApplications(rw http.ResponseWriter, req *http.Request) {
	fmt.Println("Inside GetApplications")
	ctx := req.Context()
	applications, err := r.ApplicationService.GetApplications(ctx)
	if err != nil {
		fmt.Printf("error while fetching applications %v:", err)
	}
	fmt.Fprintf(rw, "Applications: %v", applications)
}

func (r Router) createApplication(rw http.ResponseWriter, req *http.Request) {
	fmt.Println("Inside CreateApplication")
	ctx := req.Context()
	var application model.ApplicationRequest
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Printf("error while reading the request body: %v", err)
	}
	defer req.Body.Close()
	err = json.Unmarshal(body, &application)
	if err != nil {
		fmt.Printf("unmarshalling the request body failed: %v", err)
	}
	err = r.ApplicationService.CreateApplication(ctx, application)
	if err != nil {
		fmt.Printf("Application could not be created due to: %v", err)
	}
	fmt.Fprint(rw, "Inside CreateApplication")
}
