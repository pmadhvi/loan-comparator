package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
	_ "github.com/lib/pq"

	"github.com/pmadhvi/loan-comparator/model"
	log "github.com/sirupsen/logrus"
)

type applicationRepository struct {
	log log.Logger
	db  *sql.DB
}

func NewApplicationRepository(log log.Logger, db *sql.DB) *applicationRepository {
	return &applicationRepository{
		log: log,
		db:  db,
	}
}

func (ar applicationRepository) GetApplications(ctx context.Context) ([]model.Application, error) {
	var application model.Application

	rows, err := ar.db.QueryContext(ctx, "select * from applications")
	if err != nil {
		fmt.Printf("error while fetching all applications: %v", err)
		return nil, err
	}
	defer rows.Close()
	applications := make([]model.Application, 0)
	for rows.Next() {
		err := rows.Scan(application.ID, application.FirstName, application.LastName, application.Status)
		if err != nil {
			fmt.Printf("error while scaning row: ", err)
			return nil, err
		}
		applications = append(applications, application)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return applications, nil
}

func (ar applicationRepository) FindApplication(ctx context.Context, id uuid.UUID) (model.Application, error) {
	var application model.Application
	row := ar.db.QueryRowContext(ctx, "select * from applications where id=$1", id)
	err := row.Scan(application.ID, application.FirstName, application.LastName, application.Status)
	if err != nil {
		fmt.Printf("error while scaning row: ", err)
		return model.Application{}, err
	}
	return application, nil
}

func (ar applicationRepository) CreateApplication(ctx context.Context, application model.ApplicationRequest) error {
	id := uuid.New()
	_, err := ar.db.ExecContext(ctx,
		`INSERT INTO applications(id, first_name, last_name, status) values ($1, $2, $3, $4)`,
		id, "madhvi", "patel", "created", time.Now(), time.Now())
	if err != nil {
		fmt.Printf("error while creating application in db: %v", err)
		return err
	}
	return nil
}
