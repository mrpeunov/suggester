package internal

import (
	"database/sql"
	"log"
)

func GetDBConnection() *sql.DB {
	connStr := "postgres://postgres:password@localhost/postgres?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
