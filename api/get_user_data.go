package api

import (
	"giiku5/model"
	"giiku5/supabase"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUserData(c *gin.Context) {
	userID := c.Param("user_id")

	supabaseClient, err := supabase.GetClient()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to initialize Supabase client"})
		return
	}

	var users []model.User
	err = supabaseClient.DB.From("users").Select("*").Eq("user_id", userID).Execute(&users)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to query database"})
		return
	}

	if len(users) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, users[0])
}
