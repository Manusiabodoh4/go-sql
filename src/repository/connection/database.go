package connection

import (
	"database/sql"
	"fmt"
	"time"
)

var constrPg string = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
	"127.0.0.1", 2345, "postgres", "123", "todo")

var databasePGSQL *sql.DB
var err error

func GetConnectionPostgres() *sql.DB {

	if databasePGSQL == nil {

		fmt.Println("Create Connection")

		databasePGSQL, err = sql.Open("postgres", constrPg)

		if err != nil {
			panic(err)
		}

		databasePGSQL.SetMaxIdleConns(10)
		databasePGSQL.SetMaxOpenConns(100)

		databasePGSQL.SetConnMaxIdleTime(5 * time.Minute)
		databasePGSQL.SetConnMaxLifetime(1 * time.Hour)
	}

	return databasePGSQL

}
