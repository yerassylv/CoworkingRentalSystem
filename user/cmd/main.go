package main

import (
	"context"
	"log"
	"net"
	"time"

	"user/internal/repository"
	Yerassyl "user/internal/transport/grpc"
	"user/internal/usecase"

	pb "user/proto/gen"

	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
)

func main() {
	ctx := context.Background()
	mongoClient, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://mongo:27017"))

	if err != nil {
		log.Fatalf("Mongo connection failed: %v", err)
	}
	db := mongoClient.Database("coworking")
	redisClient := redis.NewClient(&redis.Options{
		Addr: "redis:6379",
		DB:   0,
	})
	repo := repository.NewUserMongoRepository(db.Collection("users"))
	cache := repository.NewRedisCache(redisClient, time.Minute*10)
	uc := usecase.NewUserUseCase(repo, cache)
	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, Yerassyl.NewUserHandler(uc))

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Println("User gRPC service is running on port 50051")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
