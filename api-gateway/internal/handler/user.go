package handler

import (
	"api-gateway/internal/client"
	pb "api-gateway/proto"
	"encoding/json"
	"net/http"
)

type RegisterRequest struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  string `json:"role"`
}

func RegisterUserHandler(w http.ResponseWriter, r *http.Request) {
	var req RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	client, err := client.NewUserClient()
	if err != nil {
		http.Error(w, "cannot connect to user-service", 500)
		return
	}

	_, err = client.RegisterUser(r.Context(), &pb.RegisterRequest{
		Id:    req.ID,
		Name:  req.Name,
		Email: req.Email,
		Role:  req.Role,
	})

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"status":"ok"}`))
}
