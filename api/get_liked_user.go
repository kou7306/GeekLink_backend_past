package api

import (
	"fmt"
	"giiku5/model"
	"giiku5/supabase"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// いいねされている人を取得
func GetLikedUser(c *gin.Context) {
	supabase, _ := supabase.GetClient()

	var body model.RequestUserID
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user_id := body.UUID
	var IDs []model.GetUserLikedID

	// 自分のuser_idがliked_user_idと一致する行を取得
	err := supabase.DB.From("likes").Select("*").Eq("liked_user_id", user_id).Execute(&IDs)
	if err != nil {
		log.Fatalf("Error fetching liked_user_id: %v", err)
	}

	var users []model.UserLikedResponse

	// いいねされている人のIDをusersから検索して情報を取得
	for _, id := range IDs {
		err := supabase.DB.From("users").Select("*").Eq("user_id", id.UserID).Execute(&users)
		if err != nil {
			log.Fatalf("Error find user: %v", err)
		}
	}

	fmt.Println(users)

	c.JSON(http.StatusOK, users)
}
