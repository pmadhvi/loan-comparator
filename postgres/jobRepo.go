package postgres

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/pmadhvi/loan-comparator/model"
	log "github.com/sirupsen/logrus"
)

type JobRepository struct {
	Log log.Logger
	Db  *sql.DB
}

func (jr JobRepository) FindJobByApplicationID(ctx context.Context, id int) (model.Job, error) {
	fmt.Println("Impemented FindJob")
	return model.Job{}, nil
}
