package main

import (
	"booking-service/internal/infrastructure/repository"
	Yerassylgrpc "booking-service/internal/transport/grpc"
	"booking-service/internal/usecase"
	pb "booking-service/proto"
	"context"
	"fmt"
	"log"
	"net"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
)

func main() {
	mongoClient, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	col := mongoClient.Database("coworking").Collection("bookings")
	repo := repository.NewBookingRepo(col)
	useCase := usecase.NewBookingUseCase(repo)
	handler := Yerassylgrpc.NewBookingHandler(useCase)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer()
	pb.RegisterBookingServiceServer(s, handler)
	fmt.Println("Booking service started on :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
