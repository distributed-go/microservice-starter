package recruiterservice

import (
	"errors"
	"net/http"
	"time"

	"github.com/jobbox-tech/recruiter-api/dal/recruitersdal"

	"github.com/go-chi/render"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/jobbox-tech/recruiter-api/database/dbmodels"
	"github.com/jobbox-tech/recruiter-api/web/renderers"

	"github.com/jobbox-tech/recruiter-api/logging"
)

type recruiterservice struct {
	logger        logging.Logger
	recruitersDal recruitersdal.RecruitersDal
}

// NewRecruiterService returns service impl
func NewRecruiterService() RecruiterService {
	return &recruiterservice{
		logger:        logging.NewLogger(),
		recruitersDal: recruitersdal.NewRecruitersDal(),
	}
}

func (rs *recruiterservice) CreateRecruiter(w http.ResponseWriter, r *http.Request) {
	txID := r.Header["transaction_id"][0]
	data := &recruitersRequest{}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, renderers.ErrorInvalidRequest(err, "Invalid request body"))
		return
	}

	t := time.Now().UTC()
	data.Recruiters.Roles = []dbmodels.Role{dbmodels.USER}
	data.Recruiters.Active = true
	data.Recruiters.CreatedTimestampUTC = &t
	data.Recruiters.UpdatedTimestampUTC = &t

	if err := data.Validate(); err != nil {
		switch err.(type) {
		case validation.Errors:
			render.Render(w, r, renderers.ErrorValidation(
				errors.New("Failed to validate the data provided in body"),
				err.(validation.Errors),
				"Incorrect details provided, please provide correct details",
			))
			return
		}
		render.Render(w, r, renderers.ErrorInvalidRequest(err, "Invalid request body"))
		return
	}

	objectID, err := rs.recruitersDal.Create(txID, data.Recruiters)
	if err != nil {
		rs.logger.Error(txID, "").Errorf("Failed to create recruiters record with error %v", err)
		render.Render(w, r, renderers.ErrorInternalServerError("Failed to create recruiter account, please try again"))
		return
	}

	data.Recruiters.ID = objectID
	render.Respond(w, r, newRecruitersResponse(data.Recruiters))
}

// ==============  Bindings  ===============
type recruitersRequest struct {
	*dbmodels.Recruiters
}

func (d *recruitersRequest) Bind(r *http.Request) error {
	return nil
}

type recruitersResponse struct {
	*dbmodels.Recruiters
}

func newRecruitersResponse(a *dbmodels.Recruiters) *recruitersResponse {
	resp := &recruitersResponse{Recruiters: a}
	return resp
}
