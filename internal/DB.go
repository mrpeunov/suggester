package internal

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"
)

func GetDBConnection() *sql.DB {
	connStr := "postgres://postgres:password@localhost/postgres?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func InitSchema() error {
	db := GetDBConnection()
	query := `
		CREATE TABLE IF NOT EXISTS query(
    			query_id serial PRIMARY KEY, 
				query VARCHAR (148) NOT NULL,
    			amount INTEGER,
				created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT now(),
    			updated_at TIMESTAMP WITHOUT TIME ZONE DEFAULT now()
    	)`

	ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFunc()

	res, err := db.ExecContext(ctx, query)
	if err != nil {
		log.Printf("Error %s when creating product table", err)
		return err
	}
	fmt.Println(res)
	return nil
}
