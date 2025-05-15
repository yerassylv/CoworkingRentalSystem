package main

import (
	"context"
	"log"
	"net/http"

	// bookingpb "api-gateway/proto/gen/booking"
	// spacepb "api-gateway/proto/gen/space"
	// userpb "api-gateway/proto/gen/user"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	// "google.golang.org/grpc"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	// opts := []grpc.DialOption{grpc.WithInsecure()}

	// if err := userpb.RegisterUserServiceHandlerFromEndpoint(ctx, mux, "user:50051", opts); err != nil {
	// 	log.Fatalf("failed to register user service: %v", err)
	// }
	// if err := spacepb.RegisterSpaceServiceHandlerFromEndpoint(ctx, mux, "space:50052", opts); err != nil {
	// 	log.Fatalf("failed to register space service: %v", err)
	// }
	// if err := bookingpb.RegisterBookingServiceHandlerFromEndpoint(ctx, mux, "booking:50053", opts); err != nil {
	// 	log.Fatalf("failed to register booking service: %v", err)
	// }

	log.Println("API Gateway listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
