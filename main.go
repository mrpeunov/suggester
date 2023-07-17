package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Suggest struct {
	Text string `json:"query"`
}

func addToStatistic(query string) {
	fmt.Println("statistic " + query)
}

func getSuggests(query string) []Suggest {
	suggesters := make([]Suggest, 0, 5)
	for i := 0; i < 5; i++ {
		suggesters = append(suggesters, Suggest{query + strconv.Itoa(i)})
	}
	return suggesters
}

func suggest(c *gin.Context) {
	query := c.DefaultQuery("q", "")
	addToStatistic(query)
	suggests := getSuggests(query)
	c.IndentedJSON(http.StatusOK, suggests)
}

func main() {
	router := gin.Default()
	router.GET("/suggest", suggest)
	router.Run("localhost:9000")
}
