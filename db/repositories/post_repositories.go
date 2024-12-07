package repositories

import (
	"database/sql"
	"github.com/nanohana2199/back_hackathon_mana-nakagawa/db/models"
	"log"
	"time"
)

type PostRepository struct {
	DB *sql.DB
}

// CreatePost はデータベースに新しい投稿を追加します
func (r *PostRepository) CreatePost(post models.Post) (*models.Post, error) { // models.Post型を使用
	// SQLクエリの実行
	query := `INSERT INTO posts (content,user_id) VALUES (?, ?)`
	result, err := r.DB.Exec(query, post.Content, post.UserID)
	if err != nil {
		log.Println("Error executing CreatePost query:", err)
		return nil, err
	}

	// 投稿IDを取得
	postID, err := result.LastInsertId()
	if err != nil {
		log.Println("Error fetching last insert ID:", err)
		return nil, err
	}

	post.ID = postID
	return &post, nil
}

// GetPosts はすべての投稿をデータベースから取得します
func (r *PostRepository) GetPosts() ([]models.Post, error) {
	var posts []models.Post

	query := `
    SELECT 
        posts.id AS post_id, 
        posts.content AS post_content, 
        posts.user_id, 
        users.username AS author,
        posts.created_at
    FROM 
        posts
    JOIN 
        users 
    ON 
        posts.user_id = users.user_id
       `
	rows, err := r.DB.Query(query)
	if err != nil {
		log.Println("Error querying posts:", err) // エラーログを記録

		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var post models.Post
		var createdAtStr string

		if err := rows.Scan(&post.ID, &post.Content, &post.UserID, &post.Author, &createdAtStr); err != nil {
			log.Println("Error scanning row:", err) // スキャン時のエラーを記録
			return nil, err
		}
		// createdAtStrをtime.Timeに変換
		parsedTime, err := time.Parse("2006-01-02 15:04:05", createdAtStr)
		if err != nil {
			log.Println("Error parsing created_at:", err)
			return nil, err
		}
		post.CreatedAt = parsedTime
		posts = append(posts, post)
	}

	return posts, nil
}

func (r *PostRepository) GetPostsByUserID(userID string) ([]models.Post, error) {
	var posts []models.Post

	query := `
    SELECT 
        posts.id AS post_id, 
        posts.content AS post_content, 
        posts.user_id, 
        users.username AS author,
        posts.created_at
    FROM 
        posts
    JOIN 
        users 
    ON 
        posts.user_id = users.user_id
    WHERE 
        posts.user_id = ?`

	rows, err := r.DB.Query(query, userID)
	if err != nil {
		log.Println("Error querying posts by user ID:", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var post models.Post
		var createdAtStr string

		if err := rows.Scan(&post.ID, &post.Content, &post.UserID, &post.Author, &createdAtStr); err != nil {
			log.Println("Error scanning row for user ID:", err)
			return nil, err
		}

		parsedTime, err := time.Parse("2006-01-02 15:04:05", createdAtStr)
		if err != nil {
			log.Println("Error parsing created_at:", err)
			return nil, err
		}
		post.CreatedAt = parsedTime
		posts = append(posts, post)
	}

	return posts, nil
}
