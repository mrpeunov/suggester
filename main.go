package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"suggester/internal"
)

func init() {
	err := internal.InitSchema()
	if err != nil {
		fmt.Println("СХема не создана")
	}
}

func main() {
	router := gin.Default()
	router.GET("/suggest", internal.GetSuggestHandler)
	router.POST("/suggest", internal.AddSuggestHandler)
	router.Run("localhost:9000")
}
