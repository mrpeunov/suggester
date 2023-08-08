package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"
)

func GetDBConnection() (*sql.DB, error) {
	connStr := "postgres://postgres:password@localhost/postgres?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return db, err
	}
	return db, nil
}

func InitSchema() error {
	db, err := GetDBConnection()
	if err != nil {
		fmt.Println("Что-то сделать")
	}
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
