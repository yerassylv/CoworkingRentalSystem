package client

import (
	pb "api-gateway/proto"

	"google.golang.org/grpc"
)

func NewUserClient() (pb.UserServiceClient, error) {
	conn, err := grpc.Dial("user-service:50052", grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	return pb.NewUserServiceClient(conn), nil
}
