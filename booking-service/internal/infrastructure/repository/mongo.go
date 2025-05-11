package repository

import (
	"booking-service/internal/domain"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type bookingRepo struct {
	col *mongo.Collection
}

func NewBookingRepo(col *mongo.Collection) domain.BookingRepository {
	return &bookingRepo{col}
}

func (r *bookingRepo) Create(b *domain.Booking) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := r.col.InsertOne(ctx, b)
	return err
}

func (r *bookingRepo) GetByID(id string) (*domain.Booking, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	var b domain.Booking
	err := r.col.FindOne(ctx, bson.M{"id": id}).Decode(&b)
	return &b, err
}

func (r *bookingRepo) ListByUser(userID string) ([]*domain.Booking, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	cursor, err := r.col.Find(ctx, bson.M{"user_id": userID})
	if err != nil {
		return nil, err
	}

	var results []*domain.Booking
	if err := cursor.All(ctx, &results); err != nil {
		return nil, err
	}
	return results, nil
}
