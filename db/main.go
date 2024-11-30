package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"db/controller"
	"db/dao"
	"db/usecase"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	//err := godotenv.Load(".env")

	// もし err がnilではないなら、"読み込み出来ませんでした"が出力されます。
	//if err != nil {
	//	fmt.Printf("読み込み出来ませんでした: %v", err)
	//}
	mysqlUser := os.Getenv("MYSQL_USER")
	mysqlPassword := os.Getenv("MYSQL_PASSWORD")
	mysqlDatabase := os.Getenv("MYSQL_DATABASE")
	mysqlHost := os.Getenv("MYSQL_HOST") // Cloud SQLの接続名

	// Cloud SQL Auth Proxyを利用した接続
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@%s/%s",
		mysqlUser,
		mysqlPassword,
		mysqlHost,
		mysqlDatabase,
	))
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer db.Close()

	dao := &dao.UserDAO{DB: db}
	registerUseCase := &usecase.RegisterUserUseCase{DAO: dao}
	searchUseCase := &usecase.SearchUserUseCase{DAO: dao}

	http.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			controller := &controller.RegisterUserController{UseCase: registerUseCase}
			controller.Handle(w, r)
		case http.MethodGet:
			controller := &controller.SearchUserController{UseCase: searchUseCase}
			controller.Handle(w, r)
		default:
			w.WriteHeader(http.StatusBadRequest)
		}
	})

	// ポートの設定（環境変数PORTに基づく）
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // デフォルトポート
	}
	log.Printf("Server started at :%s", port)

	// サーバーの起動
	go func() {
		if err := http.ListenAndServe(":"+port, nil); err != nil {
			log.Fatalf("server error: %v", err)
		}
	}()

	// Graceful shutdown
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGTERM, syscall.SIGINT)
	<-sig
	log.Println("Shutting down server...")
}
