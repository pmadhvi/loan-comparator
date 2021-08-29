package service

import (
	"context"
	"fmt"

	"github.com/pmadhvi/loan-comparator/model"
)

type jobRepo interface {
	FindJobByApplicationID(ctx context.Context, id int) (model.Job, error)
}

type JobService struct {
	JobRepository jobRepo
}

func (app JobService) GetJobByApplicationId(ctx context.Context, appID int) (model.Job, error) {
	fmt.Println("Inside GetJobByApplicationId service")
	job, _ := app.JobRepository.FindJobByApplicationID(ctx, appID)
	return job, nil
}
