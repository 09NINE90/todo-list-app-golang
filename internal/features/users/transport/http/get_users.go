package users_transport_http

import (
	"fmt"
	"net/http"

	corelogger "github.com/09NINE90/todo-list-app-golang/internal/core/logger"
	corehttpresponse "github.com/09NINE90/todo-list-app-golang/internal/core/transport/http/response"
	corehttputils "github.com/09NINE90/todo-list-app-golang/internal/core/transport/http/utils"
)

type GetUsersRs []UserDTORs

func (h *UsersHttpHandler) GetUsers(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := corelogger.FromContext(ctx)
	responseHandler := corehttpresponse.NewHttpResponseHandler(log, rw)

	limit, offset, err := getLimitOffsetQueryParams(r)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"fail on get `limit` and `offset` query params",
		)
		return
	}

	userDomains, err := h.usersService.GetUsers(ctx, limit, offset)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"fail on get `users` service",
		)
		return
	}

	response := GetUsersRs(usersDTOFromDomains(userDomains))

	responseHandler.JsonResponse(response, http.StatusOK)
}

func getLimitOffsetQueryParams(r *http.Request) (*int, *int, error) {
	limit, err := corehttputils.GetIntQueryParam(r, "limit")
	if err != nil {
		return nil, nil, fmt.Errorf("get limit: %w", err)
	}
	offset, err := corehttputils.GetIntQueryParam(r, "offset")
	if err != nil {
		return nil, nil, fmt.Errorf("get offset: %w", err)
	}

	return limit, offset, nil
}
