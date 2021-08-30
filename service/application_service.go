package service

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/pmadhvi/loan-comparator/model"
)

type applicationRepo interface {
	CreateApplication(ctx context.Context, application model.ApplicationRequest) error
	FindApplication(ctx context.Context, id uuid.UUID) (model.Application, error)
	GetApplications(ctx context.Context) ([]model.Application, error)
}

type ApplicationService struct {
	ApplicationRepository applicationRepo
}

func (app ApplicationService) GetApplicationById(ctx context.Context, appID uuid.UUID) (model.Application, error) {
	fmt.Println("Inside GetApplicationById service")
	application, err := app.ApplicationRepository.FindApplication(ctx, appID)
	if err != nil {
		fmt.Println(err)
		return model.Application{}, err
	}
	return application, nil
}

func (app ApplicationService) GetApplications(ctx context.Context) ([]model.Application, error) {
	fmt.Println("Inside GetApplications service")
	applications, err := app.ApplicationRepository.GetApplications(ctx)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return applications, nil
}

func (app ApplicationService) CreateApplication(ctx context.Context, data model.ApplicationRequest) error {
	fmt.Println("Inside CreateApplication service")
	err := app.ApplicationRepository.CreateApplication(ctx, data)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
