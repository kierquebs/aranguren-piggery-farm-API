package database

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var CCDB *sql.DB
var err error

func CCDBCon() {
	CCDB, err = sql.Open("postgres", os.Getenv("POSTGRESQL_ADDON_URI"))
	if err != nil {
		log.Fatalf("UNABLE to establish database connection :: '%v' ", err)
	}

	log.Println("Database connection SUCCESSFULLY established!")
}
