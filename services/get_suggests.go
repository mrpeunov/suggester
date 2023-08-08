package services

import (
	"database/sql"
	"fmt"
	"log"
	"suggester/models"
)

func getSuggests(query string, db *sql.DB) []models.Suggest {
	suggesters := make([]models.Suggest, 0, 5)

	rows, err := db.Query(
		`SELECT "query" FROM "query" 
               WHERE query LIKE $1 || '%' 
               ORDER BY "amount" DESC 
               LIMIT 5`, query,
	)
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var suggest models.Suggest
		if err := rows.Scan(&suggest.Text); err != nil {
			fmt.Println("Error")
		}
		suggesters = append(suggesters, suggest)
	}

	return suggesters
}
