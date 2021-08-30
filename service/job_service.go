package service

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/pmadhvi/loan-comparator/model"
)

type jobRepo interface {
	FindJobByApplicationID(ctx context.Context, id uuid.UUID) (model.Job, error)
	//UpdateJobByApplicationID(ctx context.Context, id uuid.UUID, status string) error
	//CreateJob(ctx context.Context, application_id uuid.UUID)
}

type JobService struct {
	JobRepository jobRepo
}

func (job JobService) GetJobByApplicationId(ctx context.Context, appID uuid.UUID) (model.Job, error) {
	fmt.Println("Inside GetJobByApplicationId service")
	appJob, err := job.JobRepository.FindJobByApplicationID(ctx, appID)
	if err != nil {
		fmt.Println(err)
		return model.Job{}, err
	}
	return appJob, nil
}

// func (job JobService) UpdateApplicationStatus(ctx context.Context, appID uuid.UUID, status string) error {
// 	fmt.Println("Inside UpdateApplicationStatus service")
// 	err := job.JobRepository.UpdateJobByApplicationID(ctx, appID, status)
// 	if err != nil {
// 		fmt.Println(err)
// 		return err
// 	}
// 	return nil
// }
