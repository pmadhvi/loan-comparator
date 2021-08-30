package handler

import (
	"context"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/pmadhvi/loan-comparator/model"
)

type jobService interface {
	GetJobByApplicationId(ctx context.Context, id uuid.UUID) (model.Job, error)
}

func (r Router) getJobByApplicationId(rw http.ResponseWriter, req *http.Request) {
	fmt.Println("Inside getJobByApplicationId")
	ctx := req.Context()
	id, _ := uuid.Parse(req.URL.Query().Get("id"))
	job, err := r.JobService.GetJobByApplicationId(ctx, id)
	if err != nil {
		fmt.Printf("error while fetching job with application_id: %d %v", id, err)
	}
	fmt.Fprintf(rw, "Job with application_id %d: %v ", id, job)
}
