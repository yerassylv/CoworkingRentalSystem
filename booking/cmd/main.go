package main

import (
	"booking/internal/repository"
	"booking/internal/transport/grpc"
	"booking/internal/usecase"
	"context"
	"log"
	"net"
	"time"

	pb "booking/proto"

	"github.com/go-redis/redis/v8"
	"github.com/nats-io/nats.go"
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
	nc, err := nats.Connect("nats://nats:4222")
	if err != nil {
		log.Fatalf("NATS connection failed: %v", err)
	}
	defer nc.Close()
	repo := repository.NewBookingMongoRepository(db.Collection("bookings"))
	cache := repository.NewRedisCache(redisClient, time.Minute*10)
	uc := usecase.NewBookingUseCase(repo, cache, nc)
	grpcServer := grpc.NewServer()
	pb.RegisterBookingServiceServer(grpcServer, grpc.NewBookingHandler(uc))

	listener, err := net.Listen("tcp", ":50053")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Println("Booking gRPC service is running on port 50053")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
