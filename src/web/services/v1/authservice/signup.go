package authservice

import (
	"net/http"
	"time"

	"github.com/distributed-go/microservice-starter/database/dbmodels"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/distributed-go/microservice-starter/models/authmodel"
	"github.com/distributed-go/microservice-starter/models/recruitermodel"
	"github.com/distributed-go/microservice-starter/web/interfaces/v1/authinterface"
	"github.com/distributed-go/microservice-starter/web/renderers"
	"github.com/go-chi/render"
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

	var orgID primitive.ObjectID
	var organization *dbmodels.Organization
	var newOrg bool
	if len(data.Company.CompanyID) > 0 {
		organizationID, err := primitive.ObjectIDFromHex(data.Company.CompanyID)
		if err != nil {
			as.logger.Error(txID, authmodel.FailedToSignUp).Errorf("Failed to parse ID with error %v", err)
			render.Render(w, r, renderers.ErrorBadRequest(authmodel.ErrIncorrectDetails))
			return
		}

		_, err = as.orgDal.GetByID(organizationID)
		if err != nil {
			if err != nil {
				as.logger.Error(txID, authmodel.FailedToSignUp).Errorf("Failed to get organization record with error %v", err)
				render.Render(w, r, renderers.ErrorBadRequest(authmodel.ErrIncorrectDetails))
				return
			}
		}
		orgID = organizationID
	} else {
		org, err := as.orgDal.Create(txID, &dbmodels.Organization{
			OrganizationName: data.Company.CompanyName,
			IsActive:         true,
		})
		if err != nil {
			as.logger.Error(txID, authmodel.FailedToSignUp).Errorf("Failed to create organization record with error %v", err)
			render.Render(w, r, renderers.ErrorBadRequest(authmodel.ErrServerError))
			return
		}
		orgID = org.ID
		organization = org
		newOrg = true
	}

	recruiter := &dbmodels.Recruiter{}
	recruiter.Roles = []recruitermodel.Role{recruitermodel.USER}
	recruiter.Active = true
	recruiter.CreatedTimestampUTC = time.Now().UTC()
	recruiter.UpdatedTimestampUTC = time.Now().UTC()
	recruiter.Email = data.Email
	recruiter.FirstName = data.FirstName
	recruiter.Designation = data.Designation
	recruiter.OrganizationID = orgID

	acc, err := as.recruiterDal.Create(txID, recruiter)
	if err != nil {
		as.logger.Error(txID, authmodel.FailedToSignUp).Errorf("Failed to create recruiter record with error %v", err)
		render.Render(w, r, renderers.ErrorInternalServerError(authmodel.ErrServerError))
		return
	}

	if newOrg {
		organization.CreatedBy = recruiter.ID
		organization.Admins = []primitive.ObjectID{recruiter.ID}
		err = as.orgDal.Update(organization)
		if err != nil {
			as.logger.Error(txID, authmodel.FailedToSignUp).Errorf("Failed to update organization record with error %v", err)
			render.Render(w, r, renderers.ErrorInternalServerError(authmodel.ErrServerError))
			return
		}
	}

	err = as.loginWithAccount(acc, txID, r)
	if err != nil {
		render.Render(w, r, renderers.ErrorInternalServerError(authmodel.ErrServerError))
		return
	}

	render.Respond(w, r, http.NoBody)
}
