// models/reply.go
package models

import "time"

// Reply represents a reply to a post
type Reply struct {
	ID        int64     `json:"id"`
	Content   string    `json:"content"`
	PostID    int64     `json:"post_id"`
	UserID    string    `json:"user_id"`
	Author    string    `json:"author"` // 作成者のユーザー名
	CreatedAt time.Time `json:"created_at"`
}
