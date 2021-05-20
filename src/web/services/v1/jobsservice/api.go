package jobsservice

import (
	"errors"
	"net/http"
)

// JobService interface
type JobsService interface {
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
}

// The list of error types presented to the end user as error message.
var (
	ErrIncompleteDetails = errors.New("Incorrect details provided, please provice correct details")
)

// List of error codes used in jobs service/model
var (
	FailedToCreateJob = "Failed-To-Create-Job"
	FailedToGetJob    = "Failed-To-Get-Job"
	FailedToUpdateJob = "Failed-To-Update-Job"
	FailedToDeleteJob = "Failed-To-Delete-Job"
)
