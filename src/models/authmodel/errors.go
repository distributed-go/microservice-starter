package authmodel

import "errors"

// The list of error types presented to the end user as error message.
var (
	ErrInvalidLogin        = errors.New("invalid email address")
	ErrUnknownLogin        = errors.New("email not registered")
	ErrLoginDisabled       = errors.New("login for account disabled")
	ErrLoginToken          = errors.New("invalid or expired login token")
	ErrInvalidRefreshToken = errors.New("invalid refresh token")
	ErrInsufficientRights  = errors.New("token insufficient provoledges")
)
