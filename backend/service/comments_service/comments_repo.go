package comments_service

import "socialmedia/domain"

type CommentsRepo interface {
	CreateComment(data domain.Comments) (domain.Comments, error)
	GetAllComments(post_id int64, user_id string) ([]domain.CommentsWithUsername, error)
	GetCommentByID(id int64) (domain.CommentsWithUsername, error)
	UpdateComment(data domain.Comments) error
	DeleteComment(id int64, user_id string) error
}
