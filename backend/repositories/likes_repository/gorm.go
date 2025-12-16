package likes_repository

import (
	"context"
	"errors"
	"socialmedia/domain"

	"gorm.io/gorm"
)

type GormLikesRepository struct {
	DB  *gorm.DB
	ctx context.Context
}

func NewGormLikesRepository(db *gorm.DB) *GormLikesRepository {
	return &GormLikesRepository{
		DB:  db.Table("likes"),
		ctx: context.Background(),
	}
}

func (r *GormLikesRepository) Likes(data domain.Likes) (domain.Likes, error) {
	err := r.DB.WithContext(r.ctx).Create(&data).Error
	if err != nil {
		return domain.Likes{}, err
	}

	return data, nil
}

func (r *GormLikesRepository) UnLikes(id int, user_id string) error {
	row := r.DB.WithContext(r.ctx).Where("id=?", id).Where("user_id=?", user_id).Delete(domain.Likes{})

	if err := row.Error; err != nil {
		return err
	}

	if row.RowsAffected == 0 {
		return errors.New("no rows affected, likes doesnt exists")
	}

	return nil
}

func (r *GormLikesRepository) FindLikesByID(user_id string, post_id, comment_id *int) (domain.Likes, error) {
	var like domain.Likes
	query := r.DB.WithContext(r.ctx).Where("user_id=?", user_id)
	if post_id != nil && comment_id == nil {
		query = query.Where("post_id=?", post_id)
	} else if comment_id != nil && post_id == nil {
		query = query.Where("comment_id=?", comment_id)
	} else {
		return domain.Likes{}, errors.New("post_id or comment_id must not appear twice in the same row")
	}

	err := query.First(&like).Error
	if err != nil {
		return domain.Likes{}, err
	}

	return like, nil
}
