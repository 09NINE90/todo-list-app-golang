package users_transport_http

import (
	"net/http"

	"github.com/09NINE90/todo-list-app-golang/internal/core/domain"
	corelogger "github.com/09NINE90/todo-list-app-golang/internal/core/logger"
	corehttprequest "github.com/09NINE90/todo-list-app-golang/internal/core/transport/http/request"
	corehttpresponse "github.com/09NINE90/todo-list-app-golang/internal/core/transport/http/response"
)

type CreateUserRq UserDTORq
type CreateUserRs UserDTORs

func (h *UsersHttpHandler) CreateUser(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := corelogger.FromContext(ctx)
	responseHandler := corehttpresponse.NewHttpResponseHandler(log, rw)

	log.Debug("invoke CreateUser handler")

	var request CreateUserRq
	if err := corehttprequest.DecodeAndValidateRequest(r, &request); err != nil {
		responseHandler.ErrorResponse(err, "failed to decode and validate HTTP request")

		return
	}

	userDomain := domainFromDTO(request)

	userDomain, err := h.usersService.CreateUser(ctx, userDomain)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to create user")
		return
	}

	response := CreateUserRs(userDTOFromDomain(userDomain))

	responseHandler.JsonResponse(response, http.StatusCreated)
}

func domainFromDTO(dto CreateUserRq) domain.User {
	return domain.NewUserUninitialized(
		dto.FirstName,
		dto.LastName,
		dto.PhoneNumber,
	)
}
