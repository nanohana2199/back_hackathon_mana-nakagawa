package services

import (
	"fmt"
	"github.com/nanohana2199/back_hackathon_mana-nakagawa/db/models"
	"github.com/nanohana2199/back_hackathon_mana-nakagawa/db/repositories"
	"log"
)

type UserService struct {
	UserRepo *repositories.UserRepository
}

// CreateUser creates a new user in the system using user_id (Firebase UID).
func (s *UserService) CreateUser(user *models.User) error {
	// ユーザーがすでに存在するかチェック（user_idで検索）
	existingUser, _ := s.UserRepo.FindByUserID(user.UserID)
	if existingUser != nil {
		log.Printf("User with user_id %s already exists", user.UserID)
		return fmt.Errorf("user with user_id %s already exists", user.UserID)
	}

	// ユーザーの作成
	err := s.UserRepo.Create(user)
	if err != nil {
		log.Printf("Error creating user %s", user.UserID)
		log.Printf("error creating user: %v", err)
		return fmt.Errorf("error creating user: %v", err)
	}
	return nil
}
