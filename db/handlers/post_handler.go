package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/nanohana2199/back_hackathon_mana-nakagawa/db/models" // models.Post をインポート
	"github.com/nanohana2199/back_hackathon_mana-nakagawa/db/services"
	"log"
	"net/http"
)

type PostHandler struct {
	PostService *services.PostService
}

// CreatePostHandler は新しい投稿を作成するハンドラー
func (h *PostHandler) CreatePostHandler(w http.ResponseWriter, r *http.Request) {
	var post models.Post // services.Post ではなく models.Post を使う

	// リクエストボディを解析して投稿データを構造体にバインド
	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		http.Error(w, "無効なデータ", http.StatusBadRequest)
		return
	}

	// 投稿内容に誹謗中傷が含まれていないかを確認
	part, err := h.PostService.CheckForHarmfulContent(post.Content)
	if err != nil {
		http.Error(w, "投稿内容のチェックに失敗しました", http.StatusInternalServerError)
		return
	}

	partStr := fmt.Sprintf("%s", part)
	log.Printf("CheckForHarmfulContent result: %s", partStr)

	// 不適切な投稿内容の場合、エラーを返す
	if partStr == "yes\n" {
		http.Error(w, "投稿内容が不適切です", http.StatusBadRequest)
		log.Printf("投稿が不適切と判断されました: %v", post.Content)
		return
	}

	log.Printf("hello")

	// 投稿を作成
	_, err = h.PostService.CreatePost(post)
	if err != nil {
		http.Error(w, "投稿作成に失敗しました", http.StatusInternalServerError)
		return
	}

	// Content-Type を application/json に設定
	w.Header().Set("Content-Type", "application/json")
	// 作成した投稿をレスポンスとして返す
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(part)
	log.Printf("part=%v", part)
}

// GetPostsHandler はすべての投稿を取得するハンドラー
func (h *PostHandler) GetPostsHandler(w http.ResponseWriter, r *http.Request) {
	posts, err := h.PostService.GetPosts()
	if err != nil {
		http.Error(w, "投稿取得に失敗しました", http.StatusInternalServerError)
		return
	}

	// Content-Type を application/json に設定
	w.Header().Set("Content-Type", "application/json")
	// 取得した投稿をレスポンスとして返す
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(posts)
}
