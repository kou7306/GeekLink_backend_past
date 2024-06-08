package api

import (
	"fmt"
	"giiku5/model"
	"giiku5/supabase"
	"log"

	"github.com/google/uuid"
)

// ユーザーがマッチングしているか確認(互いにいいねしているか)
func MatchingCheck(user_id uuid.UUID, other_user_id uuid.UUID) {

	// 片方のユーザーでチェック
	var filtered_Likes []model.Like
	var row_id int
	filtered_Likes, row_id = FilterLikes(user_id, other_user_id)

	// もう片方のユーザーでチェック
	var filtered_other_Likes []model.Like
	var other_row_id int
	filtered_other_Likes, other_row_id = FilterLikes(other_user_id, user_id)

	// 互いをいいねしていたらCreateMatchingを実行
	fmt.Println(len(filtered_Likes), len(filtered_other_Likes))
	if len(filtered_Likes) == 1 && len(filtered_other_Likes) == 1 {
		fmt.Println("matching successful")
		CreateMatching(user_id, other_user_id)
		DeleteLike(row_id, other_row_id)
	}
}

// 自分のいいねした人の中から特定のユーザーをいいねしているかフィルターする
func FilterLikes(user_id uuid.UUID, other_user_id uuid.UUID) ([]model.Like, int) {
	supabase, _ := supabase.GetClient()

	var likes []model.Like
	err := supabase.DB.From("likes").Select("*").Eq("user_id", user_id.String()).Execute(&likes)
	if err != nil {
		log.Printf("Error: %v", err)
		return []model.Like{}, 0
	}

	var filtered_Likes []model.Like
	for _, like := range likes {
		if like.LikedUserID == other_user_id {
			filtered_Likes = append(filtered_Likes, like)
		}
	}

	if len(filtered_Likes) == 0 {
		return []model.Like{}, 0
	}

	// filtered_Likes[0].IDはマッチングしたときにlikesテーブルから削除するために使用
	return filtered_Likes, filtered_Likes[0].ID
}

// ユーザーがマッチングしていたら呼び出す. マッチングIDを発行
func CreateMatching(user1_id uuid.UUID, user2_id uuid.UUID) {
	supabase, _ := supabase.GetClient()

	match := model.CreateMatch{
		User1ID: user1_id,
		User2ID: user2_id,
	}

	var row []model.CreateMatch
	err := supabase.DB.From("matches").Insert(match).Execute(&row)
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}

}
