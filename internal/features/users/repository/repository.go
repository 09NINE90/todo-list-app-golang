package users_repository

import corepoolconn "github.com/09NINE90/todo-list-app-golang/internal/core/repository/pool"

type UsersRepository struct {
	pool corepoolconn.Pool
}

func NewUsersRepository(
	pool corepoolconn.Pool,
) *UsersRepository {
	return &UsersRepository{
		pool: pool,
	}
}
