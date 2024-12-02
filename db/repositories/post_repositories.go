package repositories

import (
	"database/sql"
	"github.com/nanohana2199/back_hackathon_mana-nakagawa/db/models"
)

type PostRepository struct {
	DB *sql.DB
}

// CreatePost はデータベースに新しい投稿を追加します
func (r *PostRepository) CreatePost(post models.Post) (*models.Post, error) { // models.Post型を使用
	// SQLクエリの実行
	query := `INSERT INTO posts (content) VALUES (?)`
	result, err := r.DB.Exec(query, post.Content)
	if err != nil {
		return nil, err
	}

	// 投稿IDを取得
	postID, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	post.ID = postID
	return &post, nil
}

// GetPosts はすべての投稿をデータベースから取得します
func (r *PostRepository) GetPosts() ([]models.Post, error) {
	var posts []models.Post

	query := "SELECT id, content FROM posts"
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var post models.Post
		if err := rows.Scan(&post.ID, &post.Content); err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	return posts, nil
}
