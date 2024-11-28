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
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("/Users/mana/curriculum_6_mana-nakagawa/mysql/.env")
	if err != nil {
		log.Fatalf("failed to load .env file: %v", err)
	}

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@(localhost:3306)/%s",
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_DATABASE"),
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

	go func() {
		log.Println("Server started at :8000")
		if err := http.ListenAndServe(":8000", nil); err != nil {
			log.Fatalf("server error: %v", err)
		}
	}()

	// Graceful shutdown
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGTERM, syscall.SIGINT)
	<-sig
	log.Println("Shutting down server...")
}
