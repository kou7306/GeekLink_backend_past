package controller

import (
	"fmt"
	"giiku5/model"
	"giiku5/supabase"
	"strconv"

	"github.com/google/uuid"
)

// いいね情報をデータベースに挿入
func CreateLike() {
	supabase, _ := supabase.GetClient()

	// 仮置き. 自身のIDと相手のIDの取得方法分かり次第修正.
	my_user_id, id_err := uuid.Parse("3a54e201-9341-480b-b3a6-e8a538dfec6d")
	if id_err != nil {
		fmt.Println(id_err)
		return
	}
	other_user_id, other_id_err := uuid.Parse("dc1f8f58-b6bf-424a-9809-7abdbec16781")
	if other_id_err != nil {
		fmt.Println(other_id_err)
		return
	}

	row := model.CreateLike{
		UserID:      my_user_id,
		LikedUserID: other_user_id,
	}

	var results []model.CreateLike

	err := supabase.DB.From("likes").Insert(row).Execute(&results)
	if err != nil {
		fmt.Println(err)
	}

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
