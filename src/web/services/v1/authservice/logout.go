package authservice

import (
	"net/http"

	"github.com/jobbox-tech/recruiter-api/models/authmodel"
	"github.com/jobbox-tech/recruiter-api/web/renderers"

	"github.com/go-chi/jwtauth"
	"github.com/go-chi/render"
)

// @Summary Logout
// @Description It allows to logout users from account with JWT
// @Tags authentication
// @Param Authorization header string true "BEARER JWT"
// @Accept json
// @Produce json
// @Success 200
// @Failure 401 {object} errorinterface.ErrorResponse{}
// @Failure 404 {object} errorinterface.ErrorResponse{}
// @Router /logout [POST]
func (as *authservice) Logout(w http.ResponseWriter, r *http.Request) {
	txID := r.Header["transaction_id"][0]

	token, _, err := jwtauth.FromContext(r.Context())
	if err != nil {
		render.Render(w, r, renderers.ErrorUnauthorized(authmodel.ErrLoginToken))
		return
	}

	err = as.tokenDal.DeleteByAccessToken(token.Raw)
	if err != nil {
		as.logger.Error(authmodel.FailedToDeleteToken, txID).Errorf("Failed to delete access token with error %v", err)
		render.Respond(w, r, http.NoBody)
	}

	render.Respond(w, r, http.NoBody)
}
