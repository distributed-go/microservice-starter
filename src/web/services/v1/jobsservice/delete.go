package jobsservice

import (
	"errors"
	"net/http"
	"time"

	"github.com/distributed-go/microservice-starter/web/renderers"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// @Summary Delete a Job
// @Description It allows to delete a Job by ID
// @Tags jobs
// @Accept json
// @Produce json
// @Param jobID path string true "Job ID"
// @Param Authorization header string true "BEARER JWT"
// @Success 200
// @Failure 400 {object} errorinterface.ErrorResponse{}
// @Failure 401 {object} errorinterface.ErrorResponse{}
// @Failure 404 {object} errorinterface.ErrorResponse{}
// @Failure 500 {object} errorinterface.ErrorResponse{}
// @Router /jobs/jobID [DELETE]
func (j *jobsservice) Delete(w http.ResponseWriter, r *http.Request) {
	txID := r.Header["transaction_id"][0]
	jobIDstring := chi.URLParam(r, "jobID")
	claims := j.middlewares.ClaimsFromCtx(r.Context())

	jobID, err := primitive.ObjectIDFromHex(jobIDstring)
	if err != nil {
		j.logger.Error(txID, FailedToDeleteJob).Error(err)
		render.Render(w, r, renderers.ErrorNotFound(err))
		return
	}

	resp, err := j.jobsDal.GetByID(jobID)
	if err != nil {
		j.logger.Error(txID, FailedToDeleteJob).Error(err)
		render.Render(w, r, renderers.ErrorNotFound(err))
		return
	}

	if resp.Deleted {
		render.Render(w, r, renderers.ErrorNotFound(errors.New("This job does not exists")))
		return
	}

	acc, err := j.recruiterDal.GetByEmail(claims.Sub)
	if err != nil {
		j.logger.Error(txID, FailedToDeleteJob).Error(err)
		render.Render(w, r, renderers.ErrorInternalServerError(err))
		return
	}

	if resp.RecruiterID != acc.ID {
		render.Render(w, r, renderers.ErrorForbidden(errors.New("You do not have permission to modify this job")))
		return
	}

	resp.UpdatedTimestampUTC = time.Now().UTC()
	resp.Deleted = true

	err = j.jobsDal.Update(resp)
	if err != nil {
		j.logger.Error(txID, FailedToUpdateJob).Error(err)
		render.Render(w, r, renderers.ErrorInternalServerError(err))
		return
	}

	render.Respond(w, r, http.NoBody)
	return
}
