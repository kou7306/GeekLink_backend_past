package controller

import (
	"fmt"
	"giiku5/model"
	"giiku5/supabase"
	"strconv"
)

// いいね情報をデータベースに挿入
func CreateLike() {
	supabase, _ := supabase.GetClient()

	// 仮置き. 自身のIDと相手のIDの取得方法分かり次第修正.
	my_user_id := 1
	other_user_id := 2

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

	err := supabase.DB.From("likes").Delete().Eq("id", strconv.Itoa(row_id))
	if err != nil {
		fmt.Println(err)
	}

	other_err := supabase.DB.From("likes").Delete().Eq("id", strconv.Itoa(other_row_id))
	if other_err != nil {
		fmt.Println(other_err)
	}

}
