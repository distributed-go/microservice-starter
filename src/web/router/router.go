// Package api configures an http server for administration and application resources.

package router

import (
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/jobbox-tech/recruiter-api/logging"
	"github.com/spf13/viper"
)

// NewRouter configures application resources and routes.
func NewRouter(enableCORS bool) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.DefaultCompress)
	r.Use(middleware.Timeout(15 * time.Second))

	r.Use(logging.NewHTTPLogger())
	r.Use(render.SetContentType(render.ContentTypeJSON))

	// use CORS middleware if client is not served by this api, e.g. from other domain or CDN
	if enableCORS {
		r.Use(corsConfig().Handler)
	}

	// v1 URL router prefix
	v1Prefix := viper.GetString("web.url_prefix") + viper.GetString("web.api_version_v1")

	r.Get(v1Prefix+"/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	return r
}
