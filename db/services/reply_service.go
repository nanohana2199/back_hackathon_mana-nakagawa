// services/reply_service.go
package services

import (
	"cloud.google.com/go/vertexai/genai"
	"errors"
	"fmt"
	"github.com/nanohana2199/back_hackathon_mana-nakagawa/db/external"
	"github.com/nanohana2199/back_hackathon_mana-nakagawa/db/models"
	"github.com/nanohana2199/back_hackathon_mana-nakagawa/db/repositories"
	"log"
)

type ReplyService interface {
	// リプライ作成
	CreateReply(content string, postID int64, userID string) (*models.Reply, error)
	// 投稿IDに関連するリプライを取得
	GetRepliesByPostID(postID int) ([]models.Reply, error)
	// リプライ内容の誹謗中傷チェック
	CheckForHarmfulContent(content string) (genai.Part, error)
}

type ReplyServiceImpl struct {
	ReplyRepo repositories.ReplyRepository
}

// NewReplyService は新しい ReplyService インスタンスを作成します
func NewReplyService(replyRepo repositories.ReplyRepository) ReplyService {
	return &ReplyServiceImpl{ReplyRepo: replyRepo}
}

// CreateReply は新しいリプライを作成する処理を担当します
func (s *ReplyServiceImpl) CreateReply(content string, postID int64, userID string) (*models.Reply, error) {
	// バリデーションなどの追加ロジックをここで実装可能
	reply := &models.Reply{
		Content: content,
		PostID:  postID,
		UserID:  userID,
	}

	return s.ReplyRepo.CreateReply(reply)
}

// GetRepliesByPostID は指定された投稿IDに関連するリプライを取得します
func (s *ReplyServiceImpl) GetRepliesByPostID(postID int) ([]models.Reply, error) {
	replies, err := s.ReplyRepo.FindRepliesByPostID(postID)
	if err != nil {
		return nil, errors.New("リプライの取得に失敗しました")
	}
	return replies, nil
}

func (s *ReplyServiceImpl) CheckForHarmfulContent(content string) (genai.Part, error) {
	part, err := external.CheckHarmfulContent(content)
	if err != nil {
		log.Printf("CheckForHarmfulContent failed: %v", err)
		return nil, fmt.Errorf("誹謗中傷チェックに失敗しました: %w", err)
	}
	return part, nil
}
