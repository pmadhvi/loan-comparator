package main

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/pmadhvi/loan-comparator/handler"
	"github.com/pmadhvi/loan-comparator/postgres"
	"github.com/pmadhvi/loan-comparator/service"
	"github.com/sirupsen/logrus"
)

func main() {
	fmt.Println("Hello welcome to loan comparator!")
	//setup logger for the application
	var log = logrus.New()
	log.SetReportCaller(true)
	log.SetOutput(os.Stdout)

	//setup postgress connection
	db_user := os.Getenv("DB_USER")
	db_pass := os.Getenv("DB_PASS")
	db_name := os.Getenv("DB_NAME")
	dbinfo := fmt.Sprintf("host=localhost port= 5432 user=%s password=%s dbname=%s sslmode=disable", db_user, db_pass, db_name)
	db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		log.Error(errors.New("error connecting database, "), err)
	}

	//setup repositories
	var (
		applicationRepo = postgres.NewApplicationRepository(*log, db)
		jobRepo         = postgres.NewJobRepository(*log, db)

		//setup services
		applicationService = service.ApplicationService{ApplicationRepository: applicationRepo}
		jobService         = service.JobService{JobRepository: jobRepo}

		//setup router
		router = handler.Router{
			Log:                *log,
			ApplicationService: applicationService,
			JobService:         jobService,
		}
	)
	// setup all the routes
	http.Handle("/", router.Router())

	// start the http server on port 3000 and pass the chi router as handler
	http.ListenAndServe(":3000", nil)

}
