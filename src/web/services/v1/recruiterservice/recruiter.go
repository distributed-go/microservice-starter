package recruiterservice

import (
	"errors"
	"net/http"
	"time"

	"github.com/jobbox-tech/recruiter-api/dal/recruiterdal"
	"github.com/jobbox-tech/recruiter-api/models/recruitermodel"

	"github.com/go-chi/render"
	"github.com/jobbox-tech/recruiter-api/database/dbmodels"
	"github.com/jobbox-tech/recruiter-api/web/renderers"

	"github.com/jobbox-tech/recruiter-api/logging"
)

type recruiterservice struct {
	logger       logging.Logger
	recruiterDal recruiterdal.RecruiterDal
}

// NewRecruiterService returns service impl
func NewRecruiterService() RecruiterService {
	return &recruiterservice{
		logger:       logging.NewLogger(),
		recruiterDal: recruiterdal.NewRecruiterDal(),
	}
}

func (rs *recruiterservice) CreateRecruiter(w http.ResponseWriter, r *http.Request) {
	txID := r.Header["transaction_id"][0]
	data := &recruitersRequest{}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, renderers.ErrorInvalidRequest(err))
		return
	}

	data.Recruiter.Roles = []recruitermodel.Role{recruitermodel.USER}
	data.Recruiter.Active = true
	data.Recruiter.CreatedTimestampUTC = time.Now().UTC()
	data.Recruiter.UpdatedTimestampUTC = time.Now().UTC()

	_, err := rs.recruiterDal.Create(txID, data.Recruiter)
	if err != nil {
		rs.logger.Error(txID, "").Errorf("Failed to create recruiters record with error %v", err)
		render.Render(w, r, renderers.ErrorInternalServerError(errors.New("Failed to create recruiter account, please try again")))
		return
	}

	// data.Recruiter.ID = objectID
	render.Respond(w, r, newRecruitersResponse(data.Recruiter))
}

// ==============  Bindings  ===============
type recruitersRequest struct {
	*dbmodels.Recruiter
}

func (d *recruitersRequest) Bind(r *http.Request) error {
	return nil
}

type recruitersResponse struct {
	*dbmodels.Recruiter
}

func newRecruitersResponse(a *dbmodels.Recruiter) *recruitersResponse {
	resp := &recruitersResponse{Recruiter: a}
	return resp
}
