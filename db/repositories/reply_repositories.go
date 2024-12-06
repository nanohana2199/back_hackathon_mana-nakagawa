package repositories

import (
	"database/sql"
	"errors"

	_ "github.com/go-sql-driver/mysql"
	"github.com/nanohana2199/back_hackathon_mana-nakagawa/db/models"
	"log"
)

type ReplyRepository interface {
	// 返信作成
	CreateReply(reply *models.Reply) (*models.Reply, error)
	// 投稿IDに関連するリプライを取得
	FindRepliesByPostID(postID int) ([]models.Reply, error)
}

// ReplyRepositoryImplはReplyRepositoryインターフェースの実装
type ReplyRepositoryImpl struct {
	DB *sql.DB
}

// NewReplyRepositoryは新しいReplyRepositoryインスタンスを作成します
func NewReplyRepository(db *sql.DB) ReplyRepository {
	return &ReplyRepositoryImpl{DB: db}
}

// CreateReplyはデータベースに新しいリプライを挿入します
func (r *ReplyRepositoryImpl) CreateReply(reply *models.Reply) (*models.Reply, error) {
	query := `INSERT INTO replies (content, post_id, user_id) VALUES (?, ?, ?)`
	result, err := r.DB.Exec(query, reply.Content, reply.PostID, reply.UserID)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	reply.ID = id
	return reply, nil
}

// FindRepliesByPostIDは指定された投稿IDに関連するリプライをデータベースから取得します
func (r *ReplyRepositoryImpl) FindRepliesByPostID(postID int) ([]models.Reply, error) {
	query := `
		SELECT 
			replies.id AS reply_id,
			replies.content AS reply_content,
			replies.post_id,
			replies.user_id,
			users.username AS author
		FROM 
			replies
		JOIN 
			users
		ON 
			replies.user_id = users.user_id
		WHERE 
			replies.post_id = ?
	`

	rows, err := r.DB.Query(query, postID)
	if err != nil {
		log.Printf("Error while querying replies for postID %v: %v", postID, err)
		return nil, err
	}
	defer rows.Close()
	log.Printf("rows=%v", rows)

	var replies []models.Reply

	for rows.Next() {
		var reply models.Reply
		err := rows.Scan(&reply.ID, &reply.Content, &reply.PostID, &reply.UserID, &reply.Author)
		if err != nil {
			// エラーが発生した場合、ログを出力して詳細を確認
			log.Printf("rows.Scan error: %v", err)
			return nil, err
		}
		replies = append(replies, reply)
	}

	// rows.Err() をチェックして反復処理中のエラーを確認
	if err := rows.Err(); err != nil {
		log.Printf("rows iteration error: %v", err)
		return nil, err
	}

	// デバッグ情報をログに出力
	log.Printf("Retrieved replies: %+v", replies)
	return replies, nil

	// リプライがない場合
	if len(replies) == 0 {
		return nil, errors.New("リプライはありません")
	}

	return replies, nil
}
