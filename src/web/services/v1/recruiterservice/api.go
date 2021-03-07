package recruiterservice

import "net/http"

// RecruiterService interface
type RecruiterService interface {
	CreateRecruiter(w http.ResponseWriter, r *http.Request)
}
