package main

import (
	"context"
	"log"
	"net"
	"time"

	"space/internal/repository"
	"space/internal/transport/grpc"
	"space/internal/usecase"

	pb "space/proto"

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
	repo := repository.NewSpaceMongoRepository(db.Collection("spaces"))
	cache := repository.NewRedisCache(redisClient, time.Minute*10)
	uc := usecase.NewSpaceUseCase(repo, cache)
	grpcServer := grpc.NewServer()
	pb.RegisterSpaceServiceServer(grpcServer, grpc.NewSpaceHandler(uc))

	listener, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Println("Space gRPC service is running on port 50052")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
