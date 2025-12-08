package posts_service

import "socialmedia/domain"

type PostsRepo interface {
	CreatePost(data domain.Posts) (domain.Posts, error)
	GetAllPost(page, limit int) ([]domain.Posts, error)
	GetPostByID(id int) (domain.Posts, error)
	UpdatePost(data domain.Posts) error
	DeletePost(id int, user_id string) error
}
