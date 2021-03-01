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
	TimeStampUTC        time.Time           `json:"TimeStampUTC"`
	ServiceName         string              `json:"ServiceName"`
	ServiceProvider     string              `json:"ServiceProvider"`
	ServiceVersion      string              `json:"ServiceVersion"`
	ServiceStatus       ServiceStatus       `json:"ServiceStatus"`
	ServiceStartTimeUTC time.Time           `json:"ServiceStartTimeUTC"`
	Uptime              float64             `json:"Uptime"`
	InboundInterfaces   []InboundInterface  `json:"InboundInterfaces"`
	OutboundInterfaces  []OutboundInterface `json:"OutboundInterfaces"`
}

// InboundInterface inbound network inferfaces
type InboundInterface struct {
	ApplicationName  string           `json:"ApplicationName"`
	ConnectionStatus ConnectionStatus `json:"ConnectionStatus"`
	TimeStampUTC     time.Time        `json:"TimeStampUTC"`
	Hostname         string           `json:"Hostname"`
	Address          string           `json:"Address"`
	OS               string           `json:"OS"`
}

// OutboundInterface outbound network interfaces
type OutboundInterface struct {
	ApplicationName  string           `json:"ApplicationName"`
	TimeStampUTC     time.Time        `json:"TimeStampUTC"`
	URLs             []string         `json:"URLs"`
	ConnectionStatus ConnectionStatus `json:"ConnectionStatus"`
}
