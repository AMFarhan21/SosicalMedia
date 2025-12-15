package comments_service

import (
	"socialmedia/domain"
	"time"
)

type CommentsService struct {
	commentsRepo CommentsRepo
}

type Service interface {
	CreateComment(data domain.Comments) (domain.Comments, error)
	GetAllComments(post_id int64) ([]domain.Comments, error)
	GetCommentByID(id int64) (domain.Comments, error)
	UpdateComment(data domain.Comments) error
	DeleteComment(id int64, user_id string) error
}

func NewCommentsService(repo CommentsRepo) Service {
	return &CommentsService{
		commentsRepo: repo,
	}
}

func (s CommentsService) CreateComment(data domain.Comments) (domain.Comments, error) {
	now := time.Now()
	data.CreatedAt = now
	data.UpdatedAt = now

	return s.commentsRepo.CreateComment(data)
}
func (s CommentsService) GetAllComments(post_id int64) ([]domain.Comments, error) {
	return s.commentsRepo.GetAllComments(post_id)
}
func (s CommentsService) GetCommentByID(id int64) (domain.Comments, error) {
	return s.commentsRepo.GetCommentByID(id)
}
func (s CommentsService) UpdateComment(data domain.Comments) error {
	data.UpdatedAt = time.Now()
	return s.commentsRepo.UpdateComment(data)
}
func (s CommentsService) DeleteComment(id int64, user_id string) error {
	return s.commentsRepo.DeleteComment(id, user_id)
}
