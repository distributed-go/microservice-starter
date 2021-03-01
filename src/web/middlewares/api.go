package middlewares

import "net/http"

// LoggingMiddleware interfaces
type LoggingMiddleware interface {
	Logger() func(h http.Handler) http.Handler
}
