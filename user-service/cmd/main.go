package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"user-service/internal/infrastructure/repository"
	tgrpc "user-service/internal/transport/grpc"
	"user-service/internal/usecase"

	pb "user-service/proto"

	"google.golang.org/grpc"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://mongo:27017"))
	if err != nil {
		log.Fatal(err)
	}

	col := client.Database("coworking").Collection("users")
	repo := repository.NewMongoRepo(col)
	uc := usecase.NewUserUseCase(repo)
	handler := tgrpc.NewUserHandler(uc)

	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, handler)

	fmt.Println("User service running on :50052")
	if err := s.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
