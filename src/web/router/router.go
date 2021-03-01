// Package api configures an http server for administration and application resources.

package router

import (
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/jobbox-tech/recruiter-api/logging"
	"github.com/jobbox-tech/recruiter-api/web/middlewares"
	"github.com/jobbox-tech/recruiter-api/web/services/health"
	"github.com/spf13/viper"
)

type router struct {
	logger logging.Logger
	health health.Health
}

// NewRouter returns the router implementation
func NewRouter() Router {
	return &router{
		logger: logging.NewLogger(),
		health: health.NewHealth(),
	}
}

// Router configures application resources and routes.
func (router *router) Router(enableCORS bool) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.DefaultCompress)
	r.Use(middleware.Timeout(time.Duration(viper.GetInt("web.request_timeout_in_sec")) * time.Second))

	// set up logging
	r.Use(middlewares.NewLoggingMiddleware().Logger())

	// settin up content-type
	r.Use(render.SetContentType(render.ContentTypeJSON))

	// use CORS middleware if client is not served by this api, e.g. from other domain or CDN
	if enableCORS {
		r.Use(corsConfig().Handler)
	}

	// v1 URL router prefix
	v1Prefix := viper.GetString("web.url_prefix") + viper.GetString("web.api_version_v1")

	// =================  health routes ======================
	r.Get(viper.GetString("web.url_prefix")+"/health", router.health.GetHealth)

	// =================  ping pong ======================
	r.Get(v1Prefix+"/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	return r
}
