package main

import (
	"log"
	"socialmedia/app/echo-server/handler"
	"socialmedia/app/echo-server/router"
	"socialmedia/repositories/posts_repository"
	"socialmedia/repositories/users_repository"
	"socialmedia/service/posts_service"
	"socialmedia/service/users_service"
	"socialmedia/utils/config"
	"socialmedia/utils/database"

	"github.com/labstack/echo/v4"
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

	// repo
	var userRepo users_service.UsersRepo
	var postRepo posts_service.PostsRepo
	switch cfg.DBType {
	case "GORM":
		userRepo = users_repository.NewGormUsersRepository(db.Gorm)
		postRepo = posts_repository.NewGormPostRepository(db.Gorm)
	case "MONGO":
		userRepo = users_repository.NewMongoUsersRepository(db.Mongo)
		// postRepo = posts_repository.NewMongoPostRepository(db.Mongo)
	case "RAW":
		userRepo = users_repository.NewRawUsersRepository(db.Raw)
		// postRepo = posts_repository.NewRawPostRepository(db.Raw)
	}

	// service
	userService := users_service.NewUsersService(userRepo, cfg.JwtSecret)
	postService := posts_service.NewPostsService(postRepo)

	// handler
	userHandler := handler.NewUsersHandler(userService)
	postHandler := handler.NewPostsHandler(postService)

	e := echo.New()

	router.Router(e, cfg, userHandler, postHandler)

	log.Println("Successfully connected to the server")

	e.Logger.Fatal(e.Start(":" + cfg.Port))
}
