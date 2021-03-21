package middlewares

import (
	"context"
	"net/http"

	"github.com/jobbox-tech/recruiter-api/models/authmodel"
	"github.com/jobbox-tech/recruiter-api/models/recruitermodel"
)

// Middlewares interfaces
type Middlewares interface {
	Logger() func(h http.Handler) http.Handler
	RequiresRole(role recruitermodel.Role) func(next http.Handler) http.Handler
	AuthenticateRefreshJWT(next http.Handler) http.Handler
	Authenticator(next http.Handler) http.Handler
	RefreshTokenFromCtx(ctx context.Context) string
	ClaimsFromCtx(ctx context.Context) authmodel.AppClaims
}
