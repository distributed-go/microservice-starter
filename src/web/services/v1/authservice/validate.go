package authservice

import (
	"net/http"

	"github.com/distributed-go/microservice-starter/models/authmodel"
	"github.com/distributed-go/microservice-starter/web/renderers"
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
// @Router /validate [POST]
func (as *authservice) Validate(w http.ResponseWriter, r *http.Request) {
	_, _, err := jwtauth.FromContext(r.Context())
	if err != nil {
		render.Render(w, r, renderers.ErrorUnauthorized(authmodel.ErrLoginToken))
		return
	}

	render.Respond(w, r, http.NoBody)
}
