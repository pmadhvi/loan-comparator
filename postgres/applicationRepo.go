package postgres

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/pmadhvi/loan-comparator/model"
	log "github.com/sirupsen/logrus"
)

type ApplicationRepository struct {
	Log log.Logger
	Db  *sql.DB
}

func (ar ApplicationRepository) FindApplication(ctx context.Context, id int) (model.Application, error) {
	fmt.Println("Impemented FindApplication")
	return model.Application{}, nil
}

func (ar ApplicationRepository) GetApplications(ctx context.Context) ([]model.Application, error) {
	fmt.Println("Impemented GetApplications")
	return nil, nil
}

func (ar ApplicationRepository) CreateApplication(ctx context.Context, application model.ApplicationRequest) error {
	fmt.Println("created application to db")
	return nil

}
