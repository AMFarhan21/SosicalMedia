package posts_repository

import (
	"context"
	"socialmedia/domain"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type MongoPostsRepository struct {
	DB  *mongo.Collection
	ctx context.Context
}

func NewMongoPostsRepository(client *mongo.Client) *MongoPostsRepository {
	return &MongoPostsRepository{
		DB:  client.Database("socialmedia").Collection("posts"),
		ctx: context.Background(),
	}
}

func (r *MongoPostsRepository) CreatePost(data domain.Posts) (domain.Posts, error) {
	data.ID = int(time.Now().UnixNano())
	_, err := r.DB.InsertOne(r.ctx, data)
	if err != nil {
		return domain.Posts{}, err
	}

	return data, nil
}

func (r *MongoPostsRepository) GetAllPost(page, limit int) ([]domain.Posts, error) {
	cursor, err := r.DB.Find(r.ctx, bson.M{}, options.Find().SetSkip(int64((page-1)*limit)).SetLimit(int64(limit)))
	if err != nil {
		return nil, err
	}

	var users []domain.Posts
}
