package driver

import (
	"database/sql"
	"log"
	"os"

	"github.com/lib/pq"
)

var db *sql.DB

func LogFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func ConnectDB() *sql.DB {

	//get our env variables, test for errors
	pgURL, err := pq.ParseURL(os.Getenv("ELEPHANTSQL_URL"))
	LogFatal(err)

	//create connection to our database, test for errors. pgURL holds all of our database info such as: db name, host, password, port and user
	db, err = sql.Open("postgres", pgURL)
	LogFatal(err)

	//ping data base test for errors
	err = db.Ping()
	LogFatal(err)

	return db
}
