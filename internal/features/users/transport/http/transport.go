package users_transport_http

import (
	"context"
	"net/http"

	"github.com/09NINE90/todo-list-app-golang/internal/core/domain"
	corehttpserver "github.com/09NINE90/todo-list-app-golang/internal/core/transport/http/server"
)

type UsersHttpHandler struct {
	usersService UsersService
}

type UsersService interface {
	CreateUser(
		ctx context.Context,
		user domain.User,
	) (domain.User, error)
}

func NewUsersHttpHandler(
	usersService UsersService,
) *UsersHttpHandler {
	return &UsersHttpHandler{
		usersService: usersService,
	}
}

func (h *UsersHttpHandler) Routers() []corehttpserver.Route {
	return []corehttpserver.Route{
		{
			http.MethodPost,
			"/users",
			h.CreateUser,
		},
	}
}
