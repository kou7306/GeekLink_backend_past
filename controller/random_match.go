package controller

import (
	"encoding/json"
	"giiku5/model"
	"giiku5/supabase"
	"math/rand"
	"net/http"
	"time"
)

func Random_Match(w http.ResponseWriter, r *http.Request) {
	supabase, _ := supabase.GetClient()

	// 自分のidを除外する. 一旦空値に
	my_user_id := ""

	var users []model.UserRandomResponse

	rand.Seed(time.Now().UnixNano())

	err := supabase.DB.From("users").Select("*").Filter("user_id", "neq", my_user_id).Execute(&users)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
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

	json.NewEncoder(w).Encode(random_users)
}
