package jobsservice

import (
	"errors"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/jobbox-tech/recruiter-api/web/interfaces/v1/jobsinterface"
	"github.com/jobbox-tech/recruiter-api/web/renderers"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// @Summary Get a Job
// @Description It allows to reterive a Job by ID
// @Tags jobs
// @Accept json
// @Produce json
// @Param jobID path string true "Job ID"
// @Success 200 {object} jobsinterface.JobResponse{}
// @Failure 400 {object} errorinterface.ErrorResponse{}
// @Failure 404 {object} errorinterface.ErrorResponse{}
// @Failure 500 {object} errorinterface.ErrorResponse{}
// @Router /jobs/jobID [GET]
func (j *jobsservice) Get(w http.ResponseWriter, r *http.Request) {
	txID := r.Header["transaction_id"][0]
	job := chi.URLParam(r, "jobID")

	jobID, err := primitive.ObjectIDFromHex(job)
	if err != nil {
		j.logger.Error(txID, FailedToGetJob).Error(err)
		render.Render(w, r, renderers.ErrorNotFound(err))
		return
	}

	resp, err := j.jobsDal.GetByID(jobID)
	if err != nil {
		j.logger.Error(txID, FailedToGetJob).Error(err)
		render.Render(w, r, renderers.ErrorNotFound(err))
		return
	}

	if resp.Deleted {
		render.Render(w, r, renderers.ErrorNotFound(errors.New("This job does not exists")))
		return
	}

	render.Respond(w, r, &jobsinterface.JobResponse{
		Job: resp,
	})
	return
}
