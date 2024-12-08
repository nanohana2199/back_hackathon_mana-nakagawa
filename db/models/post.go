package models

import "time"

// Post は投稿データを表す構造体
type Post struct {
	ID        int64     `json:"id"`
	Content   string    `json:"content"`
	Author    string    `json:"author"`
	UserID    string    `json:"user_id"`
	ImageURL  *string   `json:"image_url,omitempty"` // 画像がない場合はnull許容
	CreatedAt time.Time `json:"created_at"`
}
