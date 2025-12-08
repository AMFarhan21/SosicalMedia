package posts_repository

import (
	"context"
	"errors"
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
	data.ID = int64(time.Now().UnixNano())
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

	var posts []domain.Posts
	err = cursor.All(r.ctx, &posts)
	if err != nil {
		return nil, err
	}

	return posts, nil
}

func (r *MongoPostsRepository) GetPostByID(id int64) (domain.Posts, error) {
	var post domain.Posts
	err := r.DB.FindOne(r.ctx, bson.M{"id": id}).Decode(&post)
	if err != nil {
		return domain.Posts{}, err
	}

	return post, nil
}

func (r *MongoPostsRepository) UpdatePost(data domain.Posts) error {
	cursor, err := r.DB.UpdateOne(r.ctx, bson.M{"id": data.ID}, bson.M{"$set": data})
	if err != nil {
		return err
	}

	if cursor.ModifiedCount == 0 {
		return errors.New("no rows affected, post doesnt exists")
	}

	return nil
}

func (r *MongoPostsRepository) DeletePost(id int64) error {
	cursor, err := r.DB.DeleteOne(r.ctx, bson.M{"id": id})
	if err != nil {
		return err
	}

	if cursor.DeletedCount == 0 {
		return errors.New("no rows affected, post doesnt exists")
	}

	return nil
}
