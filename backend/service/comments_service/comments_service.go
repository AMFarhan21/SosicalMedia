package comments_service

import (
	"encoding/json"
	"io"
	"mime/multipart"
	"os"
	"socialmedia/domain"
	"strconv"
	"time"
)

type CommentsService struct {
	commentsRepo CommentsRepo
}

type Service interface {
	CreateComment(data domain.Comments, files []*multipart.FileHeader) (domain.Comments, error)
	GetAllComments(post_id int64, user_id string) ([]domain.CommentsWithUsername, error)
	GetCommentByID(id int64) (domain.CommentsWithUsername, error)
	UpdateComment(data domain.Comments) error
	DeleteComment(id int64, user_id string) error
}

func NewCommentsService(repo CommentsRepo) Service {
	return &CommentsService{
		commentsRepo: repo,
	}
}

func (s CommentsService) CreateComment(data domain.Comments, files []*multipart.FileHeader) (domain.Comments, error) {
	now := time.Now()
	data.CreatedAt = now
	data.UpdatedAt = now

	var imageUrls []string
	for _, file := range files {
		src, err := file.Open()
		if err != nil {
			return domain.Comments{}, err
		}

		defer src.Close()

		fileName := strconv.FormatInt(time.Now().Unix(), 10) + "_" + file.Filename
		dstPath := "uploads/" + fileName

		dst, err := os.Create(dstPath)
		if err != nil {
			return domain.Comments{}, err
		}
		_, err = io.Copy(dst, src)
		if err != nil {
			return domain.Comments{}, err
		}

		imageUrls = append(imageUrls, dstPath)
	}

	imagesJson, _ := json.Marshal(imageUrls)

	data.ImageUrl = string(imagesJson)

	return s.commentsRepo.CreateComment(data)
}
func (s CommentsService) GetAllComments(post_id int64, user_id string) ([]domain.CommentsWithUsername, error) {
	return s.commentsRepo.GetAllComments(post_id, user_id)
}
func (s CommentsService) GetCommentByID(id int64) (domain.CommentsWithUsername, error) {
	return s.commentsRepo.GetCommentByID(id)
}
func (s CommentsService) UpdateComment(data domain.Comments) error {
	data.UpdatedAt = time.Now()
	return s.commentsRepo.UpdateComment(data)
}
func (s CommentsService) DeleteComment(id int64, user_id string) error {
	return s.commentsRepo.DeleteComment(id, user_id)
}
