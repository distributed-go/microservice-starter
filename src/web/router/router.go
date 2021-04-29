// Package api configures an http server for administration and application resources.

package router

import (
	"time"

	"github.com/jobbox-tech/recruiter-api/auth/jwt"
	"github.com/jobbox-tech/recruiter-api/web/services/v1/authservice"

	"github.com/jobbox-tech/recruiter-api/web/services/v1/recruiterservice"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/jobbox-tech/recruiter-api/logging"
	"github.com/jobbox-tech/recruiter-api/web/middlewares"
	"github.com/jobbox-tech/recruiter-api/web/services/health"
	"github.com/spf13/viper"
)

type router struct {
	logger      logging.Logger
	health      health.Health
	recruiter   recruiterservice.RecruiterService
	auth        authservice.AuthService
	tokenAuth   jwt.TokenAuth
	middlewares middlewares.Middlewares
}

// NewRouter returns the router implementation
func NewRouter() Router {
	return &router{
		logger:      logging.NewLogger(),
		health:      health.NewHealth(),
		recruiter:   recruiterservice.NewRecruiterService(),
		auth:        authservice.NewAuthService(),
		tokenAuth:   jwt.NewTokenAuth(),
		middlewares: middlewares.NewMiddlewares(),
	}
}

// Router configures application resources and routes.
func (router *router) Router(enableCORS bool) *chi.Mux {
	// v1 URL router prefix
	v1Prefix := viper.GetString("web.url_prefix") + viper.GetString("web.api_version_v1")

	// ==================== Public Router ======================
	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.DefaultCompress)
	r.Use(middleware.Timeout(time.Duration(viper.GetInt("web.request_timeout_in_sec")) * time.Second))

	// set up logging
	r.Use(router.middlewares.Logger())
	// settin up content-type
	r.Use(render.SetContentType(render.ContentTypeJSON))
	// use CORS middleware if client is not served by this api, e.g. from other domain or CDN
	if enableCORS {
		r.Use(corsConfig().Handler)
	}

	// ==================== Private Router ========================
	rprivate := chi.NewRouter()
	rprivate.Use(router.tokenAuth.Verifier())
	rprivate.Use(router.middlewares.Authenticator)
	r.Mount("/", rprivate)

	// =================  health routes ======================
	r.Get(viper.GetString("web.url_prefix")+"/health", router.health.GetHealth)

	// =================  auth routes ======================
	r.Post(v1Prefix+"/signup", router.auth.SignUp)
	r.Post(v1Prefix+"/login", router.auth.Login)
	r.Post(v1Prefix+"/authenticate", router.auth.Authenticate)
	rprivate.Post(v1Prefix+"/logout", router.auth.Logout)
	rprivate.Post(v1Prefix+"/validate", router.auth.Validate)

	return r
}
