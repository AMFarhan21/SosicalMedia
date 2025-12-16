package likes_service

import "socialmedia/domain"

type LikesRepo interface {
	Likes(data domain.Likes) (domain.Likes, error)
	UnLikes(id int, user_id string) error
	FindLikesByID(user_id string, post_id, comment_id *int) (domain.Likes, error)
}
