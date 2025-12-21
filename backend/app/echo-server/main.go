package main

import (
	"context"
	"fmt"
	"log"
	"socialmedia/app/echo-server/handler"
	"socialmedia/app/echo-server/router"
	"socialmedia/repositories/comments_repository"
	"socialmedia/repositories/likes_repository"
	"socialmedia/repositories/posts_repository"
	"socialmedia/repositories/redis_repository"
	"socialmedia/repositories/users_repository"
	"socialmedia/service/comments_service"
	"socialmedia/service/likes_service"
	"socialmedia/service/posts_service"
	"socialmedia/service/users_service"
	"socialmedia/utils/config"
	"socialmedia/utils/database"

	"github.com/go-redis/redis/v8"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Error on config %s", err)
	}

	db, err := database.GetDatabaseConnection(cfg, cfg.DBType)
	if err != nil {
		log.Fatalf("Error on get database connection %s", err)
	}

	redis := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	// repo
	var userRepo users_service.UsersRepo
	var postRepo posts_service.PostsRepo
	var commentRepo comments_service.CommentsRepo
	var likesRepo likes_service.LikesRepo
	redisRepo := redis_repository.NewRedisRepository(redis)
	switch cfg.DBType {
	case "GORM":
		userRepo = users_repository.NewGormUsersRepository(db.Gorm)
		postRepo = posts_repository.NewGormPostRepository(db.Gorm)
		commentRepo = comments_repository.NewGormCommentsRepository(db.Gorm)
		likesRepo = likes_repository.NewGormLikesRepository(db.Gorm)
	case "MONGO":
		userRepo = users_repository.NewMongoUsersRepository(db.Mongo)
		// postRepo = posts_repository.NewMongoPostRepository(db.Mongo)
	case "RAW":
		userRepo = users_repository.NewRawUsersRepository(db.Raw)
		// postRepo = posts_repository.NewRawPostRepository(db.Raw)
	}

	pong, err := redis.Ping(context.Background()).Result()
	fmt.Println("REDIS PING:", pong, err)
	// service
	userService := users_service.NewUsersService(userRepo, cfg.JwtSecret)
	postService := posts_service.NewPostsService(postRepo, redisRepo)
	commentService := comments_service.NewCommentsService(commentRepo)
	likesService := likes_service.NewLikesService(likesRepo)

	// handler
	userHandler := handler.NewUsersHandler(userService)
	postHandler := handler.NewPostsHandler(postService)
	commentHandler := handler.NewCommentsHandler(commentService)
	likesHandler := handler.NewLikesHandler(likesService)

	e := echo.New()

	e.Use(middleware.CORS())
	e.Static("/uploads", "uploads")

	router.Router(e, cfg, userHandler, postHandler, commentHandler, likesHandler)

	log.Println("Successfully connected to the server")

	e.Logger.Fatal(e.Start(":" + cfg.Port))
}
