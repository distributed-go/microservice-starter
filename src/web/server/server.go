package server

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/jobbox-tech/recruiter-api/logging"
	"github.com/jobbox-tech/recruiter-api/web/router"
	"github.com/spf13/viper"
)

// Server provides an http.Server.
type server struct {
	svr               *http.Server
	logger            logging.Logger
	txID              string
	startTimeStampUTC time.Time
}

// NewServer creates and configures an APIServer serving all application routes.
func NewServer() Server {
	var addr string
	port := viper.GetString("host.port")
	txID := uuid.New().String()
	apiHandler := router.NewRouter().Router(viper.GetBool("host.enable_cors"))

	// allow port to be set as localhost:8001 in env during development to avoid "accept incoming network connection" request on restarts
	if strings.Contains(port, ":") {
		addr = port
	} else {
		addr = ":" + port
	}

	srv := http.Server{
		Addr:    addr,
		Handler: apiHandler,
	}

	return &server{
		svr:    &srv,
		logger: logging.NewLogger(),
		txID:   txID,
	}
}

// Start runs ListenAndServe on the http.Server with graceful shutdown.
func (s *server) Start() {
	s.logger.Info(s.txID).Infof("starting server at %v", s.svr.Addr)
	go func() {
		if err := s.svr.ListenAndServe(); err != http.ErrServerClosed {
			s.logger.Fatal(s.txID, FailedToStartServer).Errorf("Failed to start server with error %v", err)
		}
	}()

	s.startTimeStampUTC = time.Now().UTC()
	viper.Set("service_started_timestamp_utc", time.Now().UTC())

	s.logger.Info(s.txID).Infof("Server listening on %s", s.svr.Addr)
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	sig := <-quit

	s.logger.Info(s.txID).Infof("Shutting down server... Reason:%v", sig)
	if err := s.svr.Shutdown(context.Background()); err != nil {
		s.logger.Fatal(s.txID, FailedToStopServer).Errorf("Failed to stop server with error %v", err)
	}

	s.logger.Info(s.txID).Info("Server gracefully stopped")
}

func (s *server) StartTimeStampUTC() time.Time {
	return s.startTimeStampUTC
}
