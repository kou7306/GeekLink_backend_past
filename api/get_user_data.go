package api

import (
	"giiku5/model"
	"giiku5/supabase"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUserData(c *gin.Context) {
	log.Printf("GetUserData")
	client, err := supabase.GetClient()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to initialize Supabase client"})
		return
	}

	var requestBody model.RequestBody
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Printf("Received UUID: %s\n", requestBody.UUID)

	userid := requestBody.UUID

	var userData []model.User
	err = client.DB.From("users").Select("*").Eq("user_id", userid).Execute(&userData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to query database"})
		return
	}

	log.Printf("%+v", userData)

	if len(userData) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, userData[0])
}
