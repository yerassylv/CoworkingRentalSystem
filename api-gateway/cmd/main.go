package main

import (
	"api-gateway/internal/handler"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/register", handler.RegisterUserHandler).Methods("POST")

	fmt.Println("API Gateway running on :8080")
	http.ListenAndServe(":8080", r)
}
