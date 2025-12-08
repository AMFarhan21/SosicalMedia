package handler

import (
	"log"
	"net/http"
	"socialmedia/domain"
	"socialmedia/service/posts_service"
	"strconv"
	"strings"

	"github.com/AMFarhan21/fres"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type (
	PostsHandler struct {
		postsService posts_service.Service
		validate     *validator.Validate
	}

	CreateInput struct {
		Content  string `json:"content" bson:"content" validate:"required"`
		ImageUrl string `json:"image_url" bson:"image_url"`
	}

	GetAllPostInput struct {
		Page  int `json:"page" validate:"required"`
		Limit int `json:"limit" validate:"required"`
	}

	UpdateInput struct {
		Content  string `json:"content" bson:"content" validate:"required"`
		ImageUrl string `json:"image_url" bson:"image_url"`
	}
)

func NewPostsHandler(postsService posts_service.Service) *PostsHandler {
	return &PostsHandler{
		postsService: postsService,
		validate:     validator.New(),
	}
}

func (h PostsHandler) CreatePost(e echo.Context) error {
	user_id := e.Get("id").(string)

	var request CreateInput

	if err := e.Bind(&request); err != nil {
		log.Printf("Error on CreatePost request body: %v", err.Error())
		return e.JSON(http.StatusBadRequest, fres.Response.StatusBadRequest(err.Error()))
	}

	if err := h.validate.Struct(request); err != nil {
		log.Printf("Error on CreatePost validation: %v", err.Error())
		return e.JSON(http.StatusBadRequest, fres.Response.StatusBadRequest(err.Error()))
	}

	post, err := h.postsService.CreatePost(domain.Posts{
		UserID:   user_id,
		Content:  request.Content,
		ImageUrl: request.ImageUrl,
	})
	if err != nil {
		log.Printf("Error on CreatePost internal: %v", err.Error())
		return e.JSON(http.StatusInternalServerError, fres.Response.StatusInternalServerError(http.StatusInternalServerError))
	}

	return e.JSON(http.StatusCreated, fres.Response.StatusCreated(post))
}

func (h PostsHandler) GetAllPost(e echo.Context) error {
	var request GetAllPostInput

	if err := e.Bind(&request); err != nil {
		log.Printf("Error on GetAllPost request body: %v", err.Error())
		return e.JSON(http.StatusBadRequest, fres.Response.StatusBadRequest(err.Error()))
	}

	if err := h.validate.Struct(request); err != nil {
		log.Printf("Error on GetAllPost validation: %v", err.Error())
		return e.JSON(http.StatusBadRequest, fres.Response.StatusBadRequest(err.Error()))
	}

	posts, err := h.postsService.GetAllPost(request.Page, request.Limit)
	if err != nil {
		log.Printf("Error on GetAllPost internal: %v", err.Error())
		return e.JSON(http.StatusInternalServerError, fres.Response.StatusInternalServerError(http.StatusInternalServerError))
	}

	return e.JSON(http.StatusOK, fres.Response.StatusOK(posts))
}

func (h PostsHandler) GetPostByID(e echo.Context) error {
	id := e.Param("id")
	post_id, _ := strconv.Atoi(id)
	post, err := h.postsService.GetPostByID(int64(post_id))
	if err != nil {
		if strings.Contains(err.Error(), "found") {
			log.Printf("Error on GetPostByID request: %v", err.Error())
			return e.JSON(http.StatusNotFound, fres.Response.StatusNotFound(err.Error()))
		}
		log.Printf("Error on GetAllPost internal: %v", err.Error())
		return e.JSON(http.StatusInternalServerError, fres.Response.StatusInternalServerError(http.StatusInternalServerError))
	}

	return e.JSON(http.StatusOK, fres.Response.StatusOK(post))
}

func (h PostsHandler) UpdatePost(e echo.Context) error {
	user_id := e.Get("id").(string)

	id := e.Param("id")
	post_id, _ := strconv.Atoi(id)

	var request UpdateInput

	if err := e.Bind(&request); err != nil {
		log.Printf("Error on UpdatePost request body: %v", err.Error())
		return e.JSON(http.StatusBadRequest, fres.Response.StatusBadRequest(err.Error()))
	}

	if err := h.validate.Struct(request); err != nil {
		log.Printf("Error on UpdatePost validation: %v", err.Error())
		return e.JSON(http.StatusBadRequest, fres.Response.StatusBadRequest(err.Error()))
	}

	err := h.postsService.UpdatePost(domain.Posts{
		ID:       int64(post_id),
		UserID:   user_id,
		Content:  request.Content,
		ImageUrl: request.ImageUrl,
	})
	if err != nil {
		if strings.Contains(err.Error(), "no rows") || strings.Contains(err.Error(), "found") {
			log.Printf("Error on GetPostByID request: %v", err.Error())
			return e.JSON(http.StatusNotFound, fres.Response.StatusNotFound(err.Error()))
		}

		log.Printf("Error on UpdatePost internal: %v", err.Error())
		return e.JSON(http.StatusInternalServerError, fres.Response.StatusInternalServerError(http.StatusInternalServerError))
	}

	return e.JSON(http.StatusCreated, fres.Response.StatusCreated("Post updated successfully"))
}

func (h PostsHandler) DeletePost(e echo.Context) error {
	user_id := e.Get("id").(string)

	id := e.Param("id")
	post_id, _ := strconv.Atoi(id)

	err := h.postsService.DeletePost(int64(post_id), user_id)
	if err != nil {
		if strings.Contains(err.Error(), "found") {
			log.Printf("Error on GetPostByID request: %v", err.Error())
			return e.JSON(http.StatusNotFound, fres.Response.StatusNotFound(err.Error()))
		}
		log.Printf("Error on GetAllPost internal: %v", err.Error())
		return e.JSON(http.StatusInternalServerError, fres.Response.StatusInternalServerError(http.StatusInternalServerError))
	}

	return e.JSON(http.StatusOK, fres.Response.StatusOK("Post deleted successfully"))
}
