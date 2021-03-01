package router

import (
	"github.com/go-chi/cors"
	"github.com/spf13/viper"
)

func corsConfig() *cors.Cors {
	// Basic CORS
	// for more ideas, see: https://developer.github.com/v3/#cross-origin-resource-sharing
	return cors.New(cors.Options{
		// AllowedOrigins: []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: viper.GetStringSlice("web.allowed_origins"),
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   viper.GetStringSlice("web.allowed_methods"),
		AllowedHeaders:   viper.GetStringSlice("web.allowed_headers"),
		ExposedHeaders:   viper.GetStringSlice("web.exposed_headers"),
		AllowCredentials: viper.GetBool("web.allow_credentials"),
		MaxAge:           viper.GetInt("web.max_age"), // Maximum value not ignored by any of major browsers
	})
}
