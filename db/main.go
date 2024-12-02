package main

import (
	"database/sql"
	"github.com/nanohana2199/back_hackathon_mana-nakagawa/db/config"

	"github.com/nanohana2199/back_hackathon_mana-nakagawa/db/routes"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func CORSMiddlewareProd(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		// プリフライトリクエストの応答
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		// 次のミドルウェアまたはハンドラを呼び出す
		next.ServeHTTP(w, r)
	})
}
func main() {
	// データベース接続の初期化
	var db *sql.DB
	var err error
	if db, err = config.InitDB(); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
		return
	}

	// ルーティング設定
	router := routes.SetupRoutes(db)

	corsRouter := CORSMiddlewareProd(router)
	// サーバーの起動
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // デフォルトポート
	}
	log.Printf("Server started at :%s", port)

	// サーバー起動
	go func() {
		if err := http.ListenAndServe(":"+port, corsRouter); err != nil {
			log.Fatalf("Server error: %v", err)
		}
	}()

	// Graceful Shutdown（終了シグナルを受け取ったときの処理）
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGTERM, syscall.SIGINT)
	<-sig
	log.Println("Shutting down server...")
}
