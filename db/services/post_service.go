package services

import (
	"github.com/nanohana2199/back_hackathon_mana-nakagawa/db/models" // models.Post をインポート
	"github.com/nanohana2199/back_hackathon_mana-nakagawa/db/repositories"
)

type PostService struct {
	PostRepo *repositories.PostRepository
}

// CreatePost は新しい投稿を作成します
func (s *PostService) CreatePost(post models.Post) (*models.Post, error) { // models.Post を使用
	// リポジトリを通じてデータベースに投稿を保存
	return s.PostRepo.CreatePost(post)
}

// GetPosts はすべての投稿を取得します
func (s *PostService) GetPosts() ([]models.Post, error) {
	// リポジトリを通じてデータベースから投稿を取得
	return s.PostRepo.GetPosts()
}
