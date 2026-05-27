package users_service

import (
	"context"

	"github.com/09NINE90/todo-list-app-golang/internal/core/domain"
)

type IUsersService interface {
	CreateUser(
		ctx context.Context,
		user domain.User,
	) (domain.User, error)

	GetUsers(
		ctx context.Context,
		limit *int,
		offset *int,
	) ([]domain.User, error)
}
