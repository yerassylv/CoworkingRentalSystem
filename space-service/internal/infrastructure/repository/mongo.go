package repository

import (
	"context"
	"space-service/internal/domain"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type mongoRepo struct {
	col *mongo.Collection
}

func NewMongoRepo(col *mongo.Collection) domain.SpaceRepository {
	return &mongoRepo{col: col}
}

func (r *mongoRepo) DecreaseAvailability(spaceID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	filter := bson.M{"_id": spaceID}
	update := bson.M{"$inc": bson.M{"available": -1}}
	_, err := r.col.UpdateOne(ctx, filter, update)
	return err
}
