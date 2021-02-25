package server

const (
	FailedToStartServer = "Failed-To-Start-Server"
	FailedToStopServer  = "Failed-To-Stop-Server"
)

type Server interface {
	Start()
}
