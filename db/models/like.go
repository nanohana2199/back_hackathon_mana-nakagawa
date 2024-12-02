package models

type Like struct {
	ID     int    `json:"id"`
	PostID int    `json:"post_id"`
	UserID string `json:"user_id"`
}
