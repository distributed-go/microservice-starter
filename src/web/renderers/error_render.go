package renderers

import (
	"net/http"

	"github.com/go-chi/render"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/jobbox-tech/recruiter-api/web/interfaces/v1/errorinterface"
)

// ErrorInvalidRequest returns status 422 Unprocessable Entity including error message.
func ErrorInvalidRequest(err error, message string) render.Renderer {
	return &errorinterface.ErrorResponse{
		Err:            err,
		HTTPStatusCode: http.StatusUnprocessableEntity,
		Message:        message,
		Status:         http.StatusText(http.StatusUnprocessableEntity),
		Error:          err.Error(),
	}
}

// ErrorRender returns status 422 Unprocessable Entity rendering response error.
func ErrorRender(err error, message string) render.Renderer {
	return &errorinterface.ErrorResponse{
		Err:            err,
		HTTPStatusCode: http.StatusUnprocessableEntity,
		Message:        message,
		Status:         "Error rendering response.",
		Error:          err.Error(),
	}
}

// ErrorValidation returns status 422 Unprocessable Entity stating validation errors.
func ErrorValidation(err error, valErr validation.Errors, message string) render.Renderer {
	return &errorinterface.ErrorResponse{
		Err:              err,
		HTTPStatusCode:   http.StatusUnprocessableEntity,
		Message:          message,
		Status:           http.StatusText(http.StatusUnprocessableEntity),
		Error:            err.Error(),
		ValidationErrors: valErr,
	}
}

// ErrorBadRequest return status 400 Bad Request for malformed request body.
func ErrorBadRequest(message string) render.Renderer {
	return &errorinterface.ErrorResponse{
		HTTPStatusCode: http.StatusBadRequest,
		Status:         http.StatusText(http.StatusBadRequest),
		Message:        message,
	}
}

// ErrorNotFound returns status 404 Not Found for invalid resource request.
func ErrorNotFound(message string) render.Renderer {
	return &errorinterface.ErrorResponse{
		HTTPStatusCode: http.StatusNotFound,
		Status:         http.StatusText(http.StatusNotFound),
		Message:        message,
	}
}

// ErrorInternalServerError returns status 500 Internal Server Error.
func ErrorInternalServerError(message string) render.Renderer {
	return &errorinterface.ErrorResponse{
		HTTPStatusCode: http.StatusInternalServerError,
		Message:        message,
		Status:         http.StatusText(http.StatusInternalServerError),
	}
}
