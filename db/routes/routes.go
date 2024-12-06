package routes

import (
	"database/sql"
	"github.com/gorilla/mux"
	"github.com/nanohana2199/back_hackathon_mana-nakagawa/db/handlers"
	"github.com/nanohana2199/back_hackathon_mana-nakagawa/db/repositories"
	"github.com/nanohana2199/back_hackathon_mana-nakagawa/db/services"
)

func SetupRoutes(db *sql.DB) *mux.Router {
	router := mux.NewRouter()

	userRepo := &repositories.UserRepository{DB: db}
	userService := &services.UserService{UserRepo: userRepo}
	userHandler := &handlers.UserHandler{UserService: userService}

	// 投稿関連
	postRepo := &repositories.PostRepository{DB: db}               // PostRepository のインスタンス化
	postService := &services.PostService{PostRepo: postRepo}       // PostService のインスタンス化
	postHandler := &handlers.PostHandler{PostService: postService} // PostHandler のインスタンス化

	likeRepo := &repositories.LikeRepository{DB: db}
	likeService := &services.LikeService{LikeRepo: likeRepo}
	likeHandler := &handlers.LikeHandler{LikeService: likeService}

	replyRepo := repositories.NewReplyRepository(db)
	replyService := services.NewReplyService(replyRepo)
	replyHandler := handlers.NewReplyHandler(replyService)

	// ユーザー関連
	router.HandleFunc("/users", userHandler.CreateUserHandler).Methods("POST")

	// 投稿作成
	router.HandleFunc("/posts", postHandler.CreatePostHandler).Methods("POST")

	// 投稿取得ルート
	router.HandleFunc("/posts", postHandler.GetPostsHandler).Methods("GET")

	router.HandleFunc("/posts/{post_id}/like", likeHandler.AddLikeHandler).Methods("POST")
	router.HandleFunc("/posts/{post_id}/like/count", likeHandler.GetLikeCountHandler).Methods("GET")
	router.HandleFunc("/posts/{post_id}/like/status", likeHandler.CheckLikeStatusHandler).Methods("GET")

	router.HandleFunc("/replies", replyHandler.CreateReplyHandler).Methods("POST")
	// SetupRoutes関数の一部
	router.HandleFunc("/posts/{post_id}/replies", replyHandler.GetRepliesHandler).Methods("GET")

	return router
}
