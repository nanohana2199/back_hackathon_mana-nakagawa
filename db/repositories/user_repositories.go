package repositories

import (
	"database/sql"
	"github.com/nanohana2199/back_hackathon_mana-nakagawa/db/models"
)

type UserRepository struct {
	DB *sql.DB
}

// Create adds a new user to the database.
func (repo *UserRepository) Create(user *models.User) error {
	// ユーザーの作成（user_id, email, usernameを挿入）
	query := "INSERT INTO users (user_id, email, username) VALUES (?, ?, ?)"
	_, err := repo.DB.Exec(query, user.UserID, user.Email, user.Username) // user_idを挿入
	return err
}

// FindByEmail finds a user by email in the database.
func (repo *UserRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User
	query := "SELECT user_id, email, username, created_at FROM users WHERE email = ?"
	err := repo.DB.QueryRow(query, email).Scan(&user.UserID, &user.Email, &user.Username, &user.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // ユーザーが見つからない場合はnilを返す
		}
		return nil, err
	}
	return &user, nil
}

// FindByUserID finds a user by user_id (Firebase UID) in the database.
func (repo *UserRepository) FindByUserID(userID string) (*models.User, error) {
	var user models.User
	query := "SELECT user_id, email, username, created_at FROM users WHERE user_id = ?"
	err := repo.DB.QueryRow(query, userID).Scan(&user.UserID, &user.Email, &user.Username, &user.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // ユーザーが見つからない場合はnilを返す
		}
		return nil, err
	}
	return &user, nil
}
