package users_repository

import (
	"context"
	"errors"
	"socialmedia/domain"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type MongoUsersRepository struct {
	DB  *mongo.Collection
	ctx context.Context
}

func NewMongoUsersRepository(client *mongo.Client) *MongoUsersRepository {
	return &MongoUsersRepository{
		DB:  client.Database("socialmedia").Collection("users"),
		ctx: context.Background(),
	}
}

func (r *MongoUsersRepository) CreateUser(data domain.Users) (domain.Users, error) {
	_, err := r.DB.InsertOne(r.ctx, data)
	if err != nil {
		return domain.Users{}, err
	}

	return data, nil
}

func (r *MongoUsersRepository) GetAllUsers(page, limit int) ([]domain.Users, error) {
	var users []domain.Users

	opts := options.Find().
		SetSkip(int64((page - 1) * limit)).
		SetLimit(int64(limit))

	cursor, err := r.DB.Find(r.ctx, bson.M{}, opts)
	if err != nil {
		return nil, err
	}

	err = cursor.All(r.ctx, &users)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (r *MongoUsersRepository) GetUserByID(id string) (domain.Users, error) {
	var user domain.Users

	err := r.DB.FindOne(r.ctx, bson.M{"id": id}).Decode(&user)
	if err != nil {
		return domain.Users{}, err
	}

	return user, nil
}

func (r *MongoUsersRepository) UpdateUser(data domain.Users) error {
	cursor, err := r.DB.UpdateOne(r.ctx, bson.M{"id": data.ID}, bson.M{"$set": data})
	if err != nil {
		return err
	}

	if cursor.ModifiedCount == 0 {
		return errors.New("no rows affected, user doesnt exists")
	}

	return nil
}

func (r *MongoUsersRepository) DeleteUser(id string) error {
	cursor, err := r.DB.DeleteOne(r.ctx, bson.M{"id": id})
	if err != nil {
		return err
	}

	if cursor.DeletedCount == 0 {
		return errors.New("no rows affected, user doesnt exists")
	}

	return nil
}

func (r *MongoUsersRepository) FindByEmail(email string) (domain.Users, error) {
	var user domain.Users
	err := r.DB.FindOne(r.ctx, bson.M{"email": email}).Decode(&user)
	if err != nil {
		return domain.Users{}, err
	}

	return user, nil
}
