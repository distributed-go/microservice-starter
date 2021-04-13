package connection

import (
	"github.com/jobbox-tech/recruiter-api/proto/v1/health/v1healthpb"
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
	Health() *v1healthpb.OutboundConnection
}
