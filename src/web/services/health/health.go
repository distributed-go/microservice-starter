package health

import (
	"net/http"
	"os"
	"runtime"
	"time"

	"github.com/jobbox-tech/recruiter-api/database/connection"

	"github.com/go-chi/render"
	"github.com/jobbox-tech/recruiter-api/logging"
	"github.com/jobbox-tech/recruiter-api/proto/v1/health/v1health"
	_ "github.com/jobbox-tech/recruiter-api/web/renderers" // swag
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

// ShowAccount godoc
// @Summary Get health of the service
// @Description It returns the health of the service
// @Tags health
// @Accept  json
// @Produce  json
// @Success 200 {object} v1health.Health{}
// @Failure 400 {object} v1error.ErrorResponse{}
// @Failure 404 {object} v1error.ErrorResponse{}
// @Failure 500 {object} v1error.ErrorResponse{}
// @Router /health [get]
// GetHealth returns heath of service, can be extended if
// service is running on multile instances
func (h *health) GetHealth(w http.ResponseWriter, r *http.Request) {
	txID := r.Header["transaction_id"][0]

	healthStatus := v1health.Health{}
	healthStatus.ServiceName = viper.GetString("service_name")
	healthStatus.ServiceProvider = viper.GetString("service_provider")
	healthStatus.ServiceVersion = viper.GetString("service_version")
	healthStatus.TimestampUtc = time.Now().UTC().String()
	healthStatus.ServiceStatus = v1health.ServiceStatus_Running
	healthStatus.ServiceStartTimeUtc = viper.GetTime("service_started_timestamp_utc").String()
	healthStatus.Uptime = time.Since(viper.GetTime("service_started_timestamp_utc")).Hours()

	inbound := []*v1health.InboundConnection{}
	outbound := []*v1health.OutboundConnection{}

	// add mongo connection status
	mongo := h.db.Health()
	outbound = append(outbound, mongo)

	// add internal server details
	name, _ := os.Hostname()

	server := v1health.InboundConnection{}
	server.Hostname = name
	server.Os = runtime.GOOS
	server.TimestampUtc = time.Now().UTC().String()
	server.ApplicationName = viper.GetString("service_name")
	server.ConnectionStatus = v1health.ConnectionStatus_Active

	exIP, err := externalIP()
	if err != nil {
		h.logger.Error(txID, FailedToObtainOutboundIP).Error("Failed to obtain inbound ip address with error %v", err)
		server.ConnectionStatus = v1health.ConnectionStatus_Disconnected
	}
	server.Address = exIP
	inbound = append(inbound, &server)

	healthStatus.InboundConnections = inbound
	healthStatus.OutboundConnections = outbound

	render.JSON(w, r, &healthStatus)
}
