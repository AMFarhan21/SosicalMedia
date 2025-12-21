package posts_service

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"os"
	"socialmedia/domain"
	"strconv"
	"time"
)

type PostsService struct {
	postsRepo PostsRepo
	redisRepo RedisRepo
}

type Service interface {
	CreatePost(data domain.Posts, files []*multipart.FileHeader) (domain.Posts, error)
	GetAllPost(page, limit int, user_id string) ([]domain.PostsWithUsername, error)
	GetPostByID(id int64, user_id string) (domain.PostsWithUsername, error)
	UpdatePost(data domain.Posts) error
	DeletePost(id int64, user_id string) error
}

func NewPostsService(repo PostsRepo, redisRepo RedisRepo) Service {
	return &PostsService{
		postsRepo: repo,
		redisRepo: redisRepo,
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

	post, err := s.postsRepo.CreatePost(data)
	if err != nil {
		return domain.Posts{}, err
	}

	err = s.redisRepo.DeleteFeed()
	if err != nil {
		log.Print("Error on redis DeleteFeed()", err.Error())
	}

	return post, nil
}
func (s PostsService) GetAllPost(page, limit int, user_id string) ([]domain.PostsWithUsername, error) {
	resultRedis, err := s.redisRepo.GetAllPost(page, limit, user_id)
	if err == nil {
		fmt.Println("-----------------------------------------")
		fmt.Println("THIS IS FROM REDIS")
		return resultRedis, nil
	}

	fmt.Println("-----------------------------------------")
	fmt.Println("CACHE MISSED -> QUERY DB")

	resultDB, err := s.postsRepo.GetAllPost(page, limit, user_id)
	if err != nil {
		return nil, err
	}
	fmt.Println("-----------------------------------------")
	fmt.Println("THIS IS FROM DB")

	fmt.Println("-----------------------------------------")
	s.redisRepo.CacheGetAllPost(page, limit, user_id, resultDB)
	fmt.Println("CACHED")

	return resultDB, nil
}
func (s PostsService) GetPostByID(id int64, user_id string) (domain.PostsWithUsername, error) {
	return s.postsRepo.GetPostByID(id, user_id)
}
func (s PostsService) UpdatePost(data domain.Posts) error {
	data.UpdatedAt = time.Now()
	return s.postsRepo.UpdatePost(data)
}
func (s PostsService) DeletePost(id int64, user_id string) error {
	err := s.redisRepo.DeleteFeed()
	if err != nil {
		log.Print("Error on redis DeleteFeed()", err.Error())
	}

	return s.postsRepo.DeletePost(id, user_id)
}
