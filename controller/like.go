package controller

import (
	"encoding/json"
	"fmt"
	"giiku5/model"
	"giiku5/supabase"
	"net/http"
	"strconv"

	"github.com/google/uuid"
)

// いいね情報をデータベースに挿入
func CreateLike(w http.ResponseWriter, r *http.Request) {
	supabase, _ := supabase.GetClient()

	var data model.RequestBody
	body_err := json.NewDecoder(r.Body).Decode(&data)
	if body_err != nil {
		http.Error(w, body_err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Println("IDs:", data.IDs)
	fmt.Println("UUID:", data.UUID)

	user_id, id_err := uuid.Parse(data.UUID)
	if id_err != nil {
		fmt.Println(id_err)
		return

	}

	var results []model.CreateLike

	for _, id := range data.IDs {
		other_user_id, other_id_err := uuid.Parse(id)
		if other_id_err != nil {
			fmt.Println(other_id_err)
			continue
		}
		row := model.CreateLike{
			UserID:      user_id,
			LikedUserID: other_user_id,
		}

		err := supabase.DB.From("likes").Insert(row).Execute(&results)
		if err != nil {
			fmt.Println(err)
			continue
		}
		MatchingCheck(user_id, other_user_id)
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
