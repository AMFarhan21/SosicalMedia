package likes_service

import (
	"socialmedia/domain"
	"strings"
)

type LikesService struct {
	likesRepo LikesRepo
}

type Service interface {
	Likes(data domain.Likes) (domain.Likes, error)
}

func NewLikesService(likesRepo LikesRepo) Service {
	return &LikesService{
		likesRepo: likesRepo,
	}
}

func (s *LikesService) Likes(data domain.Likes) (domain.Likes, error) {
	like, err := s.likesRepo.FindLikesByID(data.UserId, data.PostId, data.CommentId)
	if err != nil && !strings.Contains(err.Error(), "record not found") {
		return domain.Likes{}, err
	}

	if like.ID != 0 {
		return domain.Likes{}, s.likesRepo.UnLikes(like.ID, like.UserId)
	}

	return s.likesRepo.Likes(data)
}
