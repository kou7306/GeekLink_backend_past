package api

import (
	"giiku5/model"
	"giiku5/supabase"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func RandomMatch(c *gin.Context) {
	supabase, _ := supabase.GetClient()

	var body model.RequestUserID
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user_id := body.UUID
	var users []model.UserRandomResponse

	rand.Seed(time.Now().UnixNano())

	err := supabase.DB.From("users").Select("*").Filter("user_id", "neq", user_id).Execute(&users)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// マッチングしている相手は除外する
	users = DeleteMatchingUser(user_id, users)

	// ランダムマッチで表示する相手の人数
	users_num := 5
	if len(users) < users_num {
		users_num = len(users)
	}

	var random_users []model.UserRandomResponse
	for i := 0; i < users_num; i++ {
		if len(users) == 0 {
			break
		}
		index := rand.Intn(len(users))
		random_users = append(random_users, users[index])
		users = append(users[:index], users[index+1:]...)
	}

	c.JSON(http.StatusOK, random_users)
}

// マッチングしている相手を除外
func DeleteMatchingUser(user_id string, users []model.UserRandomResponse) []model.UserRandomResponse {
	supabase, _ := supabase.GetClient()

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

	// MatchIDsに一致するUserIDを持つusersを削除
	for _, matchID := range MatchIDs {
		for i := 0; i < len(users); i++ {
			if users[i].UserID == matchID {
				users = append(users[:i], users[i+1:]...)
				i--
			}
		}
	}

	return users
}
