package grpc

import (
	"booking/internal/entity"
	"booking/internal/usecase"
	"booking/proto/gen"
	"context"
	"log"
)

type BookingHandler struct {
	gen.UnimplementedBookingServiceServer
	uc *usecase.BookingUseCase
}

func NewBookingHandler(uc *usecase.BookingUseCase) *BookingHandler {
	return &BookingHandler{uc: uc}
}

func (h *BookingHandler) CreateBooking(ctx context.Context, req *gen.CreateBookingRequest) (*gen.CreateBookingResponse, error) {
	booking := &entity.Booking{
		BookingID: req.GetBookingId(),
		UserID:    req.GetUserId(),
		SpaceID:   req.GetSpaceId(),
		Date:      req.GetDate(),
		Email:     req.GetEmail(),
	}

	if err := h.uc.CreateBooking(ctx, booking); err != nil {
		log.Printf("failed to create booking: %v", err)
		return nil, err
	}

	return &gen.CreateBookingResponse{Status: "created"}, nil
}

func (h *BookingHandler) ListBookings(ctx context.Context, req *gen.ListBookingsRequest) (*gen.ListBookingsResponse, error) {
	bookings, err := h.uc.ListBookings(ctx, req.GetUserId())
	if err != nil {
		log.Printf("failed to list bookings: %v", err)
		return nil, err
	}
	var out []*gen.BookingInfo
	for _, b := range bookings {
		out = append(out, &gen.BookingInfo{
			BookingId: b.BookingID,
			SpaceId:   b.SpaceID,
			Date:      b.Date,
		})
	}
	return &gen.ListBookingsResponse{Bookings: out}, nil
}

func (h *BookingHandler) GetBookingByID(ctx context.Context, req *gen.GetBookingByIDRequest) (*gen.GetBookingByIDResponse, error) {
	booking, err := h.uc.GetBookingByID(ctx, req.GetBookingId())
	if err != nil {
		log.Printf("failed to get booking by ID: %v", err)
		return nil, err
	}

	return &gen.GetBookingByIDResponse{
		BookingId: booking.BookingID,
		UserId:    booking.UserID,
		SpaceId:   booking.SpaceID,
		Date:      booking.Date,
		Email:     booking.Email,
		Status:    booking.Status,
	}, nil
}

func (h *BookingHandler) CancelBooking(ctx context.Context, req *gen.CancelBookingRequest) (*gen.CancelBookingResponse, error) {
	err := h.uc.CancelBooking(ctx, req.GetBookingId())
	if err != nil {
		log.Printf("failed to cancel booking: %v", err)
		return nil, err
	}

	return &gen.CancelBookingResponse{
		Status: "cancelled",
	}, nil
}
