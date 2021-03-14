package authinterface

import (
	"net/http"
	"strings"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

// LoginReqInterface holds the login details
type LoginReqInterface struct {
	Email string `json:"Email"`
}

// AuthenticateReqInterface holds the login details
type AuthenticateReqInterface struct {
	Token string `json:"Token"`
}

// AuthenticateResInterface holds the token pair
type AuthenticateResInterface struct {
	AccessToken  string `json:"AccessToken"`
	RefreshToken string `json:"RefreshToken"`
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
