package posts_repository

import (
	"context"
	"encoding/json"
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

func (r *GormPostRepository) GetAllPost(page, limit int, user_id string) ([]domain.PostsWithUsername, error) {
	var rows []domain.PostFromDB
	err := r.DB.WithContext(r.ctx).
		// Select("posts.id, posts.user_id, users.first_name, users.last_name, users.username, posts.content, posts.image_url, posts.created_at, posts.updated_at, exists(select * from likes where likes.post_id = posts.id and likes.user_id = ?) as is_liked", user_id).
		Select("posts.id, posts.user_id, users.first_name, users.last_name, users.username, posts.content, posts.image_url, posts.created_at, posts.updated_at, exists(select 1 from likes where likes.post_id = posts.id and likes.user_id = ?) as is_liked, (select count(*) from likes where likes.post_id = posts.id) as likes_count, (select count(*) from comments where comments.post_id = posts.id) as comments_count", user_id).
		Joins("JOIN users ON posts.user_id = users.id").
		Order("posts.created_at DESC").Offset((page - 1) * limit).Limit(limit).Scan(&rows).Error
	if err != nil {
		return nil, err
	}

	var posts []domain.PostsWithUsername
	for _, row := range rows {
		var images []string
		if row.ImageUrl != "" {
			_ = json.Unmarshal([]byte(row.ImageUrl), &images)
		}

		posts = append(posts, domain.PostsWithUsername{
			ID:            row.ID,
			UserID:        row.UserID,
			FirstName:     row.Username,
			LastName:      row.LastName,
			Username:      row.Username,
			Content:       row.Content,
			ImageUrl:      images,
			CreatedAt:     row.CreatedAt,
			UpdatedAt:     row.UpdatedAt,
			IsLiked:       row.IsLiked,
			LikesCount:    row.LikesCount,
			CommentsCount: row.CommentsCount,
		})
	}

	return posts, nil
}

func (r *GormPostRepository) GetPostByID(id int64, user_id string) (domain.PostsWithUsername, error) {
	var row domain.PostFromDB
	err := r.DB.WithContext(r.ctx).
		Select("posts.id, posts.user_id, users.first_name, users.last_name, users.username, posts.content, posts.image_url, posts.created_at, posts.updated_at, exists(select 1 from likes where likes.post_id = posts.id and likes.user_id = ?) as is_liked, (select count(*) from likes where likes.post_id = posts.id) as likes_count, (select count(*) from comments where comments.post_id = posts.id) as comments_count", user_id).
		Joins("JOIN users ON posts.user_id = users.id").
		Where("posts.id=?", id).Scan(&row).Error
	if err != nil {
		return domain.PostsWithUsername{}, err
	}

	var imagesUrls []string

	_ = json.Unmarshal([]byte(row.ImageUrl), &imagesUrls)

	return domain.PostsWithUsername{
		ID:            row.ID,
		UserID:        row.UserID,
		FirstName:     row.FirstName,
		LastName:      row.LastName,
		Username:      row.Username,
		Content:       row.Content,
		ImageUrl:      imagesUrls,
		CreatedAt:     row.CreatedAt,
		UpdatedAt:     row.UpdatedAt,
		IsLiked:       row.IsLiked,
		LikesCount:    row.LikesCount,
		CommentsCount: row.CommentsCount,
	}, nil
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
