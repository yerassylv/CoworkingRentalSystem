package grpc

import (
	"context"
	"log"

	"space/internal/usecase"
	pb "space/proto"
)

type SpaceHandler struct {
	pb.UnimplementedSpaceServiceServer
	uc *usecase.SpaceUseCase
}

func NewSpaceHandler(uc *usecase.SpaceUseCase) *SpaceHandler {
	return &SpaceHandler{uc: uc}
}

func (h *SpaceHandler) GetSpaceByID(ctx context.Context, req *pb.GetSpaceByIDRequest) (*pb.GetSpaceByIDResponse, error) {
	space, err := h.uc.GetSpaceByID(ctx, req.GetSpaceId())
	if err != nil {
		log.Printf("failed to get space by id: %v", err)
		return nil, err
	}
	if space == nil {
		return nil, nil
	}
	return &pb.GetSpaceByIDResponse{
		SpaceId:   space.SpaceID,
		Name:      space.Name,
		Location:  space.Location,
		Capacity:  space.Capacity,
		Available: space.Available,
	}, nil
}
