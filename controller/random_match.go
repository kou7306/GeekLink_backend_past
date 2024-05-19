package controller

import (
	"encoding/json"
	"fmt"
	"giiku5/model"
	"giiku5/supabase"
	"math/rand"
	"net/http"
	"time"
)

func Random_Match(w http.ResponseWriter, r *http.Request) {
	supabase, _ := supabase.GetClient()

	var body model.RequestUserID
	_ = json.NewDecoder(r.Body).Decode(&body)

	user_id := body.UUID
	var users []model.UserRandomResponse

	rand.Seed(time.Now().UnixNano())

	err := supabase.DB.From("users").Select("*").Filter("user_id", "neq", user_id).Execute(&users)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var users_num = 5
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

	jsonRandomUsers, err := json.Marshal(random_users)
	if err != nil {
		fmt.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonRandomUsers)
}
