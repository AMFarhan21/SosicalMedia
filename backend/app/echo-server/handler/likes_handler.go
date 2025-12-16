package handler

import (
	"log"
	"net/http"
	"socialmedia/domain"
	"socialmedia/service/likes_service"

	"github.com/AMFarhan21/fres"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type (
	LikesHandler struct {
		likesService likes_service.Service
		validate     *validator.Validate
	}

	LikesRequest struct {
		PostId    *int `json:"post_id" bson:"post_id"`
		CommentId *int `json:"comment_id" bson:"comment_id"`
	}
)

func NewLikesHandler(service likes_service.Service) *LikesHandler {
	return &LikesHandler{
		likesService: service,
		validate:     validator.New(),
	}
}

func (h *LikesHandler) Likes(e echo.Context) error {
	user_id := e.Get("id").(string)

	var request LikesRequest

	if err := e.Bind(&request); err != nil {
		log.Printf("Error on Likes request body: %v", err.Error())
		return e.JSON(http.StatusBadRequest, fres.Response.StatusBadRequest(err.Error()))
	}

	if err := h.validate.Struct(request); err != nil {
		log.Printf("Error on Likes validation: %v", err.Error())
		return e.JSON(http.StatusBadRequest, fres.Response.StatusBadRequest(err.Error()))
	}

	like, err := h.likesService.Likes(domain.Likes{
		UserId:    user_id,
		PostId:    request.PostId,
		CommentId: request.CommentId,
	})
	if err != nil {
		log.Printf("Error on Likes internal: %v", err.Error())
		return e.JSON(http.StatusInternalServerError, fres.Response.StatusInternalServerError(http.StatusInternalServerError))
	}

	if like.ID != 0 {
		return e.JSON(http.StatusCreated, fres.Response.StatusCreated(like))
	}

	return e.JSON(http.StatusOK, fres.Response.StatusOK("Unliked"))

}
