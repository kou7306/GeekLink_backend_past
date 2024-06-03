package controller

import (
	"fmt"
	"giiku5/model"
	"giiku5/supabase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/google/uuid"
)

// いいね情報をデータベースに挿入
func CreateLike(c *gin.Context) {
	supabase, _ := supabase.GetClient()

	var data model.RequestBody
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println("IDs:", data.IDs)
	fmt.Println("UUID:", data.UUID)

	user_id, err := uuid.Parse(data.UUID)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID"})
		return
	}

	var results []model.CreateLike

	for _, id := range data.IDs {
		other_user_id, err := uuid.Parse(id)
		if err != nil {
			fmt.Println(err)
			continue
		}
		row := model.CreateLike{
			UserID:      user_id,
			LikedUserID: other_user_id,
		}

		err = supabase.DB.From("likes").Insert(row).Execute(&results)
		if err != nil {
			fmt.Println(err)
			continue
		}
		MatchingCheck(user_id, other_user_id)
	}

	c.JSON(http.StatusOK, gin.H{"message": "Likes created successfully"})
}

// マッチングしたら対象のいいね情報を削除
func DeleteLike(row_id int, other_row_id int) {
	supabase, _ := supabase.GetClient()

	var row model.DeleteLike
	err := supabase.DB.From("likes").Delete().Eq("id", strconv.Itoa(row_id)).Execute(&row)
	if err != nil {
		fmt.Println(err)
	}

	var other_row model.DeleteLike
	other_err := supabase.DB.From("likes").Delete().Eq("id", strconv.Itoa(other_row_id)).Execute(&other_row)
	if other_err != nil {
		fmt.Println(other_err)
	}

}
