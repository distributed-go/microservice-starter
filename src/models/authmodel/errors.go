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
	ErrServerError         = errors.New("Something went wrong please try again after some time")
)

// List of error codes used in authentication service/model
var (
	FailedToCreateAccessToken = "Failed-To-CreateAccess-Token"
	FailedToAuthenticateToken = "Failed-To-Authenticate-Token"
	FailedToSignUp            = "Failed-To-Sign-Up"
)
