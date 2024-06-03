package api

import (
	"giiku5/supabase"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// リクエストボディの構造体を定義
type RequestBody struct {
	UUID string `json:"uuid"`
}

type Match struct {
	ID      int    `json:"id"`
	User1ID string `json:"user1_id"`
	User2ID string `json:"user2_id"`
}

type User struct {
	ID        string   `json:"user_id"`
	Name      string   `json:"name"`
	ImgURL    string   `json:"img_url"`
	Languages []string `json:"languages"`
	Age       string   `json:"age"`
	Sex       string   `json:"sex"`
}

func GetMatchingUser(c *gin.Context) {
	client, err := supabase.GetClient()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to initialize Supabase client"})
		return
	}

	// リクエストボディの読み取り
	var requestBody RequestBody
	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// 取得したUUIDを使用してデータベースクエリなどの処理を実行
	log.Printf("Received UUID: %s\n", requestBody.UUID)

	userid := requestBody.UUID
	var matches1 []Match
	var matches2 []Match
	var allMatches []Match

	// user1_idに一致する行を取得
	err = client.DB.From("matches").
		Select("*").
		Eq("user1_id", userid).
		Execute(&matches1)
	if err != nil {
		log.Fatalf("Error fetching matches for user1_id: %v", err)
	}

	// user2_idに一致する行を取得
	err = client.DB.From("matches").
		Select("*").
		Eq("user2_id", userid).
		Execute(&matches2)
	if err != nil {
		log.Fatalf("Error fetching matches for user2_id: %v", err)
	}

	// matches1とmatches2を結合
	allMatches = append(matches1, matches2...)

	log.Print(allMatches)

	var matchingUserIDs []string

	for _, match := range allMatches {
		if match.User1ID == userid {
			log.Print(match.User2ID)
			matchingUserID := match.User2ID
			matchingUserIDs = append(matchingUserIDs, matchingUserID)
		} else {
			matchingUserID := match.User1ID
			matchingUserIDs = append(matchingUserIDs, matchingUserID)
		}
	}

	log.Print(matchingUserIDs)

	// マッチングユーザーの情報をユーザーテーブルから取得
	var matchingUsers []User
	for _, matchingUserID := range matchingUserIDs {
		var matchingUser []User
		err = client.DB.From("users").
			Select("user_id", "name").
			Eq("user_id", matchingUserID).
			Execute(&matchingUser)

		if err != nil {
			log.Fatalf("Error fetching user with id %s: %v", matchingUserID, err)
		}

		matchingUsers = append(matchingUsers, matchingUser[0])
	}

	c.JSON(http.StatusOK, matchingUsers)
}

