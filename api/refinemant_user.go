package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nedpals/supabase-go"
)

// FilterRequestはフィルタリングのためのフィールドと値を受け取る構造体
type FilterRequest struct {
	Field string `json:"field" binding:"required"`
	Value string `json:"value" binding:"required"`
	Limit int `json:"limit" binding:"required"`
}

	type User struct {
		Auth_id  string `json:"auth_id"`	
		Name  string `json:"name"`	
		Sex string `json:"sex"`
		Age string `json:"age"`
		Place string `json:"place"`
		Occupation string `json:occupation"`
		Techs []string `json:"techs"`
		TopTechs []string `json:"topTechs"`
		Image_url string `json:" image_url"`
		Hobby string `json:"hobby"`
		Editor string `json:" editor"`
		Affiliation string `json:"affiliation"`
		Qualification []string `json:"qualification"`
		Message string `json:"message"`
		Portfolio string `json:"portfolio"`
		Graduate string `json:"graduate"`
		DesiredOccupation string `json:"desiredOccupation"`
		Faculty string `json:"faculty"`
		Experience []string `json:"experience"`
		Github string `json:"github"`
		Twitter string `json:"twitter"`
		Zenn string `json:"zenn"`
		Qiita string `json:"qiita"`
		Atcoder string `json:"atcoder"`
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
	query := client.DB.From("users").Select("*").Eq(req.Field, req.Value)
	err := query.Execute(&users)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}
