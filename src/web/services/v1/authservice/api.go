package authservice

import "net/http"

// AuthService interface
type AuthService interface {
	Logout(w http.ResponseWriter, r *http.Request)
	Login(w http.ResponseWriter, r *http.Request)
}
