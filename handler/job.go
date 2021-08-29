package handler

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/pmadhvi/loan-comparator/model"
)

type jobService interface {
	GetJobByApplicationId(ctx context.Context, id int) (model.Job, error)
}

func (r Router) getJobByApplicationId(rw http.ResponseWriter, req *http.Request) {
	fmt.Println("Inside getJobByApplicationId")
	ctx := req.Context()
	id, _ := strconv.Atoi(req.URL.Query().Get("id"))
	job, _ := r.JobService.GetJobByApplicationId(ctx, id)
	fmt.Fprintf(rw, "Inside getJobByApplicationId: %v ", job)
}
