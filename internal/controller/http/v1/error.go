package v1

import (
	"errors"
	"net/http"

	"github.com/ilyasbulat/rest_api/internal/apperror"
	"github.com/julienschmidt/httprouter"
)

func errorResponse(writer http.ResponseWriter, code int, msg []byte) {
	writer.Header().Set(
		"Content-Type",
		"application/json",
	)
	writer.WriteHeader(code)
	writer.Write(msg)
}

type customHandler func(http.ResponseWriter, *http.Request, httprouter.Params) error

func middleware(next customHandler) httprouter.Handle {
	return func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		var (
			appErr *apperror.AppError
		)
		err := next(writer, request, params)
		if err != nil {
			if errors.As(err, &appErr) {
				if errors.Is(err, apperror.ErrNotFound) {
					errorResponse(writer, http.StatusNotFound, apperror.ErrNotFound.Marshal())
					return
				}

				err = err.(*apperror.AppError)
				errorResponse(writer, http.StatusBadRequest, apperror.ErrBadRequest.Marshal())
				return
			}
			errorResponse(writer, http.StatusInternalServerError, apperror.NewAppError(err, err.Error(), "something went wrong").Marshal())
		}
	}
}
