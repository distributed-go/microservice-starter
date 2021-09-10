package renderers

import (
	"net/http"

	"github.com/distributed-go/microservice-starter/web/interfaces/v1/errorinterface"
	"github.com/go-chi/render"
)

// ErrorInvalidRequest returns status 422 Unprocessable Entity including error message.
func ErrorInvalidRequest(err error) render.Renderer {
	return &errorinterface.ErrorResponse{
		Err:            err,
		HTTPStatusCode: http.StatusUnprocessableEntity,
		Status:         http.StatusText(http.StatusUnprocessableEntity),
		Error:          err.Error(),
	}
}

// ErrorUnauthorized renders status 401 Unauthorized with custom error message.
func ErrorUnauthorized(err error) render.Renderer {
	return &errorinterface.ErrorResponse{
		Err:            err,
		HTTPStatusCode: http.StatusUnauthorized,
		Status:         http.StatusText(http.StatusUnauthorized),
		Error:          err.Error(),
	}
}

// ErrorForbidden renders status 403 forbidden with custom error message.
func ErrorForbidden(err error) render.Renderer {
	return &errorinterface.ErrorResponse{
		Err:            err,
		Error:          err.Error(),
		HTTPStatusCode: http.StatusForbidden,
		Status:         http.StatusText(http.StatusForbidden),
	}
}

// ErrorBadRequest return status 400 Bad Request for malformed request body.
func ErrorBadRequest(err error) render.Renderer {
	return &errorinterface.ErrorResponse{
		HTTPStatusCode: http.StatusBadRequest,
		Status:         http.StatusText(http.StatusBadRequest),
		Err:            err,
		Error:          err.Error(),
	}
}

// ErrorNotFound returns status 404 Not Found for invalid resource request.
func ErrorNotFound(err error) render.Renderer {
	return &errorinterface.ErrorResponse{
		HTTPStatusCode: http.StatusNotFound,
		Status:         http.StatusText(http.StatusNotFound),
		Err:            err,
		Error:          err.Error(),
	}
}

// ErrorInternalServerError returns status 500 Internal Server Error.
func ErrorInternalServerError(err error) render.Renderer {
	return &errorinterface.ErrorResponse{
		HTTPStatusCode: http.StatusInternalServerError,
		Status:         http.StatusText(http.StatusInternalServerError),
		Err:            err,
		Error:          err.Error(),
	}
}
