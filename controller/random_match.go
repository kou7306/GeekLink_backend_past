package controller

import (
	"giiku5/api"
	"giiku5/model"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func Random_Match(c *gin.Context) {
	supabase := api.SupabaseClient()

	// 自分のidを除外するため, 一旦固定の値に
	my_user_id := "2"

	var users []model.UserRandomResponse

	rand.Seed(time.Now().UnixNano())

	err := supabase.DB.From("users").Select("*").Filter("user_id", "neq", my_user_id).Execute(&users)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 何人の情報を返すか不確定, ひとまず2人のユーザー情報をランダムに抽出
	const users_num = 2
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
