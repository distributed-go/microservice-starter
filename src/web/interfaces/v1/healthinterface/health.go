package healthinterface

import "time"

// ConnectionType ...
type ServiceStatus string
type ConnectionStatus string

var (
	ServiceRunning  ServiceStatus = "Running"
	ServiceDegraded ServiceStatus = "Degraded"
	ServiceStopped  ServiceStatus = "Stopped"

	ConnectionActive       ConnectionStatus = "Active"
	ConnectionDisconnected ConnectionStatus = "Disconnected"
)

// Health ...
type Health struct {
	TimeStampUTC        time.Time           `json:"timeStampUTC"`
	ServiceName         string              `json:"serviceName"`
	ServiceProvider     string              `json:"serviceProvider"`
	ServiceVersion      string              `json:"serviceVersion"`
	ServiceStatus       ServiceStatus       `json:"serviceStatus"`
	ServiceStartTimeUTC time.Time           `json:"serviceStartTimeUTC"`
	Uptime              float64             `json:"uptime"`
	InboundInterfaces   []InboundInterface  `json:"inboundInterfaces"`
	OutboundInterfaces  []OutboundInterface `json:"outboundInterfaces"`
}

// InboundInterface inbound network inferfaces
type InboundInterface struct {
	ApplicationName  string           `json:"applicationName"`
	ConnectionStatus ConnectionStatus `json:"connectionStatus"`
	TimeStampUTC     time.Time        `json:"timeStampUTC"`
	Hostname         string           `json:"hostname"`
	Address          string           `json:"address"`
	OS               string           `json:"os"`
}

// OutboundInterface outbound network interfaces
type OutboundInterface struct {
	ApplicationName  string           `json:"applicationName"`
	TimeStampUTC     time.Time        `json:"timeStampUTC"`
	URLs             []string         `json:"urls"`
	ConnectionStatus ConnectionStatus `json:"connectionStatus"`
}
