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

// マッチングしている人を取得
func GetMatchUser(c *gin.Context) {
	supabase, _ := supabase.GetClient()

	var body model.RequestUserID
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user_id := body.UUID
	var IDs1 []model.GetMatchingUser
	var IDs2 []model.GetMatchingUser
	var IDs []model.GetMatchingUser

	// user_id = user2_idとなる行からuser1_idを取得
	err := supabase.DB.From("matches").Select("*").Eq("user2_id", user_id).Execute(&IDs1)
	if err != nil {
		log.Fatalf("Error fetching liked_user_id: %v", err)
	}

	// user_id = user1_idとなる行からuser2_idを取得
	err = supabase.DB.From("matches").Select("*").Eq("user1_id", user_id).Execute(&IDs2)
	if err != nil {
		log.Fatalf("Error fetching liked_user_id: %v", err)
	}

	IDs = append(IDs1, IDs2...)

	var MatchIDs []string

	for _, match := range IDs {
		if match.User1ID == user_id {
			matchingUserID := match.User2ID
			MatchIDs = append(MatchIDs, matchingUserID)
		} else {
			matchingUserID := match.User1ID
			MatchIDs = append(MatchIDs, matchingUserID)
		}
	}

	// マッチングしている人のユーザー情報をusersから取得
	var matchingUsers []model.UserLikedResponse
	for _, matchingUserID := range MatchIDs {
		var matchingUser []model.UserLikedResponse
		err = supabase.DB.From("users").Select("user_id", "name").Eq("user_id", matchingUserID).Execute(&matchingUser)

		if err != nil {
			log.Fatalf("Error fetching user with id %s: %v", matchingUserID, err)
		}

		matchingUsers = append(matchingUsers, matchingUser[0])
	}

	c.JSON(http.StatusOK, matchingUsers)

}
