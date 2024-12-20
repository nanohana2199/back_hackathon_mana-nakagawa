package config

import (
	"database/sql"
	"fmt"
	//"github.com/joho/godotenv"

	_ "github.com/go-sql-driver/mysql"
	"os"

	"log"
)

// DB はデータベース接続を保持するグローバル変数
var DB *sql.DB

// InitDB はデータベース接続を初期化する関数
func InitDB() (*sql.DB, error) {

	var err error
	//err = godotenv.Load()
	//if err != nil {
	//	log.Fatal("Error loading .env file")
	//} // 環境変数からデータベース設定を取得
	mysqlUser := os.Getenv("MYSQL_USER")
	mysqlPassword := os.Getenv("MYSQL_PASSWORD")
	mysqlDatabase := os.Getenv("MYSQL_DATABASE")
	mysqlHost := os.Getenv("MYSQL_HOST")
	//if mysqlHost == "" {
	//	mysqlHost = "127.0.0.1:3306" // TCP接続用のlocalhost
	//}
	log.Printf("MYSQL_USER: %s, MYSQL_PASSWORD:%s, MYSQL_DATABASE: %s, MYSQL_HOST: %s", mysqlUser, mysqlDatabase, mysqlHost)

	if mysqlUser == "" || mysqlPassword == "" || mysqlHost == "" || mysqlDatabase == "" {
		log.Fatalf("Missing required environment variables")
	}

	// データベース接続を確立
	DB, err = sql.Open("mysql", fmt.Sprintf("%s:%s@unix(%s)/%s", mysqlUser, mysqlPassword, mysqlHost, mysqlDatabase))

	if err != nil {
		return nil, err
	}

	// 接続確認
	if err := DB.Ping(); err != nil {
		return nil, err
	}

	log.Println("Database connected successfully")
	return DB, nil
}
