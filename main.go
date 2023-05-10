package main

import (
	"log"

	"github.com/ereminiu/apitest/pkg/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/solve", handlers.Solve)

	log.Println("Server is running at port 2048")
	r.Run(":2048")
}
