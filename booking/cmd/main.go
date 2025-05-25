package main

import (
	"booking/internal/infra/email"
	"booking/internal/repository"
	Yerassylgrpc "booking/internal/transport/grpc"
	"booking/internal/usecase"
	pb "booking/proto/gen"
	"context"
	"log"
	"net"
	"os"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/joho/godotenv"
	"github.com/nats-io/nats.go"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
)

func main() {
	// Load environment variables from .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	ctx := context.Background()

	// MongoDB connection
	mongoClient, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://mongo:27017"))
	if err != nil {
		log.Fatalf("Mongo connection failed: %v", err)
	}
	db := mongoClient.Database("coworking")

	// Redis cache
	redisClient := redis.NewClient(&redis.Options{
		Addr: "redis:6379",
		DB:   0,
	})

	// NATS connection
	nc, err := nats.Connect("nats://nats:4222")
	if err != nil {
		log.Fatalf("NATS connection failed: %v", err)
	}
	defer nc.Close()

	// Parse SMTP port from environment
	smtpPort, err := strconv.Atoi(os.Getenv("SMTP_PORT"))
	if err != nil {
		log.Fatalf("Invalid SMTP_PORT: %v", err)
	}

	// Create SMTP mailer
	mailer := email.NewSmtpSender(
		os.Getenv("SMTP_HOST"),
		smtpPort,
		os.Getenv("SMTP_USER"),
		os.Getenv("SMTP_PASS"),
		os.Getenv("SMTP_FROM"),
	)

	// Initialize repository, cache, usecase
	repo := repository.NewBookingMongoRepository(db.Collection("bookings"))
	cache := repository.NewRedisCache(redisClient, time.Minute*10)
	uc := usecase.NewBookingUseCase(repo, cache, nc, mailer)

	// Set up gRPC server
	grpcServer := grpc.NewServer()
	pb.RegisterBookingServiceServer(grpcServer, Yerassylgrpc.NewBookingHandler(uc))

	listener, err := net.Listen("tcp", ":50053")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Println("Booking gRPC service is running on port 50053")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
