package usecase

import (
	"booking/internal/entity"
	"context"
	"encoding/json"

	"github.com/nats-io/nats.go"
)

type BookingRepository interface {
	CreateBooking(ctx context.Context, booking *entity.Booking) error
	ListBookings(ctx context.Context, userID string) ([]*entity.Booking, error)
}

type BookingCache interface {
	GetBookings(ctx context.Context, userID string) ([]*entity.Booking, error)
	SetBookings(ctx context.Context, userID string, bookings []*entity.Booking) error
	InvalidateBookings(ctx context.Context, userID string) error
}

type BookingUseCase struct {
	repo  BookingRepository
	cache BookingCache
	nats  *nats.Conn
}

func NewBookingUseCase(r BookingRepository, c BookingCache, nc *nats.Conn) *BookingUseCase {
	return &BookingUseCase{repo: r, cache: c, nats: nc}
}

func (uc *BookingUseCase) CreateBooking(ctx context.Context, booking *entity.Booking) error {
	err := uc.repo.CreateBooking(ctx, booking)
	if err != nil {
		return err
	}
	_ = uc.cache.InvalidateBookings(ctx, booking.UserID)

	payload, _ := json.Marshal(booking)
	_ = uc.nats.Publish("booking.created", payload)
	return nil
}

func (uc *BookingUseCase) ListBookings(ctx context.Context, userID string) ([]*entity.Booking, error) {
	bookings, err := uc.cache.GetBookings(ctx, userID)
	if err != nil {
		return nil, err
	}
	if bookings != nil {
		return bookings, nil
	}

	bookings, err = uc.repo.ListBookings(ctx, userID)
	if err != nil {
		return nil, err
	}
	_ = uc.cache.SetBookings(ctx, userID, bookings)
	return bookings, nil
}
