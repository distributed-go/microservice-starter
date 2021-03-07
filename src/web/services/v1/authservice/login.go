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
	"github.com/jobbox-tech/recruiter-api/email/authemail"
	"github.com/jobbox-tech/recruiter-api/email/mailer"
	"github.com/jobbox-tech/recruiter-api/models/authmodel"
	"github.com/jobbox-tech/recruiter-api/web/renderers"
	"github.com/spf13/viper"
)

func (as *authservice) Login(w http.ResponseWriter, r *http.Request) {
	body := &loginRequest{}
	if err := render.Bind(r, body); err != nil {
		as.logger.WithField("email", body.Email).Error(err)
		render.Render(w, r, renderers.ErrorUnauthorized(authmodel.ErrInvalidLogin))
		return
	}

	acc, err := as.recruiterDal.GetAccountByEmail(body.Email)
	if err != nil {
		as.logger.WithField("email", body.Email).Error(err)
		render.Render(w, r, renderers.ErrorUnauthorized(authmodel.ErrInvalidLogin))
		return
	}

	if !acc.CanLogin() {
		render.Render(w, r, renderers.ErrorUnauthorized(authmodel.ErrLoginDisabled))
		return
	}

	go func() {
		token := uuid.New().String()
		content := authemail.LoginEmail{
			Email:  acc.Email,
			Name:   acc.FirstName,
			Token:  token,
			Expiry: time.Now().Add(as.loginTokenExpiry),
			URL:    fmt.Sprintf("%s%s/%s", viper.GetString("website.domain_name"), viper.GetString("website.auth_login_url"), token),
		}
		fmt.Println(content)
		if err := as.authemail.SendLoginEmail(mailer.Recipient{Name: acc.FirstName, Address: acc.Email}, content); err != nil {
			as.logger.WithField("email", body.Email).Error(err)
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
