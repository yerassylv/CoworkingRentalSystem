package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	bookingpb "api-gateway/proto/gen/booking"
	spacepb "api-gateway/proto/gen/space"
	userpb "api-gateway/proto/gen/user"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// Подключение к gRPC-сервисам
	userConn, err := grpc.Dial("user:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to user service: %v", err)
	}
	defer userConn.Close()
	userClient := userpb.NewUserServiceClient(userConn)

	spaceConn, err := grpc.Dial("space:50052", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to space service: %v", err)
	}
	defer spaceConn.Close()
	spaceClient := spacepb.NewSpaceServiceClient(spaceConn)

	bookingConn, err := grpc.Dial("booking:50053", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to booking service: %v", err)
	}
	defer bookingConn.Close()
	bookingClient := bookingpb.NewBookingServiceClient(bookingConn)

	// HTTP → gRPC маршруты

	http.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		userID := r.URL.Query().Get("id")
		if userID == "" {
			http.Error(w, "missing user_id", http.StatusBadRequest)
			return
		}
		resp, err := userClient.GetUserProfile(context.Background(), &userpb.GetUserProfileRequest{
			UserId: userID,
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(resp)
	})

	http.HandleFunc("/space", func(w http.ResponseWriter, r *http.Request) {
		spaceID := r.URL.Query().Get("id")
		if spaceID == "" {
			http.Error(w, "missing space_id", http.StatusBadRequest)
			return
		}
		resp, err := spaceClient.GetSpaceByID(context.Background(), &spacepb.GetSpaceByIDRequest{
			SpaceId: spaceID,
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(resp)
	})

	http.HandleFunc("/bookings", func(w http.ResponseWriter, r *http.Request) {
		userID := r.URL.Query().Get("user_id")
		if userID == "" {
			http.Error(w, "missing user_id", http.StatusBadRequest)
			return
		}
		resp, err := bookingClient.ListBookings(context.Background(), &bookingpb.ListBookingsRequest{
			UserId: userID,
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(resp)
	})

	http.HandleFunc("/book", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var req bookingpb.CreateBookingRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "invalid JSON", http.StatusBadRequest)
			return
		}

		resp, err := bookingClient.CreateBooking(context.Background(), &req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(resp)
	})

	log.Println("🚀 API Gateway running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
