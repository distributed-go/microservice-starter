package authinterface

import (
	"net/http"
	"strings"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

// LoginReqInterface holds the login details
type LoginReqInterface struct {
	Email string `json:"email"`
}

// SignUpReqInterface sign up request details
type SignUpReqInterface struct {
	CompanyName string `json:"companyName,omitempty"`
	Email       string `json:"email,omitempty"`
	FirstName   string `json:"firstName,omitempty"`
	Designation string `json:"designation,omitempty"`
}

// AuthenticateReqInterface holds the login details
type AuthenticateReqInterface struct {
	Token string `json:"token"`
}

// AuthenticateResInterface holds the token pair
type AuthenticateResInterface struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

// ============== Validation ============== //

// Bind valide the login request interface with rules given
func (body *LoginReqInterface) Bind(r *http.Request) error {
	body.Email = strings.TrimSpace(body.Email)
	body.Email = strings.ToLower(body.Email)
	return validation.ValidateStruct(body,
		validation.Field(&body.Email, validation.Required, is.Email),
	)
}

// Bind valide the authenticate request interface with rules given
func (body *AuthenticateReqInterface) Bind(r *http.Request) error {
	body.Token = strings.TrimSpace(body.Token)
	return validation.ValidateStruct(body,
		validation.Field(&body.Token, validation.Required, is.UUID),
	)
}

// Bind valide the sign up request interface with rules given
func (body *SignUpReqInterface) Bind(r *http.Request) error {
	return validation.ValidateStruct(body,
		validation.Field(&body.CompanyName, validation.Required, validation.Length(1, 50)),
		validation.Field(&body.Email, validation.Required, is.Email),
		validation.Field(&body.FirstName, validation.Required, validation.Length(1, 50)),
		validation.Field(&body.Designation, validation.Required, validation.Length(1, 256)),
	)
}
