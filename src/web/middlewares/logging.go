package middlewares

import (
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/distributed-go/microservice-starter/logging"
	"github.com/go-chi/chi/middleware"
	"github.com/sirupsen/logrus"
)

type logger struct {
	logger logging.Logger
}

// NewMiddlewares returns new logging middleware for http
func NewMiddlewares() Middlewares {
	return &logger{
		logger: logging.NewLogger(),
	}
}

// Logger returns a request logging middleware
func (l *logger) Logger() func(h http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			reqID := middleware.GetReqID(r.Context())
			r.Header["transaction_id"] = []string{reqID}
			ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)
			t1 := time.Now()
			defer func() {
				remoteIP, _, err := net.SplitHostPort(r.RemoteAddr)
				if err != nil {
					remoteIP = r.RemoteAddr
				}
				scheme := "http"
				if r.TLS != nil {
					scheme = "https"
				}
				fields := logrus.Fields{
					"ts_started_utc":    t1.UTC(),
					"ts_finished_utc":   time.Now().UTC(),
					"status_code":       ww.Status(),
					"bytes_transferred": ww.BytesWritten(),
					"latency_ms":        time.Since(t1).Nanoseconds() / 1000000.0,
					"remote_ip":         remoteIP,
					"proto":             r.Proto,
					"method":            r.Method,
					"user_agent":        r.UserAgent(),
					"uri":               fmt.Sprintf("%s://%s%s", scheme, r.Host, r.RequestURI),
					"transaction_id":    reqID,
				}
				l.logger.WithFields(fields).Info("")
			}()

			h.ServeHTTP(ww, r)
		}

		return http.HandlerFunc(fn)
	}
}
