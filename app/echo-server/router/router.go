package router

import (
	"socialmedia/app/echo-server/handler"
	"socialmedia/app/echo-server/middleware"
	"socialmedia/utils/config"

	"github.com/labstack/echo/v4"
)

func Router(e *echo.Echo, cfg *config.Config, userHandler *handler.UsersHandler, postHandler *handler.PostsHandler) {
	jwtMiddleware := middleware.JWTMiddleware(cfg.JwtSecret)
	adminAccess := middleware.ACLMiddleware(map[string]bool{
		"admin": true,
	})

	// userAccess := middleware.ACLMiddleware(map[string]bool{
	// 	"user": true,
	// })

	userNAdminAccess := middleware.ACLMiddleware(map[string]bool{
		"admin": true,
		"user":  true,
	})

	api := e.Group("/api/v1")

	auth := api.Group("/auth")
	auth.POST("/register", userHandler.RegisterNewUser)
	auth.POST("/login", userHandler.LoginUser)

	users := api.Group("/users")
	users.POST("", userHandler.GetAllUsers)
	users.GET("", userHandler.GetMe, jwtMiddleware)
	users.GET("/:id", userHandler.GetUserByID, jwtMiddleware, adminAccess)

	posts := api.Group("/posts", jwtMiddleware)
	posts.POST("", postHandler.CreatePost, userNAdminAccess)
	posts.POST("/feed", postHandler.GetAllPost, userNAdminAccess)
	posts.GET("/:id", postHandler.GetPostByID, userNAdminAccess)
	posts.PUT("/:id", postHandler.UpdatePost, userNAdminAccess)
	posts.DELETE("/:id", postHandler.DeletePost, userNAdminAccess)
}
