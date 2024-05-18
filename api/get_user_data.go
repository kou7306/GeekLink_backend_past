package api

import (
	"encoding/json"
	"giiku5/supabase"
	"log"
	"net/http"
)


type UserData struct {
    UserID string `json:"user_id"`
	Name string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	Sex string `json:"sex"`
	Age string `json:"age"`
	Place string `json:"place"`
	TopTeches string `json:"top_teches"`
	Teches string `json:"teches"`
	Hobby string `json:"hobby"`
	Occupation string `json:"occupation"`
	Affiliation string `json:"affiliation"`
	Qualification string `json:"qualification"`
	Editor string `json:"editor"`
	Github string `json:"github"`
	Twitter string `json:"twitter"`
	Qiita string `json:"qiita"`
	Zenn string `json:"zenn"`	
	Message string `json:"message"`
	Portfolio string `json:"portfolio"`
	Graduate string `json:"graduate"`
	DesiredOcupation string `json:"desiredOccupation"`
	Faculty string `json:"faculty"`
	Experience string `json:"experience"`
	ImageURL string `json:"image_url"`
} 




func GetUserData(w http.ResponseWriter, r *http.Request) {

	log.Printf("GetUserData")
	client, err := supabase.GetClient()
    if err != nil {
        http.Error(w, "Failed to initialize Supabase client", http.StatusInternalServerError)
        return
    }

	

	
    // リクエストボディの読み取り
    var requestBody RequestBody
    // リクエストボディをデコード
    _ = json.NewDecoder(r.Body).Decode(&requestBody)
	

    // 取得したUUIDを使用してデータベースクエリなどの処理を実行
    log.Printf("Received UUID: %s\n", requestBody.UUID)


	userid := requestBody.UUID

	var userData []UserData
	err = client.DB.From("users").Select("*").Eq("user_id", userid).Execute(&userData)
		if err != nil {
		  panic(err)
		}

		log.Printf("%+v", userData)

		// Convert messages to JSON byte slice
		userDataJSON, err := json.Marshal(userData[0])
		if err != nil {
				http.Error(w, "Failed to marshal userData to JSON", http.StatusInternalServerError)
				return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(userDataJSON)
}