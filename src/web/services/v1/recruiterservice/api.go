package recruiterservice

import "net/http"

const ()

// RecruiterService interface
type RecruiterService interface {
	CreateRecruiter(w http.ResponseWriter, r *http.Request)
}
