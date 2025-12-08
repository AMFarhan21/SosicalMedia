package posts_service

import (
	"socialmedia/domain"
	"time"
)

type PostsService struct {
	postsRepo PostsRepo
}

type Service interface {
	CreatePost(data domain.Posts) (domain.Posts, error)
	GetAllPost(page, limit int) ([]domain.Posts, error)
	GetPostByID(id int) (domain.Posts, error)
	UpdatePost(data domain.Posts) error
	DeletePost(id int, user_id string) error
}

func NewPostsService(repo PostsRepo) Service {
	return &PostsService{
		postsRepo: repo,
	}
}

func (s PostsService) CreatePost(data domain.Posts) (domain.Posts, error) {
	now := time.Now()
	data.CreatedAt = now
	data.UpdatedAt = now

	return s.postsRepo.CreatePost(data)
}
func (s PostsService) GetAllPost(page, limit int) ([]domain.Posts, error) {
	return s.postsRepo.GetAllPost(page, limit)
}
func (s PostsService) GetPostByID(id int) (domain.Posts, error) {
	return s.postsRepo.GetPostByID(id)
}
func (s PostsService) UpdatePost(data domain.Posts) error {
	data.UpdatedAt = time.Now()
	return s.postsRepo.UpdatePost(data)
}
func (s PostsService) DeletePost(id int, user_id string) error {
	return s.postsRepo.DeletePost(id, user_id)
}
