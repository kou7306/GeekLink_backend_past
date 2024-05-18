package api

import (
	"encoding/json"
	"giiku5/supabase"
	"log"
	"net/http"
	"strconv"
)

// リクエストボディの構造体を定義
type RequestBody struct {
    UUID int `json:"uuid"`
}

type Match struct {
	ID int `json:"id"`
    User1ID int `json:"user1_id"`
    User2ID int `json:"user2_id"`
}

func GetMatchingUser(w http.ResponseWriter, r *http.Request) {


    client, err := supabase.GetClient()
    if err != nil {
        http.Error(w, "Failed to initialize Supabase client", http.StatusInternalServerError)
        return
    }

	// Content-Typeのチェック
    if r.Header.Get("Content-Type") != "application/json" {
        http.Error(w, "Content-Type is not application/json", http.StatusUnsupportedMediaType)
        return
    }

    // リクエストボディの読み取り
    var requestBody RequestBody
    decoder := json.NewDecoder(r.Body)
    err = decoder.Decode(&requestBody)



	userid := strconv.Itoa(requestBody.UUID)
    var matches1 []Match
    var matches2 []Match
    var allMatches []Match

    // user1_idに一致する行を取得
    err = client.DB.From("matches").
        Select("*").
        Eq("user1_id", userid).
        Execute(&matches1)
    if err != nil {
        log.Fatalf("Error fetching matches for user1_id: %v", err)
    }

    // user2_idに一致する行を取得
    err = client.DB.From("matches").
        Select("*").
        Eq("user2_id", userid).
        Execute(&matches2)
    if err != nil {
        log.Fatalf("Error fetching matches for user2_id: %v", err)
    }

    // matches1とmatches2を結合
    allMatches = append(matches1, matches2...)


    var matchingUserIDs []int
var matchingUserIDsBytes []byte
for _, match := range allMatches {
	if strconv.Itoa(match.User1ID) == userid {
		matchingUserIDs = append(matchingUserIDs, match.User2ID)
	} else {
		matchingUserIDs = append(matchingUserIDs, match.User1ID)
	}
}

matchingUserIDsBytes, err = json.Marshal(matchingUserIDs)
if err != nil {
	http.Error(w, "Failed to marshal matchingUserIDs", http.StatusInternalServerError)
	return
}

w.Header().Set("Content-Type", "application/json")
w.WriteHeader(http.StatusOK)
w.Write(matchingUserIDsBytes)
}