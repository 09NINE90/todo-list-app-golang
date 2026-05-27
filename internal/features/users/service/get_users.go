package users_service

import (
	"context"
	"fmt"

	"github.com/09NINE90/todo-list-app-golang/internal/core/domain"
	coreerrors "github.com/09NINE90/todo-list-app-golang/internal/core/errors"
)

func (s *UsersService) GetUsers(
	ctx context.Context,
	limit *int,
	offset *int,
) ([]domain.User, error) {
	if limit != nil && *limit < 0 {
		return nil, fmt.Errorf(
			"limit must be greater than or equal to 0: %w",
			coreerrors.ErrInvalidArgument,
		)
	}

	if offset != nil && *offset < 0 {
		return nil, fmt.Errorf(
			"offset must be greater than or equal to 0: %w",
			coreerrors.ErrInvalidArgument,
		)
	}

	users, err := s.usersRepository.GetUsers(ctx, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("get Users from repository: %w", err)
	}

	return users, nil
}
