package authmodel

import "errors"

// The list of error types presented to the end user as error message.
var (
	ErrInvalidLogin        = errors.New("Incorrect email address provided")
	ErrUnknownLogin        = errors.New("We could not find any account associated with given email address")
	ErrAlreadyRegistered   = errors.New("This email address is already used, please use a different email address")
	ErrLoginDisabled       = errors.New("Login for this account has been disabled")
	ErrLoginToken          = errors.New("Invalid or expired login token provided")
	ErrInvalidRefreshToken = errors.New("Invalid refresh token")
	ErrInsufficientRights  = errors.New("token insufficient proviledges")
	ErrServerError         = errors.New("Something went wrong, please try again")
	ErrIncorrectDetails    = errors.New("Incorrect details provided, please provide correct details")
)

// List of error codes used in authentication service/model
var (
	FailedToCreateAccessToken = "Failed-To-CreateAccess-Token"
	FailedToAuthenticateToken = "Failed-To-Authenticate-Token"
	FailedToSignUp            = "Failed-To-Sign-Up"
	FailedToDeleteToken       = "Failed-To-Delete-Token"
)
