package postgres

import (
	"database/sql"
	"fmt"

	"github.com/pmadhvi/loan-comparator/model"
)

func create(db *sql.DB, application model.Application) error {
	fmt.Println("adding application to db")

	return nil

}
