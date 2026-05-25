package users_transport_http

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
)

type CreateUserRq struct {
	FirstName   string `json:"fullName"`
	LastName    string `json:"lastName"`
	PhoneNumber string `json:"phoneNumber"`
}

type CreateUserRs struct {
	ID          uuid.UUID `json:"id"`
	FirstName   string    `json:"fullName"`
	LastName    string    `json:"lastName"`
	PhoneNumber string    `json:"phoneNumber"`
}

func (h *UsersHttpHandler) CreateUser(rw http.ResponseWriter, r *http.Request) {
	var request CreateUserRq
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {

	}
}
