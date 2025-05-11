package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"user-service/infrastructure/repository"
	"user-service/internal/usecase"
	"user-service/transport/grpc"

	pb "user-service/proto"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
)

func main() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://mongo:27017"))
	if err != nil {
		log.Fatal(err)
	}

	col := client.Database("coworking").Collection("users")
	repo := repository.NewMongoRepo(col)
	uc := usecase.NewUserUseCase(repo)
	handler := grpc.NewUserHandler(uc)

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
