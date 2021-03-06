package errorinterface

import (
	"net/http"

	"github.com/go-chi/render"
	validation "github.com/go-ozzo/ozzo-validation"
)

// ErrorResponse renderer type for handling all sorts of errors.
type ErrorResponse struct {
	Err            error `json:"-"` // low-level runtime error
	HTTPStatusCode int   `json:"-"` // http response status code

	Message          string            `json:"Message,omitempty"`          // user-level error message
	Status           string            `json:"Status,omitempty"`           // user-level status message
	Code             int64             `json:"Code,omitempty"`             // application-specific error code
	Error            string            `json:"Error,omitempty"`            // application-level error message, for debugging
	ValidationErrors validation.Errors `json:"ValidationErrors,omitempty"` // user level model validation errors
}

// Render sets the application-specific error code in AppCode.
func (e *ErrorResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}
