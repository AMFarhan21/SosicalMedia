package posts_service

import "socialmedia/domain"

type PostsRepo interface {
	CreatePost(data domain.Posts) (domain.Posts, error)
	GetAllPost(page, limit int) ([]domain.Posts, error)
	GetPostByID(id int64) (domain.Posts, error)
	UpdatePost(data domain.Posts) error
	DeletePost(id int64, user_id string) error
}
