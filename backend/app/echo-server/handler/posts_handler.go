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

	CreatePostInput struct {
		Content string `form:"content"`
	}

	UpdatePostInput struct {
		Content string `json:"content" bson:"content"`
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

	var request CreatePostInput

	if err := e.Bind(&request); err != nil {
		log.Printf("Error on CreatePost request body: %v", err.Error())
		return e.JSON(http.StatusBadRequest, fres.Response.StatusBadRequest(err.Error()))
	}

	if err := h.validate.Struct(request); err != nil {
		log.Printf("Error on CreatePost validation: %v", err.Error())
		return e.JSON(http.StatusBadRequest, fres.Response.StatusBadRequest(err.Error()))
	}

	form, err := e.MultipartForm()
	if err != nil {
		log.Printf("Error on multipartform on CreatePost: %v", err.Error())
		return e.JSON(http.StatusBadRequest, fres.Response.StatusBadRequest(err.Error()))
	}

	files := form.File["images"]

	post, err := h.postsService.CreatePost(domain.Posts{
		UserID:  user_id,
		Content: request.Content,
	}, files)
	if err != nil {
		log.Printf("Error on CreatePost internal: %v", err.Error())
		return e.JSON(http.StatusInternalServerError, fres.Response.StatusInternalServerError(http.StatusInternalServerError))
	}

	return e.JSON(http.StatusCreated, fres.Response.StatusCreated(post))
}

func (h PostsHandler) GetAllPost(e echo.Context) error {
	pReq := e.QueryParam("page")
	lReq := e.QueryParam("limit")

	page, _ := strconv.Atoi(pReq)
	limit, _ := strconv.Atoi(lReq)

	if page < 1 {
		page = 1
	}

	if limit < 1 {
		limit = 10
	}

	user_id := e.Get("id").(string)

	posts, err := h.postsService.GetAllPost(page, limit, user_id)
	if err != nil {
		log.Printf("Error on GetAllPost internal: %v", err.Error())
		return e.JSON(http.StatusInternalServerError, fres.Response.StatusInternalServerError(http.StatusInternalServerError))
	}

	return e.JSON(http.StatusOK, fres.Response.StatusOK(posts))
}

func (h PostsHandler) GetPostByID(e echo.Context) error {
	id := e.Param("id")
	post_id, _ := strconv.Atoi(id)
	user_id := e.Get("id").(string)
	post, err := h.postsService.GetPostByID(int64(post_id), user_id)
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

	var request UpdatePostInput

	if err := e.Bind(&request); err != nil {
		log.Printf("Error on UpdatePost request body: %v", err.Error())
		return e.JSON(http.StatusBadRequest, fres.Response.StatusBadRequest(err.Error()))
	}

	if err := h.validate.Struct(request); err != nil {
		log.Printf("Error on UpdatePost validation: %v", err.Error())
		return e.JSON(http.StatusBadRequest, fres.Response.StatusBadRequest(err.Error()))
	}

	err := h.postsService.UpdatePost(domain.Posts{
		ID:      int64(post_id),
		UserID:  user_id,
		Content: request.Content,
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
