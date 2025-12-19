package posts_service

import (
	"encoding/json"
	"io"
	"mime/multipart"
	"os"
	"socialmedia/domain"
	"strconv"
	"time"
)

type PostsService struct {
	postsRepo PostsRepo
}

type Service interface {
	CreatePost(data domain.Posts, files []*multipart.FileHeader) (domain.Posts, error)
	GetAllPost(page, limit int, user_id string) ([]domain.PostsWithUsername, error)
	GetPostByID(id int64, user_id string) (domain.PostsWithUsername, error)
	UpdatePost(data domain.Posts) error
	DeletePost(id int64, user_id string) error
}

func NewPostsService(repo PostsRepo) Service {
	return &PostsService{
		postsRepo: repo,
	}
}

func (s PostsService) CreatePost(data domain.Posts, files []*multipart.FileHeader) (domain.Posts, error) {
	now := time.Now()
	data.CreatedAt = now
	data.UpdatedAt = now

	var imageUrls []string
	for _, file := range files {
		src, err := file.Open()
		if err != nil {
			return domain.Posts{}, err
		}

		defer src.Close()

		fileName := strconv.FormatInt(time.Now().Unix(), 10) + "_" + file.Filename
		dstPath := "uploads/" + fileName

		dst, err := os.Create(dstPath)
		if err != nil {
			return domain.Posts{}, err
		}

		_, err = io.Copy(dst, src)
		if err != nil {
			return domain.Posts{}, err
		}

		imageUrls = append(imageUrls, dstPath)
	}

	imagesJson, _ := json.Marshal(imageUrls)

	data.ImageUrl = string(imagesJson)

	return s.postsRepo.CreatePost(data)
}
func (s PostsService) GetAllPost(page, limit int, user_id string) ([]domain.PostsWithUsername, error) {
	return s.postsRepo.GetAllPost(page, limit, user_id)
}
func (s PostsService) GetPostByID(id int64, user_id string) (domain.PostsWithUsername, error) {
	return s.postsRepo.GetPostByID(id, user_id)
}
func (s PostsService) UpdatePost(data domain.Posts) error {
	data.UpdatedAt = time.Now()
	return s.postsRepo.UpdatePost(data)
}
func (s PostsService) DeletePost(id int64, user_id string) error {
	return s.postsRepo.DeletePost(id, user_id)
}
