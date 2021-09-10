package authservice

import (
	"fmt"
	"net/http"
	"time"

	"github.com/distributed-go/microservice-starter/database/dbmodels"
	"github.com/distributed-go/microservice-starter/email/authemail"
	"github.com/distributed-go/microservice-starter/email/mailer"
	"github.com/distributed-go/microservice-starter/models/authmodel"
	"github.com/distributed-go/microservice-starter/web/interfaces/v1/authinterface"
	"github.com/distributed-go/microservice-starter/web/renderers"
	"github.com/go-chi/render"
	"github.com/google/uuid"
	"github.com/mssola/user_agent"
	"github.com/spf13/viper"
)

// @Summary Login to account with email
// @Description It allows to login to account with email address registered
// @Tags authentication
// @Accept json
// @Produce json
// @Param * body authinterface.LoginReqInterface{} true "login with email"
// @Success 200
// @Failure 400 {object} errorinterface.ErrorResponse{}
// @Failure 404 {object} errorinterface.ErrorResponse{}
// @Failure 500 {object} errorinterface.ErrorResponse{}
// @Router /login [POST]
func (as *authservice) Login(w http.ResponseWriter, r *http.Request) {
	txID := r.Header["transaction_id"][0]
	body := &authinterface.LoginReqInterface{}
	if err := render.Bind(r, body); err != nil {
		as.logger.Error(txID, authmodel.FailedToCreateAccessToken).Error(err)
		render.Render(w, r, renderers.ErrorBadRequest(authmodel.ErrInvalidLogin))
		return
	}

	acc, err := as.recruiterDal.GetByEmail(body.Email)
	if err != nil {
		as.logger.Error(txID, authmodel.FailedToCreateAccessToken).Error(err)
		render.Render(w, r, renderers.ErrorBadRequest(authmodel.ErrUnknownLogin))
		return
	}

	if !acc.CanLogin() {
		render.Render(w, r, renderers.ErrorBadRequest(authmodel.ErrLoginDisabled))
		return
	}

	err = as.loginWithAccount(acc, txID, r)
	if err != nil {
		render.Render(w, r, renderers.ErrorInternalServerError(authmodel.ErrServerError))
	}

	render.Respond(w, r, http.NoBody)
}

func (as *authservice) loginWithAccount(acc *dbmodels.Recruiter, txID string, r *http.Request) error {
	token := uuid.New().String()
	ua := user_agent.New(r.UserAgent())
	browser, _ := ua.Browser()
	accessToken := &dbmodels.Token{
		CreatedTimestampUTC: time.Now().UTC(),
		UpdatedTimestampUTC: time.Now().UTC(),
		AccountID:           acc.ID,
		ExpiryTimestampUTC:  time.Now().UTC().Add(viper.GetDuration("jwt.auth_login_token_expiry")),
		TokenUUID:           token,
		UserAgent:           fmt.Sprintf("%s on %s", browser, ua.OS()),
		Mobile:              ua.Mobile(),
	}

	_, err := as.tokenDal.Create(txID, accessToken)
	if err != nil {
		as.logger.Error(txID, authmodel.FailedToCreateAccessToken).Errorf("Failed to create access token with error %v", err)
		return err
	}

	go func() {
		content := authemail.LoginEmail{
			Email:  acc.Email,
			Name:   acc.FirstName,
			Token:  token,
			Expiry: time.Now().Add(viper.GetDuration("jwt.auth_login_token_expiry")),
			URL:    fmt.Sprintf("%s%s/%s", viper.GetString("website.domain_name"), viper.GetString("website.auth_login_url"), token),
		}
		if err := as.authemail.SendLoginEmail(mailer.Recipient{Name: acc.FirstName, Address: acc.Email}, content); err != nil {
			as.logger.Error(txID, authmodel.FailedToCreateAccessToken).Errorf("Failed to send login link with error %v", err)
		}
	}()

	return nil
}
