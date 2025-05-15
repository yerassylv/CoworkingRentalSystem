package repository

import (
	"context"
	"errors"
	"space/internal/entity"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type SpaceMongoRepository struct {
	collection *mongo.Collection
}

func NewSpaceMongoRepository(c *mongo.Collection) *SpaceMongoRepository {
	return &SpaceMongoRepository{collection: c}
}

func (r *SpaceMongoRepository) GetSpaceByID(ctx context.Context, spaceID string) (*entity.Space, error) {
	filter := bson.M{"space_id": spaceID}
	var space entity.Space
	err := r.collection.FindOne(ctx, filter).Decode(&space)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}
		return nil, err
	}
	return &space, nil
}
