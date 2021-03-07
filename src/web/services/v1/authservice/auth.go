package authservice

import (
	"time"

	"github.com/jobbox-tech/recruiter-api/dal/recruiterdal"
	"github.com/jobbox-tech/recruiter-api/dal/tokendal"
	"github.com/jobbox-tech/recruiter-api/email/authemail"
	"github.com/jobbox-tech/recruiter-api/logging"
	"github.com/spf13/viper"
)

type authservice struct {
	logger           logging.Logger
	loginURL         string
	loginTokenExpiry time.Duration

	tokenDal     tokendal.TokenDal
	recruiterDal recruiterdal.RecruiterDal
	authemail    authemail.AuthEmail
}

// NewAuthService returns service impl
func NewAuthService() AuthService {
	return &authservice{
		logger:           logging.NewLogger(),
		loginTokenExpiry: viper.GetDuration("jwt.auth_login_token_expiry"),

		tokenDal:     tokendal.NewTokenDal(),
		recruiterDal: recruiterdal.NewRecruiterDal(),
		authemail:    authemail.NewAuthEmail(),
	}
}
