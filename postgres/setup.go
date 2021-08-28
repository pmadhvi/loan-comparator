package postgres

import (
	"database/sql"
	"fmt"

	log "github.com/sirupsen/logrus"
)

// TODO: Use env
// TODO: Also pass the db obejct from main to this module
const (
	DB_USER = "madhvip"
	DB_PASS = "madhvip"
	DB_NAME = "loan"
)

func connectToDatabase() (db *sql.DB, err error) {
	dbinfo := fmt.Sprintf("host=localhost port= 5432 user=%s password=%s dbname=%s sslmode=disable", DB_USER, DB_PASS, DB_NAME)
	db, err = sql.Open("postgres", dbinfo)
	chkerr(err)
	log.Info("Connect to todolist database")
	return
}

func chkerr(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}
