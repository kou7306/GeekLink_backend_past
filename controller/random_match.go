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
	w.Header().Set("Access-Control-Allow-Origin", "https://giiku5-frontend.vercel.app")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	w.Header().Set("Content-Type", "application/json")
	// Handle preflight request
	if r.Method == http.MethodOptions {

		w.WriteHeader(http.StatusOK)
		return
	}
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

	w.WriteHeader(http.StatusOK)
	w.Write(jsonRandomUsers)
}
