package users_repository

import (
	"context"
	"fmt"

	"github.com/09NINE90/todo-list-app-golang/internal/core/domain"
)

func (r *UsersRepository) CreateUser(
	ctx context.Context,
	user domain.User,
) (domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
		INSERT INTO todoapp.users (first_name, last_name, phone_number)
		VALUES ($1, $2, $3)
		RETURNING id, first_name, last_name, phone_number;
	`

	row := r.pool.QueryRow(
		ctx,
		query,
		user.FirstName,
		user.LastName,
		user.PhoneNumber,
	)

	var userModel UserModel
	err := row.Scan(
		&userModel.ID,
		&userModel.FirstName,
		&userModel.LastName,
		&userModel.PhoneNumber,
	)

	if err != nil {
		return domain.User{}, fmt.Errorf("scan Error: %w", err)
	}

	userDomain := domain.NewUser(
		userModel.ID,
		userModel.FirstName,
		userModel.LastName,
		userModel.PhoneNumber,
	)

	return userDomain, nil
}
