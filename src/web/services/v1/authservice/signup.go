package authservice

import (
	"net/http"
	"time"

	"github.com/jobbox-tech/recruiter-api/database/dbmodels"

	"github.com/go-chi/render"
	"github.com/jobbox-tech/recruiter-api/models/authmodel"
	"github.com/jobbox-tech/recruiter-api/models/recruitermodel"
	"github.com/jobbox-tech/recruiter-api/web/interfaces/v1/authinterface"
	"github.com/jobbox-tech/recruiter-api/web/renderers"
)

// @Summary Sign up with email
// @Description It allows to sign up with email address and personal details
// @Tags authentication
// @Accept json
// @Produce json
// @Param * body authinterface.SignUpReqInterface{} true "signup with email"
// @Success 200
// @Failure 400 {object} errorinterface.ErrorResponse{}
// @Failure 404 {object} errorinterface.ErrorResponse{}
// @Failure 500 {object} errorinterface.ErrorResponse{}
// @Router /signup [POST]
func (as *authservice) SignUp(w http.ResponseWriter, r *http.Request) {
	txID := r.Header["transaction_id"][0]
	data := &authinterface.SignUpReqInterface{}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, renderers.ErrorBadRequest(authmodel.ErrIncorrectDetails))
		return
	}

	account, _ := as.recruiterDal.GetByEmail(data.Email)
	if account != nil {
		render.Render(w, r, renderers.ErrorBadRequest(authmodel.ErrAlreadyRegistered))
		return
	}

	recruiter := &dbmodels.Recruiter{}
	recruiter.Roles = []recruitermodel.Role{recruitermodel.USER}
	recruiter.Active = true
	recruiter.CreatedTimestampUTC = time.Now().UTC()
	recruiter.UpdatedTimestampUTC = time.Now().UTC()
	recruiter.Email = data.Email
	recruiter.FirstName = data.FirstName
	recruiter.Designation = data.Designation

	acc, err := as.recruiterDal.Create(txID, recruiter)
	if err != nil {
		as.logger.Error(txID, authmodel.FailedToSignUp).Errorf("Failed to create recruiter record with error %v", err)
		render.Render(w, r, renderers.ErrorInternalServerError(authmodel.ErrServerError))
		return
	}

	err = as.loginWithAccount(acc, txID, r)
	if err != nil {
		render.Render(w, r, renderers.ErrorInternalServerError(authmodel.ErrServerError))
		return
	}

	render.Respond(w, r, http.NoBody)
}
