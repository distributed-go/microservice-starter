package jobsservice

import (
	"errors"
	"net/http"
	"time"

	"github.com/distributed-go/microservice-starter/web/interfaces/v1/jobsinterface"
	"github.com/distributed-go/microservice-starter/web/renderers"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// @Summary Update a  Job
// @Description It allows to update a Job
// @Tags jobs
// @Accept json
// @Produce json
// @Param * body jobsinterface.JobRequest{} true "Job Details"
// @Param Authorization header string true "BEARER JWT"
// @Param jobID path string true "Job ID"
// @Success 200 {object} jobsinterface.JobResponse{}
// @Failure 400 {object} errorinterface.ErrorResponse{}
// @Failure 401 {object} errorinterface.ErrorResponse{}
// @Failure 500 {object} errorinterface.ErrorResponse{}
// @Router /jobs/jobID [PUT]
func (j *jobsservice) Update(w http.ResponseWriter, r *http.Request) {
	txID := r.Header["transaction_id"][0]
	jobIDstring := chi.URLParam(r, "jobID")
	claims := j.middlewares.ClaimsFromCtx(r.Context())

	jobID, err := primitive.ObjectIDFromHex(jobIDstring)
	if err != nil {
		j.logger.Error(txID, FailedToUpdateJob).Error(err)
		render.Render(w, r, renderers.ErrorNotFound(err))
		return
	}

	job := &jobsinterface.JobRequest{}
	if err := render.Bind(r, job); err != nil {
		j.logger.Error(txID, FailedToUpdateJob).Error(err)
		render.Render(w, r, renderers.ErrorBadRequest(ErrIncompleteDetails))
		return
	}

	resp, err := j.jobsDal.GetByID(jobID)
	if err != nil {
		j.logger.Error(txID, FailedToUpdateJob).Error(err)
		render.Render(w, r, renderers.ErrorNotFound(err))
		return
	}

	acc, err := j.recruiterDal.GetByEmail(claims.Sub)
	if err != nil {
		j.logger.Error(txID, FailedToUpdateJob).Error(err)
		render.Render(w, r, renderers.ErrorInternalServerError(err))
		return
	}

	if resp.RecruiterID != acc.ID {
		render.Render(w, r, renderers.ErrorForbidden(errors.New("You do not have permission to modify this job")))
		return
	}

	job.Job.UpdatedTimestampUTC = time.Now().UTC()
	job.Job.RecruiterID = acc.ID
	job.Job.OrganizationID = acc.OrganizationID
	job.Job.ID = resp.ID

	err = j.jobsDal.Update(&job.Job)
	if err != nil {
		j.logger.Error(txID, FailedToUpdateJob).Error(err)
		render.Render(w, r, renderers.ErrorInternalServerError(err))
		return
	}

	render.Respond(w, r, &jobsinterface.JobResponse{
		Job: &job.Job,
	})
	return
}
