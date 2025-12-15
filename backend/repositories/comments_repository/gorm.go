package comments_repository

import (
	"context"
	"errors"
	"socialmedia/domain"

	"gorm.io/gorm"
)

type GormCommentsRepository struct {
	DB  *gorm.DB
	ctx context.Context
}

func NewGormCommentsRepository(db *gorm.DB) *GormCommentsRepository {
	return &GormCommentsRepository{
		DB:  db,
		ctx: context.Background(),
	}
}

func (r *GormCommentsRepository) CreateComment(data domain.Comments) (domain.Comments, error) {
	err := r.DB.WithContext(r.ctx).Create(&data).Error
	if err != nil {
		return domain.Comments{}, err
	}

	return data, nil
}

func (r *GormCommentsRepository) GetAllComments(post_id int64) ([]domain.Comments, error) {
	var comments []domain.Comments
	err := r.DB.WithContext(r.ctx).Order("created_at DESC").Where("post_id=?", post_id).Find(&comments).Error
	if err != nil {
		return nil, err
	}

	return comments, nil
}

func (r *GormCommentsRepository) GetCommentByID(id int64) (domain.Comments, error) {
	var comment domain.Comments
	err := r.DB.WithContext(r.ctx).Where("id=?", id).First(&comment).Error
	if err != nil {
		return domain.Comments{}, err
	}

	return comment, nil
}

func (r *GormCommentsRepository) UpdateComment(data domain.Comments) error {
	row := r.DB.WithContext(r.ctx).Where("id=?", data.ID).Where("user_id=?", data.UserID).Updates(data)
	if err := row.Error; err != nil {
		return err
	}

	if row.RowsAffected == 0 {
		return errors.New("no rows affected, user doesnt exists")
	}

	return nil
}

func (r *GormCommentsRepository) DeleteComment(id int64, user_id string) error {
	row := r.DB.WithContext(r.ctx).Where("id=?", id).Where("user_id=?", user_id).Delete(domain.Comments{})
	if err := row.Error; err != nil {
		return err
	}

	if row.RowsAffected == 0 {
		return errors.New("no rows affected, user doesnt exists")
	}

	return nil
}
