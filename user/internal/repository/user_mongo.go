package repository

import (
	"context"
	"errors"
	"user/internal/entity"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserMongoRepository struct {
	collection *mongo.Collection
}

func NewUserMongoRepository(c *mongo.Collection) *UserMongoRepository {
	return &UserMongoRepository{collection: c}
}

func (r *UserMongoRepository) GetUserByID(ctx context.Context, userID string) (*entity.User, error) {
	filter := bson.M{"user_id": userID}
	var user entity.User
	err := r.collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (r *UserMongoRepository) GetUserByEmail(ctx context.Context, email string) (*entity.User, error) {
	filter := bson.M{"email": email}
	var user entity.User
	err := r.collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (r *UserMongoRepository) CreateUser(ctx context.Context, user *entity.User) error {
	_, err := r.collection.InsertOne(ctx, user)
	return err
}
