package usecase

import (
	"booking/internal/entity"
	"booking/internal/infra/email"
	"context"
	"encoding/json"
	"fmt"

	"github.com/nats-io/nats.go"
)

type BookingRepository interface {
	CreateBooking(ctx context.Context, booking *entity.Booking) error
	ListBookings(ctx context.Context, userID string) ([]*entity.Booking, error)
	GetBookingByID(ctx context.Context, bookingID string) (*entity.Booking, error)
	CancelBooking(ctx context.Context, bookingID string) error // ✅ добавлено
}

type BookingCache interface {
	GetBookings(ctx context.Context, userID string) ([]*entity.Booking, error)
	SetBookings(ctx context.Context, userID string, bookings []*entity.Booking) error
	InvalidateBookings(ctx context.Context, userID string) error
}

type BookingUseCase struct {
	repo   BookingRepository
	cache  BookingCache
	nats   *nats.Conn
	mailer email.Sender
}

func NewBookingUseCase(r BookingRepository, c BookingCache, nc *nats.Conn, m email.Sender) *BookingUseCase {
	return &BookingUseCase{
		repo:   r,
		cache:  c,
		nats:   nc,
		mailer: m,
	}
}

func (uc *BookingUseCase) CreateBooking(ctx context.Context, booking *entity.Booking) error {
	err := uc.repo.CreateBooking(ctx, booking)
	if err != nil {
		return err
	}

	_ = uc.cache.InvalidateBookings(ctx, booking.UserID)

	payload, _ := json.Marshal(booking)
	_ = uc.nats.Publish("booking.created", payload)

	_ = uc.mailer.Send(
		booking.Email,
		"Booking Confirmed",
		fmt.Sprintf("Hello! Your booking with ID %s for date %s has been confirmed.", booking.BookingID, booking.Date),
	)

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

func (uc *BookingUseCase) GetBookingByID(ctx context.Context, bookingID string) (*entity.Booking, error) {
	return uc.repo.GetBookingByID(ctx, bookingID)
}

func (uc *BookingUseCase) CancelBooking(ctx context.Context, bookingID string) error {
	err := uc.repo.CancelBooking(ctx, bookingID)
	if err != nil {
		return fmt.Errorf("failed to cancel booking: %w", err)
	}

	_ = uc.cache.InvalidateBookings(ctx, "") // можно указать userID, если знаешь его

	return nil
}
