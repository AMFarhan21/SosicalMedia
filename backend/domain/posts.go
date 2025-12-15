package domain

import "time"

type Posts struct {
	ID        int64     `json:"id" bson:"id" gorm:"primaryKey;autoIncrement"`
	UserID    string    `json:"user_id" bson:"user_id"`
	Content   string    `json:"content" bson:"content"`
	ImageUrl  string    `json:"image_url" bson:"image_url"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
}

type PostsWithUsername struct {
	ID        int64     `json:"id" bson:"id" gorm:"primaryKey;autoIncrement"`
	UserID    string    `json:"user_id" bson:"user_id"`
	Username  string    `json:"username" bson:"username"`
	Content   string    `json:"content" bson:"content"`
	ImageUrl  string    `json:"image_url" bson:"image_url"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
}
