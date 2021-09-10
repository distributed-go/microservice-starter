package authservice

import (
	"github.com/distributed-go/microservice-starter/auth/jwt"
	"github.com/distributed-go/microservice-starter/dal/organizationdal"
	"github.com/distributed-go/microservice-starter/dal/recruiterdal"
	"github.com/distributed-go/microservice-starter/dal/tokendal"
	"github.com/distributed-go/microservice-starter/email/authemail"
	"github.com/distributed-go/microservice-starter/logging"
)

type authservice struct {
	logger   logging.Logger
	loginURL string

	tokenDal     tokendal.TokenDal
	recruiterDal recruiterdal.RecruiterDal

	tokenAuth jwt.TokenAuth
	authemail authemail.AuthEmail

	orgDal organizationdal.OrganizationDal
}

// NewAuthService returns service impl
func NewAuthService() AuthService {
	return &authservice{
		logger: logging.NewLogger(),

		tokenDal:     tokendal.NewTokenDal(),
		recruiterDal: recruiterdal.NewRecruiterDal(),

		authemail: authemail.NewAuthEmail(),
		tokenAuth: jwt.NewTokenAuth(),

		orgDal: organizationdal.NewOrganizationDal(),
	}
}
