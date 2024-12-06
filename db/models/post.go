package models

// Post は投稿データを表す構造体
type Post struct {
	ID      int64  `json:"id"`
	Content string `json:"content"`
	Author  string `json:"author"`
	UserID  string `json:"user_id"`
}
