package model

import (
	"time"

	"github.com/google/uuid"
)

type Match struct {
	ID        int       `json:"id"`
	User1ID   uuid.UUID `json:"user1_id"`
	User2ID   uuid.UUID `json:"user2_id"`
	CreatedAt time.Time `json:"created_at"`
}

// マッチング情報を挿入
type CreateMatch struct {
	User1ID uuid.UUID `json:"user1_id"`
	User2ID uuid.UUID `json:"user2_id"`
}
