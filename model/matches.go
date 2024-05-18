package model

import "time"

type Match struct {
	ID        int       `json:"id"`
	User1ID   int       `json:"user1_id"`
	User2ID   int       `json:"user2_id"`
	CreatedAt time.Time `json:"created_at"`
}

// マッチング情報を挿入
type CreateMatch struct {
	User1ID int `json:"user1_id"`
	User2ID int `json:"user2_id"`
}
