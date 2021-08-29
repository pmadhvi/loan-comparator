package service

import (
	"context"
	"fmt"

	"github.com/pmadhvi/loan-comparator/model"
)

type applicationRepo interface {
	CreateApplication(ctx context.Context, application model.ApplicationRequest) error
	FindApplication(ctx context.Context, id int) (model.Application, error)
	GetApplications(ctx context.Context) ([]model.Application, error)
}

type ApplicationService struct {
	ApplicationRepository applicationRepo
}

func (app ApplicationService) GetApplicationById(ctx context.Context, appID int) (model.Application, error) {
	fmt.Println("Inside GetApplicationById service")
	application, _ := app.ApplicationRepository.FindApplication(ctx, appID)
	return application, nil
}

func (app ApplicationService) GetApplications(ctx context.Context) ([]model.Application, error) {
	fmt.Println("Inside GetApplications service")
	applications, _ := app.ApplicationRepository.GetApplications(ctx)
	return applications, nil
}

func (app ApplicationService) CreateApplication(ctx context.Context, data model.ApplicationRequest) error {
	fmt.Println("Inside CreateApplication service")
	_ = app.ApplicationRepository.CreateApplication(ctx, data)
	return nil
}
