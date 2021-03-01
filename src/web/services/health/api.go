package health

import "net/http"

const (
	FailedToObtainOutboundIP = "Failed-To-Obtain-Outbound-IP"
)

// Health interface
type Health interface {
	GetHealth(w http.ResponseWriter, r *http.Request)
}
