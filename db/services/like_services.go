package services

import (
	"fmt"
	"github.com/nanohana2199/back_hackathon_mana-nakagawa/db/models" // modelsパッケージのインポート
	"github.com/nanohana2199/back_hackathon_mana-nakagawa/db/repositories"
)

type LikeService struct {
	LikeRepo *repositories.LikeRepository
}

// いいねを追加する
func (service *LikeService) AddLike(postId int, userId string) error {
	// models.Like 型のインスタンスを作成
	like := models.Like{
		PostID: postId,
		UserID: userId,
	}

	// LikeRepo.AddLike メソッドを呼び出す
	return service.LikeRepo.AddLike(like)
}

// いいね数を取得するメソッド
func (s *LikeService) GetLikeCount(postID string) (int, error) {
	// いいね数をカウントするクエリ
	query := "SELECT COUNT(*) FROM likes WHERE post_id = ?"
	var count int
	err := s.LikeRepo.DB.QueryRow(query, postID).Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("いいね数の取得エラー: %v", err)
	}
	return count, nil
}
