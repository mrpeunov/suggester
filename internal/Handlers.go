package internal

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type Suggest struct {
	Text string `json:"query"`
}

func getSuggests(query string, db *sql.DB) []Suggest {
	suggesters := make([]Suggest, 0, 5)

	rows, err := db.Query(`SELECT "query" FROM "queries"`)
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var suggest Suggest
		if err := rows.Scan(&suggest.Text); err != nil {
			fmt.Println("Error")
		}
		suggesters = append(suggesters, suggest)
	}

	return suggesters
}

func GetSuggestHandler(c *gin.Context) {
	query := c.DefaultQuery("q", "")
	db := GetDBConnection()
	suggests := getSuggests(query, db)
	c.IndentedJSON(http.StatusOK, suggests)
}
