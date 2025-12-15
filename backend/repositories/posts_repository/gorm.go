package posts_repository

import (
	"context"
	"errors"
	"socialmedia/domain"

	"gorm.io/gorm"
)

type GormPostRepository struct {
	DB  *gorm.DB
	ctx context.Context
}

func NewGormPostRepository(db *gorm.DB) *GormPostRepository {
	return &GormPostRepository{
		DB:  db.Table("posts"),
		ctx: context.Background(),
	}
}

func (r *GormPostRepository) CreatePost(data domain.Posts) (domain.Posts, error) {
	err := r.DB.WithContext(r.ctx).Create(&data).Error
	if err != nil {
		return domain.Posts{}, err
	}

	return data, nil
}

func (r *GormPostRepository) GetAllPost(page, limit int) ([]domain.PostsWithUsername, error) {
	var posts []domain.PostsWithUsername
	err := r.DB.WithContext(r.ctx).
		Select("posts.id, posts.user_id, users.username, posts.content, posts.image_url, posts.created_at, posts.updated_at").
		Joins("JOIN users ON posts.user_id = users.id").
		Order("posts.created_at DESC").Offset((page - 1) * limit).Limit(limit).Scan(&posts).Error
	if err != nil {
		return nil, err
	}

	return posts, nil
}

func (r *GormPostRepository) GetPostByID(id int64) (domain.PostsWithUsername, error) {
	var post domain.PostsWithUsername
	err := r.DB.WithContext(r.ctx).
		Select("posts.id, posts.user_id, users.username, posts.content, posts.image_url, posts.created_at, posts.updated_at").
		Joins("JOIN users ON posts.user_id = users.id").
		Where("posts.id=?", id).Scan(&post).Error
	if err != nil {
		return domain.PostsWithUsername{}, err
	}

	return post, nil
}

func (r *GormPostRepository) UpdatePost(data domain.Posts) error {
	row := r.DB.WithContext(r.ctx).Where("id=?", data.ID).Where("user_id=?", data.UserID).Updates(data)
	if err := row.Error; err != nil {
		return err
	}

	if row.RowsAffected == 0 {
		return errors.New("no rows affected, user doesnt exists")
	}

	return nil
}

func (r *GormPostRepository) DeletePost(id int64, user_id string) error {
	row := r.DB.WithContext(r.ctx).Where("id=?", id).Where("user_id=?", user_id).Delete(domain.Posts{})
	if err := row.Error; err != nil {
		return err
	}

	if row.RowsAffected == 0 {
		return errors.New("no rows affected, user doesnt exists")
	}

	return nil
}
