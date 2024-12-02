package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/nanohana2199/back_hackathon_mana-nakagawa/db/services"
	"io/ioutil"
	"net/http"
	"strconv"
)

type LikeHandler struct {
	LikeService *services.LikeService
}

func (handler *LikeHandler) AddLikeHandler(w http.ResponseWriter, r *http.Request) {
	// パスパラメータからpost_idを取得
	vars := mux.Vars(r)
	postIdStr := vars["post_id"]
	postId, err := strconv.Atoi(postIdStr)
	if err != nil {
		// post_idが整数に変換できなければエラー
		http.Error(w, "Invalid post ID", http.StatusBadRequest)
		return
	}

	// ログを追加してpost_idの値を確認
	fmt.Println("Received post_id:", postIdStr)

	// リクエストボディを読み取る
	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}

	// ボディの内容を表示
	fmt.Println("Request Body:", string(bodyBytes))

	// ボディをデコードするために再度r.Bodyを設定
	r.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

	// リクエストボディからuser_idを取得
	var likeRequest struct {
		UserID string `json:"user_id"` // Firebase UIDは文字列型で送られる
	}
	if err := json.NewDecoder(r.Body).Decode(&likeRequest); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// 受け取ったuser_idの確認ログ
	fmt.Println("Received user_id:", likeRequest.UserID)

	// user_idはstring型のままで渡す
	err = handler.LikeService.AddLike(postId, likeRequest.UserID) // user_idはstring型で渡す
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 成功レスポンス
	w.WriteHeader(http.StatusOK)

}

func (h *LikeHandler) GetLikeCountHandler(w http.ResponseWriter, r *http.Request) {
	// URLパラメータから投稿IDを取得
	postID := mux.Vars(r)["post_id"]
	fmt.Println("Received post_id:", postID) // 追加したデバッグ出力

	// 投稿IDが空の場合のチェック
	if postID == "" {
		http.Error(w, "投稿IDが指定されていません", http.StatusBadRequest)
		return
	}

	// いいね数を取得
	likeCount, err := h.LikeService.GetLikeCount(postID)
	if err != nil {
		// エラーログを追加してエラーの詳細を表示
		fmt.Println("Error fetching like count:", err)
		http.Error(w, "いいね数の取得に失敗しました", http.StatusInternalServerError)
		return
	}

	// JSON形式でレスポンスを返す
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(map[string]int{"count": likeCount})
	if err != nil {
		// エンコードエラーのログ出力
		fmt.Println("Error encoding JSON response:", err)
		http.Error(w, "レスポンスの作成に失敗しました", http.StatusInternalServerError)
		return
	}

	// 成功レスポンスを返す
	w.WriteHeader(http.StatusOK)
}
