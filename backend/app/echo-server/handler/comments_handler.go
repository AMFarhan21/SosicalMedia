package handler

import (
	"log"
	"net/http"
	"socialmedia/domain"
	"socialmedia/service/comments_service"
	"strconv"
	"strings"

	"github.com/AMFarhan21/fres"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type (
	CommentsHandler struct {
		commentsService comments_service.Service
		validate        *validator.Validate
	}

	CreateCommentInput struct {
		Content  string `json:"content" bson:"content" validate:"required"`
		ImageUrl string `json:"image_url" bson:"image_url"`
	}

	UpdateCommentInput struct {
		Content  string `json:"content" bson:"content" validate:"required"`
		ImageUrl string `json:"image_url" bson:"image_url"`
	}
)

func NewCommentsHandler(commentsService comments_service.Service) *CommentsHandler {
	return &CommentsHandler{
		commentsService: commentsService,
		validate:        validator.New(),
	}
}

func (h CommentsHandler) CreateComment(e echo.Context) error {
	user_id := e.Get("id").(string)
	id := e.Param("id")
	post_id, _ := strconv.Atoi(id)

	var request CreateCommentInput

	if err := e.Bind(&request); err != nil {
		log.Printf("Error on CreateComment request body: %v", err.Error())
		return e.JSON(http.StatusBadRequest, fres.Response.StatusBadRequest(err.Error()))
	}

	if err := h.validate.Struct(request); err != nil {
		log.Printf("Error on CreateComment validation: %v", err.Error())
		return e.JSON(http.StatusBadRequest, fres.Response.StatusBadRequest(err.Error()))
	}

	post, err := h.commentsService.CreateComment(domain.Comments{
		UserID:   user_id,
		PostID:   int64(post_id),
		Content:  request.Content,
		ImageUrl: request.ImageUrl,
	})
	if err != nil {
		log.Printf("Error on CreateComment internal: %v", err.Error())
		return e.JSON(http.StatusInternalServerError, fres.Response.StatusInternalServerError(http.StatusInternalServerError))
	}

	return e.JSON(http.StatusCreated, fres.Response.StatusCreated(post))
}

func (h CommentsHandler) GetAllComments(e echo.Context) error {
	id := e.Param("id")
	post_id, _ := strconv.Atoi(id)

	comments, err := h.commentsService.GetAllComments(int64(post_id))
	if err != nil {
		log.Printf("Error on GetAllComments internal: %v", err.Error())
		return e.JSON(http.StatusInternalServerError, fres.Response.StatusInternalServerError(http.StatusInternalServerError))
	}

	return e.JSON(http.StatusOK, fres.Response.StatusOK(comments))
}

func (h CommentsHandler) GetCommentByID(e echo.Context) error {
	strCommentID := e.Param("comment_id")
	comment_id, _ := strconv.Atoi(strCommentID)

	comment, err := h.commentsService.GetCommentByID(int64(comment_id))
	if err != nil {
		if strings.Contains(err.Error(), "found") {
			log.Printf("Error on GetCommentByID request: %v", err.Error())
			return e.JSON(http.StatusNotFound, fres.Response.StatusNotFound(err.Error()))
		}
		log.Printf("Error on GetCommentByID internal: %v", err.Error())
		return e.JSON(http.StatusInternalServerError, fres.Response.StatusInternalServerError(http.StatusInternalServerError))
	}

	return e.JSON(http.StatusOK, fres.Response.StatusOK(comment))
}

func (h CommentsHandler) UpdateComment(e echo.Context) error {
	user_id := e.Get("id").(string)

	id := e.Param("id")
	post_id, _ := strconv.Atoi(id)

	strCommentID := e.Param("comment_id")
	comment_id, _ := strconv.Atoi(strCommentID)

	var request UpdateCommentInput

	if err := e.Bind(&request); err != nil {
		log.Printf("Error on UpdateComment request body: %v", err.Error())
		return e.JSON(http.StatusBadRequest, fres.Response.StatusBadRequest(err.Error()))
	}

	if err := h.validate.Struct(request); err != nil {
		log.Printf("Error on UpdateComment validation: %v", err.Error())
		return e.JSON(http.StatusBadRequest, fres.Response.StatusBadRequest(err.Error()))
	}

	err := h.commentsService.UpdateComment(domain.Comments{
		ID:       int64(comment_id),
		UserID:   user_id,
		PostID:   int64(post_id),
		Content:  request.Content,
		ImageUrl: request.ImageUrl,
	})
	if err != nil {
		if strings.Contains(err.Error(), "no rows") || strings.Contains(err.Error(), "found") {
			log.Printf("Error on UpdateComment request: %v", err.Error())
			return e.JSON(http.StatusNotFound, fres.Response.StatusNotFound(err.Error()))
		}

		log.Printf("Error on UpdateComment internal: %v", err.Error())
		return e.JSON(http.StatusInternalServerError, fres.Response.StatusInternalServerError(http.StatusInternalServerError))
	}

	return e.JSON(http.StatusCreated, fres.Response.StatusCreated("Post updated successfully"))
}

func (h CommentsHandler) DeleteComment(e echo.Context) error {
	user_id := e.Get("id").(string)

	strCommentID := e.Param("comment_id")
	comment_id, _ := strconv.Atoi(strCommentID)

	err := h.commentsService.DeleteComment(int64(comment_id), user_id)
	if err != nil {
		if strings.Contains(err.Error(), "found") {
			log.Printf("Error on DeleteComment request: %v", err.Error())
			return e.JSON(http.StatusNotFound, fres.Response.StatusNotFound(err.Error()))
		}
		log.Printf("Error on DeleteComment internal: %v", err.Error())
		return e.JSON(http.StatusInternalServerError, fres.Response.StatusInternalServerError(http.StatusInternalServerError))
	}

	return e.JSON(http.StatusOK, fres.Response.StatusOK("Post deleted successfully"))
}
