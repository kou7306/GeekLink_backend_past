package model

import (
	"time"

	"github.com/google/uuid"
)

type Like struct {
	ID          int       `json:"id"`
	UserID      uuid.UUID `json:"user_id"`       // いいねを押した人のuser_id
	LikedUserID uuid.UUID `json:"liked_user_id"` // いいねを押された人のuser_id
	CreatedAt   time.Time `json:"created_at"`
}

// いいね情報を挿入
type CreateLike struct {
	UserID      uuid.UUID `json:"user_id"`
	LikedUserID uuid.UUID `json:"liked_user_id"`
}

// いいね情報を削除
type DeleteLike struct {
	ID          int       `json:"id"`
	UserID      uuid.UUID `json:"user_id"`
	LikedUserID uuid.UUID `json:"liked_user_id"`
	CreatedAt   time.Time `json:"created_at"`
}
