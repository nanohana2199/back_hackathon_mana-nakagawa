package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/nanohana2199/back_hackathon_mana-nakagawa/db/models"
	"github.com/nanohana2199/back_hackathon_mana-nakagawa/db/services"
	"net/http"
)

// UserHandler handles user-related requests
type UserHandler struct {
	UserService *services.UserService
}

// CreateUserHandler handles the user creation HTTP request
// ここでuser_idとしてFirebase UIDを受け取ります
func (u *UserHandler) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var userRequest struct {
		UserID   string `json:"user_id"`  // Firebase UIDを受け取る
		Username string `json:"username"` // ユーザー名
		Email    string `json:"email"`    // メールアドレス
	}

	// リクエストボディをデコード
	if err := json.NewDecoder(r.Body).Decode(&userRequest); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// モデルに変換
	user := models.User{
		UserID:   userRequest.UserID,
		Username: userRequest.Username,
		Email:    userRequest.Email,
	}

	// ユーザー作成サービスを呼び出し
	err := u.UserService.CreateUser(&user)
	if err != nil {
		http.Error(w, fmt.Sprintf("Unable to create user: %v", err), http.StatusInternalServerError)
		return
	}

	// ユーザー作成成功のレスポンス
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}
