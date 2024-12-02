package repositories

import (
	"database/sql"
	"fmt"
	"github.com/nanohana2199/back_hackathon_mana-nakagawa/db/models" // modelsパッケージをインポート
)

type LikeRepository struct {
	DB *sql.DB
}

// AddLike は Like テーブルに新しい「いいね」を追加するメソッドです。
func (repo *LikeRepository) AddLike(like models.Like) error {
	query := "INSERT INTO likes (post_id, user_id) VALUES (?, ?)"
	_, err := repo.DB.Exec(query, like.PostID, like.UserID)
	if err != nil {
		return fmt.Errorf("failed to add like: %v", err)
	}
	return nil
}

func (r *LikeRepository) GetLikeCount(postID string) (int, error) {
	query := "SELECT COUNT(*) FROM likes WHERE post_id = ?"
	var count int
	err := r.DB.QueryRow(query, postID).Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}
