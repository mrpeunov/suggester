package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"suggester/models"
	"suggester/repository"
	"suggester/services"
)

func AddSuggestHandler(c *gin.Context) {
	var suggest models.Suggest

	repo, err := repository.GetSuggestRepo()
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
	}

	if err = c.BindJSON(&suggest); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, nil)
	}

	if err = services.AddSuggest(suggest, repo); err != nil {
		c.JSON(http.StatusInternalServerError, nil)
	}

	c.IndentedJSON(http.StatusOK, nil)
}

func GetSuggestHandler(c *gin.Context) {
	repo, err := repository.GetSuggestRepo()
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
	}

	query := c.DefaultQuery("q", "")

	suggests, err := repo.Get(query, 5)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
	}

	c.IndentedJSON(http.StatusOK, suggests)
}
