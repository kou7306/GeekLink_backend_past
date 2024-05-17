package model

import "time"

type User struct {
	UserID        int       `json:"user_id"`
	Name          string    `json:"name"`
	Sex           string    `json:"sex"`
	Age           int       `json:"age"`
	Place         string    `json:"place"`
	TopTeches     []string  `json:"top_teches"`
	Teches        []string  `json:"teches"`
	Hobby         string    `json:"hobby"`
	Occupation    string    `json:"occupation"`
	Affiliation   string    `json:"affiliation"`
	Qualification string    `json:"qualification"`
	Editor        string    `json:"editor"`
	ImageURL      string    `json:"image_url"`
	GitHub        string    `json:"github"`
	Twitter       string    `json:"twitter"`
	Qiita         string    `json:"qiita"`
	Zenn          string    `json:"zenn"`
	AtCoder       string    `json:"atcoder"`
	Message       string    `json:"message"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

// ランダムマッチングで返すユーザー情報
type UserRandomResponse struct {
	UserID     int      `json:"user_id"`
	Name       string   `json:"name"`
	Sex        string   `json:"sex"`
	Age        int      `json:"age"`
	Place      string   `json:"place"`
	Occupation string   `json:"occupation"`
	TopTeches  []string `json:"top_teches"`
}
