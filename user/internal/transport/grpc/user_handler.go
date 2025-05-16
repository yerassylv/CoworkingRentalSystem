package grpc

import (
	"context"
	"log"

	"user/internal/usecase"
	pb "user/proto/gen"
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

func (h *UserHandler) RegisterUser(ctx context.Context, req *pb.RegisterUserRequest) (*pb.RegisterUserResponse, error) {
	user, err := h.uc.RegisterUser(ctx, req.GetFullName(), req.GetEmail(), req.GetPhone(), req.GetPassword())
	if err != nil {
		log.Printf("failed to register user: %v", err)
		return nil, err
	}

	return &pb.RegisterUserResponse{
		UserId: user.UserID,
	}, nil
}

func (h *UserHandler) LoginUser(ctx context.Context, req *pb.LoginUserRequest) (*pb.LoginUserResponse, error) {
	token, user, err := h.uc.LoginUser(ctx, req.GetEmail(), req.GetPassword())
	if err != nil {
		log.Printf("failed to login user: %v", err)
		return nil, err
	}

	return &pb.LoginUserResponse{
		Token:    token,
		UserId:   user.UserID,
		FullName: user.FullName,
		Email:    user.Email,
		Phone:    user.Phone,
	}, nil
}
