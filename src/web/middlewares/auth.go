package middlewares

import (
	"context"
	"net/http"

	"github.com/go-chi/jwtauth"
	"github.com/go-chi/render"
	"github.com/jobbox-tech/recruiter-api/models/authmodel"
	"github.com/jobbox-tech/recruiter-api/models/recruitermodel"
	"github.com/jobbox-tech/recruiter-api/web/renderers"
)

type ctxKey int

const (
	ctxClaims ctxKey = iota
	ctxRefreshToken
)

// ClaimsFromCtx retrieves the parsed AppClaims from request context.
func (l *logger) ClaimsFromCtx(ctx context.Context) authmodel.AppClaims {
	return ctx.Value(ctxClaims).(authmodel.AppClaims)
}

// RefreshTokenFromCtx retrieves the parsed refresh token from context.
func (l *logger) RefreshTokenFromCtx(ctx context.Context) string {
	return ctx.Value(ctxRefreshToken).(string)
}

// Authenticator is a default authentication middleware to enforce access from the
// Verifier middleware request context values. The Authenticator sends a 401 Unauthorized
// response for any unverified tokens and passes the good ones through.
func (l *logger) Authenticator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, claims, err := jwtauth.FromContext(r.Context())
		if err != nil {
			render.Render(w, r, renderers.ErrorUnauthorized(authmodel.ErrInvalidLogin))
			return
		}

		if !token.Valid {
			render.Render(w, r, renderers.ErrorUnauthorized(authmodel.ErrLoginToken))
			return
		}

		// Token is authenticated, parse claims
		var c authmodel.AppClaims
		err = c.ParseClaims(claims)
		if err != nil {
			render.Render(w, r, renderers.ErrorUnauthorized(authmodel.ErrLoginToken))
			return
		}

		// Set AppClaims on context
		ctx := context.WithValue(r.Context(), ctxClaims, c)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// AuthenticateRefreshJWT checks validity of refresh tokens and is only used for access token refresh and logout requests. It responds with 401 Unauthorized for invalid or expired refresh tokens.
func (l *logger) AuthenticateRefreshJWT(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, claims, err := jwtauth.FromContext(r.Context())
		if err != nil {
			// logging.GetLogEntry(r).Warn(err)
			render.Render(w, r, renderers.ErrorUnauthorized(authmodel.ErrInvalidLogin))
			return
		}
		if !token.Valid {
			render.Render(w, r, renderers.ErrorUnauthorized(authmodel.ErrLoginToken))
			return
		}

		// Token is authenticated, parse refresh token string
		var c authmodel.RefreshClaims
		err = c.ParseClaims(claims)
		if err != nil {
			// logging.GetLogEntry(r).Error(err)
			render.Render(w, r, renderers.ErrorUnauthorized(authmodel.ErrInvalidRefreshToken))
			return
		}
		// Set refresh token string on context
		ctx := context.WithValue(r.Context(), ctxRefreshToken, c.TokenUUID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// RequiresRole middleware restricts access to accounts having role parameter in their jwt claims.
func (l *logger) RequiresRole(role recruitermodel.Role) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		hfn := func(w http.ResponseWriter, r *http.Request) {
			claims := l.ClaimsFromCtx(r.Context())
			if !hasRole(role, claims.Roles) {
				render.Render(w, r, renderers.ErrorForbidden(authmodel.ErrInsufficientRights))
				return
			}
			next.ServeHTTP(w, r)
		}
		return http.HandlerFunc(hfn)
	}
}

func hasRole(role recruitermodel.Role, roles []recruitermodel.Role) bool {
	for _, r := range roles {
		if r == role {
			return true
		}
	}
	return false
}
