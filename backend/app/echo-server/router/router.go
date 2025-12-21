package router

import (
	"log"
	"net/http"
	"socialmedia/app/echo-server/handler"
	"socialmedia/app/echo-server/middleware"
	"socialmedia/utils/config"
	"time"

	"github.com/AMFarhan21/fres"
	"github.com/labstack/echo/v4"
	"github.com/robfig/cron/v3"
)

func Router(e *echo.Echo, c *cron.Cron, cfg *config.Config, userHandler *handler.UsersHandler, postHandler *handler.PostsHandler, commentHandler *handler.CommentsHandler, likesHandler *handler.LikesHandler) {
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

	c.AddFunc("*/10 * * * *", func() {
		response, err := http.Get(cfg.DeploymentURL + "/health")
		if err != nil {
			log.Print("Ping failed", err.Error())
			return
		}

		response.Body.Close()
		log.Print("Pinged at", time.Now())
	})

	e.GET(cfg.DeploymentURL+"/health", func(e echo.Context) error {
		return e.JSON(http.StatusOK, fres.Response.StatusOK("health ok"))
	})

	auth := api.Group("/auth")
	auth.POST("/register", userHandler.RegisterNewUser)
	auth.POST("/login", userHandler.LoginUser)

	users := api.Group("/users")
	users.POST("", userHandler.GetAllUsers)
	users.GET("", userHandler.GetMe, jwtMiddleware)
	users.GET("/:id", userHandler.GetUserByID, jwtMiddleware, adminAccess)

	posts := api.Group("/posts", jwtMiddleware)
	posts.POST("", postHandler.CreatePost, userNAdminAccess)
	posts.GET("", postHandler.GetAllPost, userNAdminAccess)
	posts.GET("/:id", postHandler.GetPostByID, userNAdminAccess)
	posts.PUT("/:id", postHandler.UpdatePost, userNAdminAccess)
	posts.DELETE("/:id", postHandler.DeletePost, userNAdminAccess)

	comments := posts.Group("/:id/comments")
	comments.POST("", commentHandler.CreateComment, userNAdminAccess)
	comments.GET("", commentHandler.GetAllComments, userNAdminAccess)
	comments.GET("/:comment_id", commentHandler.GetCommentByID, userNAdminAccess)
	comments.PUT("/:comment_id", commentHandler.UpdateComment, userNAdminAccess)
	comments.DELETE("/:comment_id", commentHandler.DeleteComment, userNAdminAccess)

	likes := api.Group("/likes", jwtMiddleware)
	likes.POST("", likesHandler.Likes)
}

// func Router(e *echo.Echo, cfg *config.Config, userHandler *handler.UsersHandler, postHandler *handler.PostsHandler, commentHandler *handler.CommentsHandler) {

// 	api := e.Group("/api/v1")

// 	auth := api.Group("/auth")
// 	auth.POST("/register", userHandler.RegisterNewUser)
// 	auth.POST("/login", userHandler.LoginUser)

// 	users := api.Group("/users")
// 	users.POST("", userHandler.GetAllUsers)
// 	users.GET("", userHandler.GetMe)
// 	users.GET("/:id", userHandler.GetUserByID)

// 	posts := api.Group("/posts")
// 	posts.POST("", postHandler.CreatePost)
// 	posts.GET("", postHandler.GetAllPost)
// 	posts.GET("/:id", postHandler.GetPostByID)
// 	posts.PUT("/:id", postHandler.UpdatePost)
// 	posts.DELETE("/:id", postHandler.DeletePost)

// 	comments := posts.Group("/:id/comments")
// 	comments.POST("", commentHandler.CreateComment)
// 	comments.GET("", commentHandler.GetAllComments)
// 	comments.GET("/:comment_id", commentHandler.GetCommentByID)
// 	comments.PUT("/:comment_id", commentHandler.UpdateComment)
// 	comments.DELETE("/:comment_id", commentHandler.DeleteComment)
// }
