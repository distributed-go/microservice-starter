package jwt

import (
	"net/http"
	"time"

	"github.com/distributed-go/microservice-starter/models/authmodel"
	"github.com/go-chi/jwtauth"
	"github.com/spf13/viper"
)

type tokenAuth struct {
	JwtAuth          *jwtauth.JWTAuth
	JwtExpiry        time.Duration
	JwtRefreshExpiry time.Duration
}

// NewTokenAuth configures and returns a JWT authentication instance.
func NewTokenAuth() TokenAuth {
	secret := viper.GetString("jwt.auth_jwt_secret")
	a := &tokenAuth{
		JwtAuth:          jwtauth.New("HS256", []byte(secret), nil),
		JwtExpiry:        viper.GetDuration("jwt.auth_jwt_expiry"),
		JwtRefreshExpiry: viper.GetDuration("jwt.auth_jwt_refresh_expiry"),
	}

	return a
}

// Verifier http middleware will verify a jwt string from a http request.
func (a *tokenAuth) Verifier() func(http.Handler) http.Handler {
	return jwtauth.Verifier(a.JwtAuth)
}

// GenTokenPair returns both an access token and a refresh token.
func (a *tokenAuth) GenTokenPair(accessClaims authmodel.AppClaims, refreshClaims authmodel.RefreshClaims) (string, string, error) {
	access, err := a.CreateJWT(accessClaims)
	if err != nil {
		return "", "", err
	}
	refresh, err := a.CreateRefreshJWT(refreshClaims)
	if err != nil {
		return "", "", err
	}
	return access, refresh, nil
}

// CreateJWT returns an access token for provided account claims.
func (a *tokenAuth) CreateJWT(c authmodel.AppClaims) (string, error) {
	c.IssuedAt = time.Now().Unix()
	c.ExpiresAt = time.Now().Add(a.JwtExpiry).Unix()
	_, tokenString, err := a.JwtAuth.Encode(c)
	return tokenString, err
}

// CreateRefreshJWT returns a refresh token for provided token Claims.
func (a *tokenAuth) CreateRefreshJWT(c authmodel.RefreshClaims) (string, error) {
	c.IssuedAt = time.Now().Unix()
	c.ExpiresAt = time.Now().Add(a.JwtExpiry).Unix()
	_, tokenString, err := a.JwtAuth.Encode(c)
	return tokenString, err
}
