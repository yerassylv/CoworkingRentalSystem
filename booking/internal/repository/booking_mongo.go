package repository

import (
	"booking/internal/entity"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type BookingMongoRepository struct {
	collection *mongo.Collection
}

func NewBookingMongoRepository(c *mongo.Collection) *BookingMongoRepository {
	return &BookingMongoRepository{collection: c}
}

func (r *BookingMongoRepository) CreateBooking(ctx context.Context, booking *entity.Booking) error {
	_, err := r.collection.InsertOne(ctx, booking)
	return err
}

func (r *BookingMongoRepository) ListBookings(ctx context.Context, userID string) ([]*entity.Booking, error) {
	filter := bson.M{"user_id": userID}
	cursor, err := r.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var bookings []*entity.Booking
	for cursor.Next(ctx) {
		var b entity.Booking
		if err := cursor.Decode(&b); err != nil {
			return nil, err
		}
		bookings = append(bookings, &b)
	}
	return bookings, nil
}
