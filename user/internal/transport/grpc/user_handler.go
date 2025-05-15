package grpc

import (
	"context"
	"log"

	"user/internal/usecase"
	pb "user/proto"
)

type UserHandler struct {
	pb.UnimplementedUserServiceServer
	uc *usecase.UserUseCase
}

func NewUserHandler(uc *usecase.UserUseCase) *UserHandler {
	return &UserHandler{uc: uc}
}

func (h *UserHandler) GetUserProfile(ctx context.Context, req *pb.GetUserProfileRequest) (*pb.GetUserProfileResponse, error) {
	user, err := h.uc.GetUserProfile(ctx, req.GetUserId())
	if err != nil {
		log.Printf("failed to get user profile: %v", err)
		return nil, err
	}
	if user == nil {
		return nil, nil
	}
	return &pb.GetUserProfileResponse{
		UserId:   user.UserID,
		FullName: user.FullName,
		Email:    user.Email,
		Phone:    user.Phone,
	}, nil
}
