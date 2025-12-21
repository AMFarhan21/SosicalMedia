package posts_service

import "socialmedia/domain"

type PostsRepo interface {
	CreatePost(data domain.Posts) (domain.Posts, error)
	GetAllPost(page, limit int, user_id string) ([]domain.PostsWithUsername, error)
	GetPostByID(id int64, user_id string) (domain.PostsWithUsername, error)
	UpdatePost(data domain.Posts) error
	DeletePost(id int64, user_id string) error
}

type RedisRepo interface {
	GetAllPost(page, limit int, user_id string) ([]domain.PostsWithUsername, error)
	CacheGetAllPost(page, limit int, user_id string, data []domain.PostsWithUsername)
	DeleteFeed() error
}
