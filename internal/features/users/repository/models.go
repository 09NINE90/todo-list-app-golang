package users_repository

import "github.com/google/uuid"

type UserModel struct {
	ID          uuid.UUID
	FirstName   string
	LastName    string
	PhoneNumber string
}
