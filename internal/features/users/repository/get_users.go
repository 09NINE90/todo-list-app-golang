package users_repository

import (
	"context"
	"fmt"

	"github.com/09NINE90/todo-list-app-golang/internal/core/domain"
)

func (r *UsersRepository) GetUsers(
	ctx context.Context,
	limit *int,
	offset *int,
) ([]domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
		SELECT id, first_name, last_name, phone_number
		FROM todoapp.users
		ORDER BY last_name, first_name
		LIMIT $1 OFFSET $2;
	`

	rows, err := r.pool.Query(
		ctx,
		query,
		limit,
		offset,
	)

	if err != nil {
		return nil, fmt.Errorf("get users: %w", err)
	}
	defer rows.Close()

	var users []domain.User
	for rows.Next() {
		var userModel UserModel
		err := rows.Scan(
			&userModel.ID,
			&userModel.FirstName,
			&userModel.LastName,
			&userModel.PhoneNumber,
		)
		if err != nil {
			return nil, fmt.Errorf("scan users: %w", err)
		}

		users = append(users, domain.NewUser(
			userModel.ID,
			userModel.FirstName,
			userModel.LastName,
			userModel.PhoneNumber,
		))
	}

	return users, nil
}
