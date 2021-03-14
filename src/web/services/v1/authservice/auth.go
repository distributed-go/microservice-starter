package authservice

import (
	"github.com/jobbox-tech/recruiter-api/auth/jwt"
	"github.com/jobbox-tech/recruiter-api/dal/recruiterdal"
	"github.com/jobbox-tech/recruiter-api/dal/tokendal"
	"github.com/jobbox-tech/recruiter-api/email/authemail"
	"github.com/jobbox-tech/recruiter-api/logging"
)

type authservice struct {
	logger   logging.Logger
	loginURL string

	tokenDal     tokendal.TokenDal
	recruiterDal recruiterdal.RecruiterDal

	tokenAuth jwt.TokenAuth
	authemail authemail.AuthEmail
}

// NewAuthService returns service impl
func NewAuthService() AuthService {
	return &authservice{
		logger: logging.NewLogger(),

		tokenDal:     tokendal.NewTokenDal(),
		recruiterDal: recruiterdal.NewRecruiterDal(),

		authemail: authemail.NewAuthEmail(),
		tokenAuth: jwt.NewTokenAuth(),
	}
}
