package handlers

import (
	"log"
	"net/http"

	"github.com/ereminiu/apitest/pkg/models"
	"github.com/ereminiu/apitest/pkg/service"
	"github.com/gin-gonic/gin"
)

func Solve(c *gin.Context) {
	var input models.Grid
	if err := c.BindJSON(&input); err != nil {
		log.Fatalf("failed to read puzzle grid\n")
		return
	}

	// check it's a valid grid
	code := service.IsValid(input)
	if code != models.CorrectGrid {
		c.JSON(403, gin.H{
			"message": models.Decode(code),
		})
		return
	}

	sol := service.Solve(input)

	c.JSON(http.StatusOK, gin.H{
		"sol": models.NewResponse(sol).Sol,
	})
}
