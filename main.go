package main

import (
	"giiku5/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello world")
	})

	router.GET("/random-match", controller.Random_Match)

	router.Run(":8080")
}
