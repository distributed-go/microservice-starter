package renderers

import (
	"net/http"

	"github.com/jobbox-tech/recruiter-api/proto/v1/error/v1error"

	"github.com/go-chi/render"
)

type errorResponse struct {
	*v1error.ErrorResponse
}

// Render sets the application-specific error code in AppCode.
func (e *errorResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, int(e.Code))
	return nil
}

// ErrorInvalidRequest returns status 422 Unprocessable Entity including error message.
func ErrorInvalidRequest(err error) render.Renderer {
	return &errorResponse{
		&v1error.ErrorResponse{
			Code:       http.StatusUnprocessableEntity,
			StatusText: http.StatusText(http.StatusUnprocessableEntity),
			Error:      err.Error(),
		},
	}
}

// ErrorUnauthorized renders status 401 Unauthorized with custom error message.
func ErrorUnauthorized(err error) render.Renderer {
	return &errorResponse{
		&v1error.ErrorResponse{
			Code:       http.StatusUnauthorized,
			StatusText: http.StatusText(http.StatusUnauthorized),
			Error:      err.Error(),
		},
	}
}

// ErrorForbidden renders status 403 forbidden with custom error message.
func ErrorForbidden(err error) render.Renderer {
	return &errorResponse{
		&v1error.ErrorResponse{
			Code:       http.StatusForbidden,
			StatusText: http.StatusText(http.StatusForbidden),
			Error:      err.Error(),
		},
	}
}

// ErrorBadRequest return status 400 Bad Request for malformed request body.
func ErrorBadRequest(err error) render.Renderer {
	return &errorResponse{
		&v1error.ErrorResponse{
			Code:       http.StatusBadRequest,
			StatusText: http.StatusText(http.StatusBadRequest),
			Error:      err.Error(),
		},
	}
}

// ErrorNotFound returns status 404 Not Found for invalid resource request.
func ErrorNotFound(err error) render.Renderer {
	return &errorResponse{
		&v1error.ErrorResponse{
			Code:       http.StatusNotFound,
			StatusText: http.StatusText(http.StatusNotFound),
			Error:      err.Error(),
		},
	}
}

// ErrorInternalServerError returns status 500 Internal Server Error.
func ErrorInternalServerError(err error) render.Renderer {
	return &errorResponse{
		&v1error.ErrorResponse{
			Code:       http.StatusInternalServerError,
			StatusText: http.StatusText(http.StatusInternalServerError),
			Error:      err.Error(),
		},
	}
}
