package jobsservice

import (
	"net/http"
	"time"

	"github.com/go-chi/render"
	"github.com/jobbox-tech/recruiter-api/web/interfaces/v1/jobsinterface"
	"github.com/jobbox-tech/recruiter-api/web/renderers"
)

// @Summary Post a new Job
// @Description It allows to Post a new Job
// @Tags jobs
// @Accept json
// @Produce json
// @Param * body jobsinterface.JobRequest{} true "Job Details"
// @Success 200 {object} jobsinterface.JobResponse{}
// @Failure 400 {object} errorinterface.ErrorResponse{}
// @Failure 401 {object} errorinterface.ErrorResponse{}
// @Failure 500 {object} errorinterface.ErrorResponse{}
// @Router /jobs [POST]
func (j *jobsservice) Create(w http.ResponseWriter, r *http.Request) {
	txID := r.Header["transaction_id"][0]
	claims := j.middlewares.ClaimsFromCtx(r.Context())

	job := &jobsinterface.JobRequest{}
	if err := render.Bind(r, job); err != nil {
		j.logger.Error(txID, FailedToCreateJob).Error(err)
		render.Render(w, r, renderers.ErrorBadRequest(ErrIncompleteDetails))
		return
	}

	acc, err := j.recruiterDal.GetByEmail(claims.Sub)
	if err != nil {
		j.logger.Error(txID, FailedToCreateJob).Error(err)
		render.Render(w, r, renderers.ErrorInternalServerError(err))
		return
	}

	job.Job.CreatedTimestampUTC = time.Now().UTC()
	job.Job.UpdatedTimestampUTC = time.Now().UTC()
	job.Job.RecruiterID = acc.ID
	job.Job.OrganizationID = acc.OrganizationID

	resp, err := j.jobsDal.Create(txID, &job.Job)
	if err != nil {
		j.logger.Error(txID, FailedToCreateJob).Error(err)
		render.Render(w, r, renderers.ErrorInternalServerError(err))
		return
	}

	render.Respond(w, r, &jobsinterface.JobResponse{
		Job: resp,
	})
	return
}
