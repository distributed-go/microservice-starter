package jobsinterface

import (
	"net/http"

	"github.com/distributed-go/microservice-starter/database/dbmodels"
	validation "github.com/go-ozzo/ozzo-validation"
)

// JobRequest ...
type JobRequest struct {
	dbmodels.Job
}

// JobResponse ...
type JobResponse struct {
	*dbmodels.Job
}

// ============== Validation ============== //

// Bind valide the login request interface with rules given
func (body *JobRequest) Bind(r *http.Request) error {
	return validation.ValidateStruct(body,
		validation.Field(&body.Job.Title, validation.Required, validation.Length(1, 256)),
		validation.Field(&body.Job.Locations, validation.Required, validation.Length(1, 48)),
		validation.Field(&body.Job.MustHaveSkills, validation.Required, validation.Length(1, 48)),
		validation.Field(&body.Job.Category, validation.Required, validation.Length(1, 48)),
		validation.Field(&body.Job.EmploymentType, validation.Required, validation.Length(1, 48)),
	)
}
