package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"suggester/internal"
)

func main() {
	router := gin.Default()
	router.GET("/suggest", internal.GetSuggestHandler)
	//router.POST("/suggest")
	router.Run("localhost:9000")
}
