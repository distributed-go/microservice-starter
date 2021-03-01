package server

import "time"

const (
	FailedToStartServer = "Failed-To-Start-Server"
	FailedToStopServer  = "Failed-To-Stop-Server"
)

// Server implements server interface
type Server interface {
	Start()
	StartTimeStampUTC() time.Time
}
