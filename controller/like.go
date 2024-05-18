package controller

import (
	"fmt"
	"giiku5/model"
	"giiku5/supabase"
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
