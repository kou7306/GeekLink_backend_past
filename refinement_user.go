package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nedpals/supabase-go"
)

func main() {
	router := gin.Default()
	router.POST("/users", getUsers) 
	router.Run(":8080")
}


// FilterRequestはフィルタリングのためのフィールドと値を受け取る構造体
type FilterRequest struct {
	Field string `json:"field" binding:"required"`
	Value string `json:"value" binding:"required"`
	Limit int `json:"limit" binding:"required"`
}

func getUsers(c *gin.Context) {
	supabaseURL := "" 
	supabaseKey := "" 

	client := supabase.CreateClient(supabaseURL, supabaseKey)

	var req FilterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var users []User
	query := client.DB.From("users").Select("*").Eq(req.Field, req.Value).Limit(req.Limit)
	err := query.Execute(&users)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}
