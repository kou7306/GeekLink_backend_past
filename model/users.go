package model

import (
	"time"
)

type User struct {
	UserID            string    `json:"user_id"`
	Name              string    `json:"name"`               // 名前
	Sex               string    `json:"sex"`                // 性別
	Age               string    `json:"age"`                // 年齢
	Place             string    `json:"place"`              // 在住
	TopTeches         []string  `json:"top_teches"`         // トップスキル
	Teches            []string  `json:"teches"`             // スキル
	Hobby             string    `json:"hobby"`              // 趣味
	Occupation        string    `json:"occupation"`         // 職種
	Affiliation       string    `json:"affiliation"`        // 所属
	Qualification     []string  `json:"qualification"`      // 資格
	Editor            string    `json:"editor"`             // エディタ
	ImageURL          string    `json:"image_url"`          // アイコン画像
	Message           string    `json:"message"`            // メッセージ
	Portfolio         string    `json:"portfolio"`          // ポートフォリオ
	Graduate          string    `json:"graduate"`           // 卒業年度
	DesiredOccupation string    `json:"desiredOccupation"` // 希望職種
	Faculty           string    `json:"faculty"`            // 学部
	Experience        []string  `json:"experience"`         // 経験
	GitHub            string    `json:"github"`
	Twitter           string    `json:"twitter"`
	Qiita             string    `json:"qiita"`
	Zenn              string    `json:"zenn"`
	AtCoder           string    `json:"atcoder"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}

// ホーム画面で返すユーザー情報
type TopUserResponse struct {
	UserID    int      `json:"user_id"`
	Name      string   `json:"name"`
	TopTeches []string `json:"top_teches"`
	ImageURL  string   `json:"image_url"`
}

// ランダムマッチングで返すユーザー情報
type UserRandomResponse struct {
	UserID     string   `json:"user_id"`
	Name       string   `json:"name"`
	Sex        string   `json:"sex"`
	Age        string   `json:"age"`
	Place      string   `json:"place"`
	Occupation string   `json:"occupation"`
	TopTeches  []string `json:"top_teches"`
	ImageURL   string   `json:"image_url"`
}

// メッセージ一覧で返すユーザー情報
type MessageUserResponse struct {
	UserID    string   `json:"user_id"`
	Name      string   `json:"name"`
	Sex       string   `json:"sex"`
	Age       string   `json:"age"`
	TopTeches []string `json:"top_teches"`
	ImageURL  string   `json:"image_url"`
}

type RequestUserID struct {
	UUID string `json:"uuid"`
}
