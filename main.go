package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func handler(c *gin.Context) {
	c.String(http.StatusOK, "Hello, world!")
}

func main() {
	router := gin.Default()

	fmt.Println("Starting server on port 8080...")

	router.GET("/", handler)

	if err := router.Run(":8080"); err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}
