package connection

import (
	"github.com/jobbox-tech/recruiter-api/web/interfaces/v1/healthinterface"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	// DBConnectionFailed used when failed to create a database client
	DBConnectionFailed = "Database-Connection-Failed"
)

// MongoStore implents the database store
type MongoStore interface {
	Client() *mongo.Client
	Database() *mongo.Database
	Health() *healthinterface.OutboundInterface
}
