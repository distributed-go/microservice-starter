package authservice

import "net/http"

// AuthService interface
type AuthService interface {
	Login(w http.ResponseWriter, r *http.Request)
	Logout(w http.ResponseWriter, r *http.Request)
	SignUp(w http.ResponseWriter, r *http.Request)
	Authenticate(w http.ResponseWriter, r *http.Request)
	Validate(w http.ResponseWriter, r *http.Request)
}
