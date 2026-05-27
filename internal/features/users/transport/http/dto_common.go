package users_transport_http

import (
	"github.com/09NINE90/todo-list-app-golang/internal/core/domain"
	"github.com/google/uuid"
)

type UserDTORq struct {
	FirstName   string `json:"firstName" validate:"required,min=3,max=50"`
	LastName    string `json:"lastName" validate:"required,min=3,max=50"`
	PhoneNumber string `json:"phoneNumber" validate:"required,min=11,max=15,startswith=+"`
}

type UserDTORs struct {
	ID          uuid.UUID `json:"id"`
	FirstName   string    `json:"firstName"`
	LastName    string    `json:"lastName"`
	PhoneNumber string    `json:"phoneNumber"`
}

func userDTOFromDomain(user domain.User) UserDTORs {
	return UserDTORs{
		ID:          user.ID,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		PhoneNumber: user.PhoneNumber,
	}
}

func usersDTOFromDomains(users []domain.User) []UserDTORs {
	usersDTO := make([]UserDTORs, len(users))

	for i, user := range users {
		usersDTO[i] = userDTOFromDomain(user)
	}

	return usersDTO
}
