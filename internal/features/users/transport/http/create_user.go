package users_transport_http

import (
	"net/http"

	"github.com/09NINE90/todo-list-app-golang/internal/core/domain"
	corelogger "github.com/09NINE90/todo-list-app-golang/internal/core/logger"
	corehttprequest "github.com/09NINE90/todo-list-app-golang/internal/core/transport/http/request"
	corehttpresponse "github.com/09NINE90/todo-list-app-golang/internal/core/transport/http/response"
	"github.com/google/uuid"
)

type CreateUserRq struct {
	FirstName   string `json:"fullName" validate:"required,min=3,max=50"`
	LastName    string `json:"lastName" validate:"required,min=3,max=50"`
	PhoneNumber string `json:"phoneNumber" validate:"required,min=11,max=15,startswith=+"`
}

type CreateUserRs struct {
	ID          uuid.UUID `json:"id"`
	FirstName   string    `json:"fullName"`
	LastName    string    `json:"lastName"`
	PhoneNumber string    `json:"phoneNumber"`
}

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
	}

	response := dtoFromDomain(userDomain)

	responseHandler.JsonResponse(response, http.StatusCreated)
}

func domainFromDTO(dto CreateUserRq) domain.User {
	return domain.NewUserUninitialized(
		dto.FirstName,
		dto.LastName,
		dto.PhoneNumber,
	)
}

func dtoFromDomain(user domain.User) CreateUserRs {
	return CreateUserRs{
		ID:          user.ID,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		PhoneNumber: user.PhoneNumber,
	}
}
