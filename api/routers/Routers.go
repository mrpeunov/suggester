package routers

import (
	"github.com/gin-gonic/gin"
	"suggester/api/handlers"
)

func InitRouters() {
	router := gin.Default()
	router.GET("/suggest", handlers.GetSuggestHandler)
	router.POST("/suggest", handlers.AddSuggestHandler)
	router.Run("localhost:9000")
}
