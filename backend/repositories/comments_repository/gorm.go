package comments_repository

import (
	"context"
	"encoding/json"
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
		DB:  db.Table("comments"),
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

func (r *GormCommentsRepository) GetAllComments(post_id int64, user_id string) ([]domain.CommentsWithUsername, error) {
	var rows []domain.CommentsFromDB
	err := r.DB.WithContext(r.ctx).
		Select("comments.id, comments.user_id, comments.post_id, users.first_name, users.last_name, users.username, comments.content, comments.image_url, comments.created_at, comments.updated_at, exists(select 1 from likes where likes.comment_id = comments.id and likes.user_id = ?) as is_liked, (select count(*) from likes where likes.comment_id = comments.id) as likes_count", user_id).
		Joins("JOIN users on comments.user_id = users.id").
		Order("comments.created_at DESC").Where("post_id=?", post_id).Find(&rows).Error
	if err != nil {
		return nil, err
	}

	var comments []domain.CommentsWithUsername
	for _, row := range rows {
		var images []string
		_ = json.Unmarshal([]byte(row.ImageUrl), &images)

		comments = append(comments, domain.CommentsWithUsername{
			ID:         row.ID,
			UserID:     row.UserID,
			FirstName:  row.FirstName,
			LastName:   row.LastName,
			Username:   row.Username,
			Content:    row.Content,
			ImageUrl:   images,
			CreatedAt:  row.CreatedAt,
			UpdatedAt:  row.UpdatedAt,
			IsLiked:    row.IsLiked,
			LikesCount: row.LikesCount,
		})
	}

	return comments, nil
}

func (r *GormCommentsRepository) GetCommentByID(id int64) (domain.CommentsWithUsername, error) {
	var row domain.CommentsFromDB
	err := r.DB.WithContext(r.ctx).Where("id=?", id).First(&row).Error
	if err != nil {
		return domain.CommentsWithUsername{}, err
	}

	var images []string
	err = json.Unmarshal([]byte(row.ImageUrl), &images)
	if err != nil {
		return domain.CommentsWithUsername{}, err
	}
	return domain.CommentsWithUsername{
		ID:         row.ID,
		UserID:     row.UserID,
		FirstName:  row.FirstName,
		LastName:   row.LastName,
		Username:   row.Username,
		Content:    row.Content,
		ImageUrl:   images,
		CreatedAt:  row.CreatedAt,
		UpdatedAt:  row.UpdatedAt,
		IsLiked:    row.IsLiked,
		LikesCount: row.LikesCount,
	}, nil
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
