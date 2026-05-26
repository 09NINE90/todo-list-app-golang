package core_http_response

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	core_errors "github.com/09NINE90/todo-list-app-golang/internal/core/errors"
	corelogger "github.com/09NINE90/todo-list-app-golang/internal/core/logger"
	"go.uber.org/zap"
)

type HttpResponseHandler struct {
	log *corelogger.Logger
	rw  http.ResponseWriter
}

func NewHttpResponseHandler(
	log *corelogger.Logger,
	rw http.ResponseWriter,
) *HttpResponseHandler {
	return &HttpResponseHandler{
		log: log,
		rw:  rw,
	}
}

func (h *HttpResponseHandler) JsonResponse(
	responseBody any,
	statusCode int,
) {
	h.rw.WriteHeader(statusCode)

	if err := json.NewEncoder(h.rw).Encode(responseBody); err != nil {
		h.log.Error("write HTTP response", zap.Error(err))
	}
}

func (h *HttpResponseHandler) ErrorResponse(err error, msg string) {
	var (
		statusCode int
		logFunc    func(string, ...zap.Field)
	)

	switch {
	case errors.Is(err, core_errors.ErrInvalidArgument):
		statusCode = http.StatusBadRequest
		logFunc = h.log.Warn
	case errors.Is(err, core_errors.ErrNotFound):
		statusCode = http.StatusNotFound
		logFunc = h.log.Debug
	case errors.Is(err, core_errors.ErrConflict):
		statusCode = http.StatusConflict
		logFunc = h.log.Warn
	default:
		statusCode = http.StatusInternalServerError
		logFunc = h.log.Error
	}

	logFunc(msg, zap.Error(err))

	h.errorResponse(
		err,
		msg,
		statusCode,
	)
}

func (h *HttpResponseHandler) PanicResponse(p any, msg string) {
	statusCode := http.StatusInternalServerError
	err := fmt.Errorf("unexpected panic: %v", p)

	h.log.Error(msg, zap.Error(err))

	h.errorResponse(
		err,
		msg,
		statusCode,
	)
}

func (h *HttpResponseHandler) errorResponse(
	err error,
	msg string,
	statusCode int,
) {
	response := map[string]string{
		"message": msg,
		"error":   err.Error(),
	}

	h.JsonResponse(
		response,
		statusCode,
	)
}
