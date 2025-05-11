package grpc

import (
	"booking-service/internal/usecase"
	pb "booking-service/proto"
	"context"
	"time"
)

type BookingHandler struct {
	pb.UnimplementedBookingServiceServer
	usecase *usecase.BookingUseCase
}

func NewBookingHandler(uc *usecase.BookingUseCase) *BookingHandler {
	return &BookingHandler{usecase: uc}
}

func (h *BookingHandler) CreateBooking(ctx context.Context, req *pb.CreateBookingRequest) (*pb.CreateBookingResponse, error) {
	start, _ := time.Parse(time.RFC3339, req.StartTime)
	end, _ := time.Parse(time.RFC3339, req.EndTime)
	err := h.usecase.Create(req.UserId, req.SpaceId, start, end)
	if err != nil {
		return nil, err
	}
	return &pb.CreateBookingResponse{Success: true}, nil
}
