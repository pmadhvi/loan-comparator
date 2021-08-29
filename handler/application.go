package handler

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/pmadhvi/loan-comparator/model"
)

type applicationService interface {
	CreateApplication(ctx context.Context, application model.ApplicationRequest) error
	GetApplicationById(ctx context.Context, id int) (model.Application, error)
	GetApplications(ctx context.Context) ([]model.Application, error)
}

func (r Router) getApplicationById(rw http.ResponseWriter, req *http.Request) {
	fmt.Println("Inside getApplicationById")
	ctx := req.Context()
	id, _ := strconv.Atoi(req.URL.Query().Get("id"))
	application, _ := r.ApplicationService.GetApplicationById(ctx, id)
	fmt.Fprintf(rw, "Inside getApplicationById: %v", application)
}

func (r Router) getApplications(rw http.ResponseWriter, req *http.Request) {
	fmt.Println("Inside GetApplications")
	ctx := req.Context()
	applications, _ := r.ApplicationService.GetApplications(ctx)
	fmt.Fprintf(rw, "Inside GetApplications: %v", applications)
}

func (r Router) createApplication(rw http.ResponseWriter, req *http.Request) {
	fmt.Println("Inside CreateApplication")
	ctx := req.Context()
	application := model.ApplicationRequest{}
	_ = r.ApplicationService.CreateApplication(ctx, application)
	fmt.Fprint(rw, "Inside CreateApplication")
}
