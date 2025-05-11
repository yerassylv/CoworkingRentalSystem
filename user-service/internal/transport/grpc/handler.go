package grpc

import (
	"context"
	"user-service/internal/domain"
	"user-service/internal/usecase"
	pb "user-service/proto"
)

type UserHandler struct {
	pb.UnimplementedUserServiceServer
	uc *usecase.UserUseCase
}

func NewUserHandler(uc *usecase.UserUseCase) *UserHandler {
	return &UserHandler{uc: uc}
}

func (h *UserHandler) RegisterUser(ctx context.Context, req *pb.RegisterRequest) (*pb.UserResponse, error) {
	u := &domain.User{
		ID:    req.Id,
		Name:  req.Name,
		Email: req.Email,
		Role:  req.Role,
	}
	if err := h.uc.Register(u); err != nil {
		return nil, err
	}
	return &pb.UserResponse{Id: u.ID, Name: u.Name, Email: u.Email, Role: u.Role}, nil
}

func (h *UserHandler) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.UserResponse, error) {
	u, err := h.uc.GetByID(req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.UserResponse{Id: u.ID, Name: u.Name, Email: u.Email, Role: u.Role}, nil
}

func (h *UserHandler) ListUsers(ctx context.Context, _ *pb.Empty) (*pb.UserListResponse, error) {
	users, err := h.uc.List()
	if err != nil {
		return nil, err
	}
	var resp []*pb.UserResponse
	for _, u := range users {
		resp = append(resp, &pb.UserResponse{
			Id: u.ID, Name: u.Name, Email: u.Email, Role: u.Role,
		})
	}
	return &pb.UserListResponse{Users: resp}, nil
}
