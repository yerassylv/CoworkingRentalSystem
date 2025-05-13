package main

import (
	"context"
	"log"
	snats "space-service/internal/infrastructure/nats"
	"space-service/internal/infrastructure/repository"
	"space-service/internal/usecase"
	"time"

	"github.com/nats-io/nats.go"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	mongoClient, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://mongo:27017"))
	if err != nil {
		log.Fatal(err)
	}

	col := mongoClient.Database("coworking").Collection("spaces")
	repo := repository.NewMongoRepo(col)
	useCase := usecase.NewSpaceUseCase(repo)

	natsConn, err := nats.Connect("nats://nats:4222")
	if err != nil {
		log.Fatal(err)
	}

	snats.SubscribeToBookingCreated(natsConn, useCase.HandleBooking)

	log.Println("Space service is listening to booking.created events...")
}
