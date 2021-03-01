package health

import (
	"net/http"
	"os"
	"runtime"
	"time"

	"github.com/jobbox-tech/recruiter-api/database/connection"

	"github.com/go-chi/render"
	"github.com/jobbox-tech/recruiter-api/logging"
	"github.com/jobbox-tech/recruiter-api/web/interfaces/v1/healthinterface"
	"github.com/spf13/viper"
)

type health struct {
	logger logging.Logger
	db     connection.MongoStore
}

// NewHealth returns health impl
func NewHealth() Health {
	return &health{
		logger: logging.NewLogger(),
		db:     connection.NewMongoStore(),
	}
}

// GetHealth returns heath of service, can be extended if
// service is running on multile instances
func (h *health) GetHealth(w http.ResponseWriter, r *http.Request) {
	txID := r.Header["transaction_id"][0]

	healthStatus := healthinterface.Health{}
	healthStatus.ServiceName = viper.GetString("service_name")
	healthStatus.ServiceProvider = viper.GetString("service_provider")
	healthStatus.ServiceVersion = viper.GetString("service_version")
	healthStatus.TimeStampUTC = time.Now().UTC()
	healthStatus.ServiceStatus = healthinterface.ServiceRunning
	healthStatus.ServiceStartTimeUTC = viper.GetTime("service_started_timestamp_utc")
	healthStatus.Uptime = time.Since(viper.GetTime("service_started_timestamp_utc")).Hours()

	inbound := []healthinterface.InboundInterface{}
	outbound := []healthinterface.OutboundInterface{}

	// add mongo connection status
	mongo := h.db.Health()
	outbound = append(outbound, *mongo)

	// add internal server details
	name, _ := os.Hostname()

	server := healthinterface.InboundInterface{}
	server.Hostname = name
	server.OS = runtime.GOOS
	server.TimeStampUTC = time.Now().UTC()
	server.ApplicationName = viper.GetString("service_name")
	server.ConnectionStatus = healthinterface.ConnectionActive

	exIP, err := externalIP()
	if err != nil {
		h.logger.Error(txID, FailedToObtainOutboundIP).Error("Failed to obtain inbound ip address with error %v", err)
		server.ConnectionStatus = healthinterface.ConnectionDisconnected
	}
	server.Address = exIP
	inbound = append(inbound, server)

	healthStatus.InboundInterfaces = inbound
	healthStatus.OutboundInterfaces = outbound

	render.JSON(w, r, healthStatus)
}
