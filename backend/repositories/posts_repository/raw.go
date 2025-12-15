package posts_repository

import (
	"database/sql"
	"errors"
	"socialmedia/domain"
)

type RawPostsRepository struct {
	DB *sql.DB
}

func NewRawPostsRepository(db *sql.DB) *RawPostsRepository {
	return &RawPostsRepository{
		DB: db,
	}
}

func (r *RawPostsRepository) CreatePost(data domain.Posts) (domain.Posts, error) {
	query := `
		insert into posts (user_id, content, image_url, created_at, updated_at) values
		(?, ?, ?, ?, ?) returning *;
	`

	var post domain.Posts

	err := r.DB.QueryRow(query, data.UserID, data.Content, data.ImageUrl, data.CreatedAt, data.UpdatedAt).Scan(
		post.ID,
		post.UserID,
		post.Content,
		post.ImageUrl,
		post.CreatedAt,
		post.UpdatedAt,
	)
	if err != nil {
		return domain.Posts{}, err
	}

	return post, nil
}

func (r *RawPostsRepository) GetAllPost(page, limit int) ([]domain.Posts, error) {
	query := `
		select * from posts order created_at desc offset (? - 1) * ? limit ?;
	`

	row, err := r.DB.Query(query, page, limit, limit)
	if err != nil {
		return nil, err
	}
	var posts []domain.Posts
	for row.Next() {
		var post domain.Posts
		err := row.Scan(post.ID, post.UserID, post.Content, post.ImageUrl, post.CreatedAt, post.UpdatedAt)
		if err != nil {
			return nil, err
		}

		posts = append(posts, post)
	}

	return posts, nil
}

func (r *RawPostsRepository) GetPostByID(id int64) (domain.Posts, error) {
	query := `
		select * from posts where id = ?;
	`
	var post domain.Posts
	err := r.DB.QueryRow(query, id).Scan(
		post.ID,
		post.UserID,
		post.Content,
		post.ImageUrl,
		post.CreatedAt,
		post.UpdatedAt,
	)
	if err != nil {
		return domain.Posts{}, err
	}

	return post, nil
}

func (r *RawPostsRepository) UpdatePost(data domain.Posts) error {
	query := `
		update posts set user_id=?, content=?, image_url=?, created_at=?, updated_at=? where id = ?;
	`

	row, err := r.DB.Exec(query, data.UserID, data.Content, data.ImageUrl, data.CreatedAt, data.UpdatedAt, data.ID)
	if err != nil {
		return err
	}

	rowsAffected, err := row.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("no rows affected, post doesnt exists")
	}

	return nil
}

func (r *RawPostsRepository) DeletePost(id int64) error {
	query := `delete users where id = ?`
	row, err := r.DB.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := row.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("no rows affected, post doesnt exists")
	}

	return nil
}
