// handlers/reply_handler.go
package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/nanohana2199/back_hackathon_mana-nakagawa/db/services"
	"log"
	"net/http"
	"strconv"
)

type ReplyHandler struct {
	ReplyService services.ReplyService
}

// NewReplyHandler creates a new ReplyHandler instance
func NewReplyHandler(replyService services.ReplyService) *ReplyHandler {
	return &ReplyHandler{ReplyService: replyService}
}

// CreateReplyHandler handles the POST request to create a new reply
func (h *ReplyHandler) CreateReplyHandler(w http.ResponseWriter, r *http.Request) {
	var replyRequest struct {
		Content string `json:"content"`
		PostID  int64  `json:"postId"`
		UserID  string `json:"user_id"`
	}

	// リクエストボディをデコード
	if err := json.NewDecoder(r.Body).Decode(&replyRequest); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	log.Printf("CreateReplyHandler: %v", replyRequest)

	// リプライを作成
	reply, err := h.ReplyService.CreateReply(replyRequest.Content, replyRequest.PostID, replyRequest.UserID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// レスポンスとしてリプライを返す
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(reply)
}

func (h *ReplyHandler) GetRepliesHandler(w http.ResponseWriter, r *http.Request) {
	// URLパラメータからpost_idを取得
	vars := mux.Vars(r)
	postID, err := strconv.Atoi(vars["post_id"])
	if err != nil {
		http.Error(w, "無効な投稿ID", http.StatusBadRequest)
		return
	}
	log.Printf("postId=%v", postID)
	// リプライを取得
	replies, err := h.ReplyService.GetRepliesByPostID(postID)
	if err != nil {
		http.Error(w, "リプライの取得に失敗しました", http.StatusInternalServerError)
		return
	}

	// レスポンスの形式を設定
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(replies)
}
