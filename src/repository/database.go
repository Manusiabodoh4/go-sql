package repository

import (
	"database/sql"
	"fmt"
	"time"
)

func GetConnectionPostgres() *sql.DB {

	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		"127.0.0.1", 2345, "postgres", "123", "todo")

	database, err := sql.Open("postgres", connStr)

	if err != nil {
		panic(err)
	}

	database.SetMaxIdleConns(10)
	database.SetMaxOpenConns(100)

	database.SetConnMaxIdleTime(5 * time.Minute)
	database.SetConnMaxLifetime(1 * time.Hour)

	return database

}
