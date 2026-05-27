package users_transport_http

import (
	"net/http"

	corehttpserver "github.com/09NINE90/todo-list-app-golang/internal/core/transport/http/server"
	usersservice "github.com/09NINE90/todo-list-app-golang/internal/features/users/service"
)

type UsersHttpHandler struct {
	usersService usersservice.IUsersService
}

func NewUsersHttpHandler(
	usersService usersservice.IUsersService,
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
		{
			http.MethodGet,
			"/users",
			h.GetUsers,
		},
	}
}
