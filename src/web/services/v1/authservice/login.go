package authservice

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/go-chi/render"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/google/uuid"
	"github.com/jobbox-tech/recruiter-api/database/dbmodels"
	"github.com/jobbox-tech/recruiter-api/email/authemail"
	"github.com/jobbox-tech/recruiter-api/email/mailer"
	"github.com/jobbox-tech/recruiter-api/models/authmodel"
	"github.com/jobbox-tech/recruiter-api/web/renderers"
	"github.com/mssola/user_agent"
	"github.com/spf13/viper"
)

func (as *authservice) Login(w http.ResponseWriter, r *http.Request) {
	txID := r.Header["transaction_id"][0]
	body := &loginRequest{}
	if err := render.Bind(r, body); err != nil {
		as.logger.Error(txID, authmodel.FailedToCreateAccessToken).Error(err)
		render.Render(w, r, renderers.ErrorUnauthorized(authmodel.ErrInvalidLogin))
		return
	}

	acc, err := as.recruiterDal.GetAccountByEmail(body.Email)
	if err != nil {
		as.logger.Error(txID, authmodel.FailedToCreateAccessToken).Error(err)
		render.Render(w, r, renderers.ErrorUnauthorized(authmodel.ErrInvalidLogin))
		return
	}

	if !acc.CanLogin() {
		render.Render(w, r, renderers.ErrorUnauthorized(authmodel.ErrLoginDisabled))
		return
	}

	identifier := uuid.New().String()
	ua := user_agent.New(r.UserAgent())
	browser, _ := ua.Browser()
	accessToken := &dbmodels.Token{
		CreatedTimestampUTC: time.Now().UTC(),
		UpdatedTimestampUTC: time.Now().UTC(),
		Token:               identifier, // initially token is set to identifier will be replaced with actual jwt on authentication
		AccountID:           acc.ID,
		ExpiryTimestampUTC:  time.Now().UTC().Add(viper.GetDuration("jwt.auth_login_token_expiry")),
		Identifier:          identifier,
		UserAgent:           browser,
	}

	_, err = as.tokenDal.Create(txID, accessToken)
	if err != nil {
		as.logger.Error(txID, authmodel.FailedToCreateAccessToken).Errorf("Failed to create access token with error %v", err)
		render.Render(w, r, renderers.ErrorInternalServerError(authmodel.ErrServerError.Error()))
		return
	}

	go func() {
		content := authemail.LoginEmail{
			Email:  acc.Email,
			Name:   acc.FirstName,
			Token:  identifier,
			Expiry: time.Now().Add(as.loginTokenExpiry),
			URL:    fmt.Sprintf("%s%s/%s", viper.GetString("website.domain_name"), viper.GetString("website.auth_login_url"), identifier),
		}
		if err := as.authemail.SendLoginEmail(mailer.Recipient{Name: acc.FirstName, Address: acc.Email}, content); err != nil {
			as.logger.Error(txID, authmodel.FailedToCreateAccessToken).Errorf("Failed to send login link with error %v", err)
		}
	}()

	render.Respond(w, r, http.NoBody)
}

// =================== bindings ========================= //
type loginRequest struct {
	Email string `json:"Email"`
}

func (body *loginRequest) Bind(r *http.Request) error {
	body.Email = strings.TrimSpace(body.Email)
	body.Email = strings.ToLower(body.Email)

	return validation.ValidateStruct(body,
		validation.Field(&body.Email, validation.Required, is.Email),
	)
}
