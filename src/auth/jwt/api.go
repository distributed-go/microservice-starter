package jwt

import (
	"net/http"

	"github.com/distributed-go/microservice-starter/models/authmodel"
)

// TokenAuth interface
type TokenAuth interface {
	Verifier() func(http.Handler) http.Handler
	GenTokenPair(accessClaims authmodel.AppClaims, refreshClaims authmodel.RefreshClaims) (string, string, error)
	CreateJWT(c authmodel.AppClaims) (string, error)
	CreateRefreshJWT(c authmodel.RefreshClaims) (string, error)
}
