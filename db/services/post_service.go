package services

import (
	"cloud.google.com/go/vertexai/genai"
	"fmt"
	"github.com/nanohana2199/back_hackathon_mana-nakagawa/db/external" // Vertex AI関連
	"github.com/nanohana2199/back_hackathon_mana-nakagawa/db/models"   // models.Post をインポート
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

// CheckForHarmfulContent は投稿内容に誹謗中傷が含まれているかをチェックします
func (s *PostService) CheckForHarmfulContent(content string) (genai.Part, error) {
	part, err := external.CheckHarmfulContent(content)
	if err != nil {
		return nil, fmt.Errorf("誹謗中傷チェックに失敗しました: %w", err)
	}
	return part, nil
}
