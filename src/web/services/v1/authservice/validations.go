package authservice

import (
	"net/http"
	"strings"

	"github.com/jobbox-tech/recruiter-api/proto/v1/auth/v1auth"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type loginRequest v1auth.LoginRequest
type authenticateRequest v1auth.AuthenticateRequest
type signupRequest v1auth.SignUpRequest

// Bind valide the login request interface with rules given
func (body *loginRequest) Bind(r *http.Request) error {
	body.Email = strings.TrimSpace(body.Email)
	body.Email = strings.ToLower(body.Email)
	return validation.ValidateStruct(body,
		validation.Field(&body.Email, validation.Required, is.Email),
	)
}

// Bind valide the authenticate request interface with rules given
func (body *authenticateRequest) Bind(r *http.Request) error {
	body.Token = strings.TrimSpace(body.Token)
	return validation.ValidateStruct(body,
		validation.Field(&body.Token, validation.Required, is.UUID),
	)
}

// Bind valide the sign up request interface with rules given
func (body *signupRequest) Bind(r *http.Request) error {
	return validation.ValidateStruct(body,
		validation.Field(&body.CompanyName, validation.Required, validation.Length(1, 50)),
		validation.Field(&body.Email, validation.Required, is.Email),
		validation.Field(&body.FirstName, validation.Required, validation.Length(1, 50)),
		validation.Field(&body.Designation, validation.Required, validation.Length(1, 256)),
	)
}
