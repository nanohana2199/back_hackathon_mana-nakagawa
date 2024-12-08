// handlers/reply_handler.go
package handlers

import (
	"encoding/json"
	"fmt"
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
		http.Error(w, "リクエストデータの解析に失敗しました: "+err.Error(), http.StatusBadRequest)
		log.Printf("Error decoding request body: %v", err)
		return
	}
	log.Printf("CreateReplyHandler: %v", replyRequest)

	reply, err := h.ReplyService.CheckForHarmfulContent(replyRequest.Content)
	if err != nil {
		http.Error(w, "投稿内容のチェックに失敗しました: "+err.Error(), http.StatusInternalServerError)
		log.Printf("CheckForHarmfulContent failed: %v", err)
		return
	}

	partStr := fmt.Sprintf("%s", reply)
	log.Printf("CheckForHarmfulContent result: %s", partStr)

	if partStr == "yes\n" {
		http.Error(w, "不適切な内容が検出されました", http.StatusForbidden)
		log.Printf("Harmful content detected: %v", replyRequest.Content)
		return
	}

	// リプライを作成
	log.Printf("Attempting to create reply for postID: %d, userID: %s", replyRequest.PostID, replyRequest.UserID)
	_, err = h.ReplyService.CreateReply(replyRequest.Content, replyRequest.PostID, replyRequest.UserID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// レスポンスとしてリプライを返す
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	log.Printf("Reply successfully created: %v", replyRequest.Content)
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
