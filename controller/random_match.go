package controller

import (
	"giiku5/model"
	"giiku5/supabase"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/labstack/gommon/log"
)

func RandomMatch(c *gin.Context) {

	
	supabase, _ := supabase.GetClient()

	var body model.RequestUserID
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Print("RandomMatch")

	user_id := body.UUID
	var users []model.UserRandomResponse

	rand.Seed(time.Now().UnixNano())

	err := supabase.DB.From("users").Select("*").Filter("user_id", "neq", user_id).Execute(&users)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

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
