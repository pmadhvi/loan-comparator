package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"github.com/pmadhvi/loan-comparator/model"
	log "github.com/sirupsen/logrus"
)

type jobRepository struct {
	log log.Logger
	db  *sql.DB
}

func NewJobRepository(log log.Logger, db *sql.DB) *jobRepository {
	return &jobRepository{
		log: log,
		db:  db,
	}
}

// func (jr jobRepository) CreateJob(ctx context.Context, application_id uuid.UUID) error {
// 	_, err := Db.ExecContext(ctx, `insert into jobs(application_id, status) values (application_id, "created")`)
// 	if err != nil {
// 		fmt.Printf("error while creating job for application_id %d in db: %v", application_id, err)
// 		return err
// 	}
// 	return nil
// }

func (jr jobRepository) FindJobByApplicationID(ctx context.Context, id uuid.UUID) (model.Job, error) {
	var job model.Job
	row := jr.db.QueryRowContext(ctx, "select * from jobs where application_id=$1", id)
	err := row.Scan(job.ID, job.ApplicationID, job.Status)
	if err != nil {
		fmt.Printf("error while scaning row: %v", err)
	}
	return job, nil
}

// func (jr JobRepository) UpdateJobByApplicationID(ctx context.Context, id uuid.UUID, status string) error {
// 	var job model.Job
// 	stmt, err := Db.PrepareContext(ctx, "update jobs set status=$1 where application_id=$2")
// 	if err != nil {
// 		fmt.Printf("error while creating prepare statement %v: ", err)
// 	}
// 	res, err := stmt.ExecContext(ctx, status, id)
// 	if err != nil {
// 		fmt.Printf("error while executing update statement %v: ", err)
// 	}
// 	_, err = res.RowsAffected()
// 	if err != nil {
// 		fmt.Printf("error while fetching updated row %v: ", err)
// 	}
// 	return nil
// }

// func fetchAllApplications(ctx context.Context) ([]model.Job, error) {
// 	var job model.Job

// 	rows, err := Db.QueryContext(ctx, "select * from jobs where status != $1", "loan_paid")
// 	if err != nil {
// 		fmt.Printf("error while fetching all applications: %v", err)
// 	}
// 	defer rows.Close()
// 	jobs := make([]model.Job, len(rows))
// 	for rows.Next() {
// 		err := rows.scan(job)
// 		if err != nil {
// 			fmt.Printf("error while scaning row: ", err)
// 			return nil, err
// 		}
// 		jobs := append(jobs, job)
// 	}
// 	return jobs, nil
// }

// func changeStatusJob(ctx context.Context) error {
// 	var jobs []model.Job
// 	jobs, err := fetchAllApplications(ctx)
// 	if err != nil {
// 		fmt.Printf("error while fetching all jobs with status not loan_approved: ", err)
// 		return err
// 	}
// 	var status string
// 	for job := range jobs {
// 		switch job.Status {
// 		case "created":
// 			status = "processing"
// 		case "processing":
// 			status = "loan_approved"
// 		case "loan_approved":
// 			status = "loan_paid"
// 		default:
// 			status = "created"
// 		}
// 		err := UpdateJobByApplicationID(ctx, job.ApplicationID, status)
// 		if err != nil {
// 			fmt.Println(err)
// 			return err
// 		}
// 	}
// 	return nil
// }
