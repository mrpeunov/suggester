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
		var suggest Suggest
		if err := rows.Scan(&suggest.Text); err != nil {
			fmt.Println("Error")
		}
		suggesters = append(suggesters, suggest)
	}

	return suggesters
}

func getSuggestId(query string, db *sql.DB) int64 {
	row := db.QueryRow(`SELECT "query_id" FROM "query" WHERE query = $1`, query)
	fmt.Println(row)
	var id int64
	if err := row.Scan(&id); err != nil {
		return -1
	}
	return id
}

func incSuggest(id int64, db *sql.DB) {
	_, err := db.Exec("UPDATE query SET amount = amount + 1 WHERE query_id = $1", id)
	if err != nil {
		log.Fatal("Error inc")
	}
}

func addSuggest(sg Suggest, db *sql.DB) {
	_, err := db.Exec("INSERT INTO query(query, amount) VALUES ($1, 1)", sg.Text)
	if err != nil {
		log.Fatal("Error add suggest")
	}
}

func AddSuggestHandler(c *gin.Context) {
	var suggest Suggest
	err := c.BindJSON(&suggest)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, nil)
	}

	db := GetDBConnection()

	if id := getSuggestId(suggest.Text, db); id == -1 {
		addSuggest(suggest, db)
	} else {
		incSuggest(id, db)
	}

	c.IndentedJSON(http.StatusOK, nil)
}

func GetSuggestHandler(c *gin.Context) {
	query := c.DefaultQuery("q", "")
	db := GetDBConnection()
	suggests := getSuggests(query, db)
	c.IndentedJSON(http.StatusOK, suggests)
}
