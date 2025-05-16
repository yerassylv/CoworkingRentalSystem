package grpc

import (
	"context"
	"log"
	"space/internal/entity"
	"space/internal/usecase"
	pb "space/proto/gen"
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

func (h *SpaceHandler) CreateSpace(ctx context.Context, req *pb.CreateSpaceRequest) (*pb.CreateSpaceResponse, error) {
	space, err := h.uc.CreateSpace(ctx, req.GetName(), req.GetLocation(), req.GetCapacity(), req.GetAvailable())
	if err != nil {
		log.Printf("failed to create space: %v", err)
		return nil, err
	}
	return &pb.CreateSpaceResponse{
		SpaceId: space.SpaceID,
	}, nil
}

func (h *SpaceHandler) UpdateSpace(ctx context.Context, req *pb.UpdateSpaceRequest) (*pb.UpdateSpaceResponse, error) {
	space := &entity.Space{
		SpaceID:   req.GetSpaceId(),
		Name:      req.GetName(),
		Location:  req.GetLocation(),
		Capacity:  req.GetCapacity(),
		Available: req.GetAvailable(),
	}

	err := h.uc.UpdateSpace(ctx, space)
	if err != nil {
		log.Printf("failed to update space: %v", err)
		return nil, err
	}
	return &pb.UpdateSpaceResponse{}, nil
}

func (h *SpaceHandler) DeleteSpace(ctx context.Context, req *pb.DeleteSpaceRequest) (*pb.DeleteSpaceResponse, error) {
	err := h.uc.DeleteSpace(ctx, req.GetSpaceId())
	if err != nil {
		log.Printf("failed to delete space: %v", err)
		return nil, err
	}
	return &pb.DeleteSpaceResponse{}, nil
}

func (h *SpaceHandler) ListSpaces(ctx context.Context, req *pb.ListSpacesRequest) (*pb.ListSpacesResponse, error) {
	spaces, err := h.uc.ListSpaces(ctx)
	if err != nil {
		log.Printf("failed to list spaces: %v", err)
		return nil, err
	}

	var pbSpaces []*pb.Space
	for _, space := range spaces {
		pbSpaces = append(pbSpaces, &pb.Space{
			SpaceId:   space.SpaceID,
			Name:      space.Name,
			Location:  space.Location,
			Capacity:  space.Capacity,
			Available: space.Available,
		})
	}

	return &pb.ListSpacesResponse{
		Spaces: pbSpaces,
	}, nil
}
