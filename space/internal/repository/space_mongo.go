package repository

import (
	"context"
	"errors"
	"space/internal/entity"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func (r *SpaceMongoRepository) CreateSpace(ctx context.Context, space *entity.Space) error {
	_, err := r.collection.InsertOne(ctx, space)
	return err
}

func (r *SpaceMongoRepository) UpdateSpace(ctx context.Context, space *entity.Space) error {
	filter := bson.M{"space_id": space.SpaceID}
	update := bson.M{
		"$set": bson.M{
			"name":      space.Name,
			"location":  space.Location,
			"capacity":  space.Capacity,
			"available": space.Available,
		},
	}
	opts := options.Update().SetUpsert(false)
	_, err := r.collection.UpdateOne(ctx, filter, update, opts)
	return err
}

func (r *SpaceMongoRepository) DeleteSpace(ctx context.Context, spaceID string) error {
	filter := bson.M{"space_id": spaceID}
	_, err := r.collection.DeleteOne(ctx, filter)
	return err
}

func (r *SpaceMongoRepository) ListSpaces(ctx context.Context) ([]*entity.Space, error) {
	cur, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	var spaces []*entity.Space
	for cur.Next(ctx) {
		var space entity.Space
		if err := cur.Decode(&space); err != nil {
			return nil, err
		}
		spaces = append(spaces, &space)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}
	return spaces, nil
}
