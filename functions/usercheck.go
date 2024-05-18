package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	supa "github.com/nedpals/supabase-go"
)

type CheckUserRequest struct {
	UUID string `json:"uuid" binding:"required"`
}

type CheckUserResponse struct {
	Exists bool `json:"exists"`
}

func CheckUser(c *gin.Context) {
	supabaseUrl := ""
	supabaseKey := ""
	client := supa.CreateClient(supabaseUrl, supabaseKey)

	var req CheckUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}



	// UUIDを使ってユーザを検索
	var users []User
	err := client.DB.From("users").Select("name").Eq("uuid", req.UUID).Execute(&users)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if len(users) == 0 {
		// ユーザが存在しない場合
		c.JSON(http.StatusOK, CheckUserResponse{Exists: false})
		return
	}

	// ユーザが存在する場合
	c.JSON(http.StatusOK, CheckUserResponse{Exists: true})
}
