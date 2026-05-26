package domain

import (
	"fmt"
	"regexp"

	coreerrors "github.com/09NINE90/todo-list-app-golang/internal/core/errors"
	"github.com/google/uuid"
)

type User struct {
	ID          uuid.UUID
	FirstName   string
	LastName    string
	PhoneNumber string
}

func NewUser(
	id uuid.UUID,
	firstname string,
	lastName string,
	phoneNumber string) User {
	return User{
		ID:          id,
		FirstName:   firstname,
		LastName:    lastName,
		PhoneNumber: phoneNumber,
	}
}

func NewUserUninitialized(
	firstname string,
	lastName string,
	phoneNumber string) User {
	return NewUser(
		UninitializedID,
		firstname,
		lastName,
		phoneNumber,
	)
}

func (u *User) Validate() error {
	firstNameLength := len([]rune(u.FirstName))
	if firstNameLength < 3 || firstNameLength > 50 {
		return fmt.Errorf(
			"invalid `FirstName` length: %d, %w",
			firstNameLength,
			coreerrors.ErrInvalidArgument,
		)
	}

	lastNameLength := len([]rune(u.LastName))
	if lastNameLength < 3 || lastNameLength > 50 {
		return fmt.Errorf(
			"invalid `LastName` length: %d, %w",
			lastNameLength,
			coreerrors.ErrInvalidArgument,
		)
	}

	phoneNumberLength := len([]rune(u.PhoneNumber))
	if phoneNumberLength < 11 || phoneNumberLength > 15 {
		return fmt.Errorf(
			"invalid `PhoneNumber` length: %d, %w",
			phoneNumberLength,
			coreerrors.ErrInvalidArgument,
		)
	}

	re := regexp.MustCompile(`^\+[0-9]+$`)
	if !re.MatchString(u.PhoneNumber) {
		return fmt.Errorf(
			"invalid `PhoneNumber` format: %w",
			coreerrors.ErrInvalidArgument,
		)
	}

	return nil
}
