package grpc

import (
	"booking/internal/entity"
	"booking/internal/usecase"
	"context"
	"log"

	pb "booking/proto/gen"
)

type BookingHandler struct {
	pb.UnimplementedBookingServiceServer
	uc *usecase.BookingUseCase
}

func NewBookingHandler(uc *usecase.BookingUseCase) *BookingHandler {
	return &BookingHandler{uc: uc}
}

func (h *BookingHandler) CreateBooking(ctx context.Context, req *pb.CreateBookingRequest) (*pb.CreateBookingResponse, error) {
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

	return &pb.CreateBookingResponse{Status: "created"}, nil
}

func (h *BookingHandler) ListBookings(ctx context.Context, req *pb.ListBookingsRequest) (*pb.ListBookingsResponse, error) {
	bookings, err := h.uc.ListBookings(ctx, req.GetUserId())
	if err != nil {
		log.Printf("failed to list bookings: %v", err)
		return nil, err
	}
	var out []*pb.BookingInfo
	for _, b := range bookings {
		out = append(out, &pb.BookingInfo{
			BookingId: b.BookingID,
			SpaceId:   b.SpaceID,
			Date:      b.Date,
		})
	}
	return &pb.ListBookingsResponse{Bookings: out}, nil
}
