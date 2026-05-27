package core_http_utils

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	coreerrors "github.com/09NINE90/todo-list-app-golang/internal/core/errors"
)

func GetIntQueryParam(r *http.Request, key string) (*int, error) {
	param := r.URL.Query().Get(key)

	if param == "" {
		return nil, nil
	}

	val, err := strconv.Atoi(param)
	if err != nil {
		return nil, fmt.Errorf(
			"query param '%s' is not a valid integer: %v :%w",
			key,
			err,
			coreerrors.ErrInvalidArgument,
		)
	}

	return &val, nil
}

func GetTimeQueryParam(r *http.Request, key string) (*time.Time, error) {
	param := r.URL.Query().Get(key)

	if param == "" {
		return nil, nil
	}

	val, err := time.Parse(time.RFC3339, param)
	if err != nil {
		return nil, fmt.Errorf(
			"query param '%s' is not a valid time: %v :%w",
			key,
			err,
			coreerrors.ErrInvalidArgument,
		)
	}

	return &val, nil
}
