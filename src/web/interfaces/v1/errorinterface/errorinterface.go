package errorinterface

import (
	"net/http"

	"github.com/go-chi/render"
)

// ErrorResponse renderer type for handling all sorts of errors.
type ErrorResponse struct {
	Err            error `json:"-"` // low-level runtime error
	HTTPStatusCode int   `json:"-"` // http response status code

	Status string `json:"status,omitempty"` // user-level status message
	Code   int64  `json:"code,omitempty"`   // application-specific error code
	Error  string `json:"error,omitempty"`  // application-level error message, for debugging
}

// Render sets the application-specific error code in AppCode.
func (e *ErrorResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}
