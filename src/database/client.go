package database

import (
	"context"
	"sync"
	"time"

	"github.com/jobbox-tech/recruiter-api/logging"

	"github.com/google/uuid"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	once   sync.Once
	db     *mongo.Database
	client *mongo.Client
)

// MongoStore ...
type mongoStore struct {
	logger logging.Logger
	txID   string
}

// New returns new instance of datastore
func New() MongoStore {
	return &mongoStore{
		logger: logging.NewLogger(),
		txID:   uuid.New().String(),
	}
}

// Client returns mongodb client instance
func (s *mongoStore) Client() *mongo.Client {
	once.Do(func() {
		db, client = s.initialize()
	})

	return client
}

// Database returns mongodb database instance
func (s *mongoStore) Database() *mongo.Database {
	once.Do(func() {
		db, client = s.initialize()
	})

	return db
}

func (s *mongoStore) initialize() (a *mongo.Database, b *mongo.Client) {
	client, err := mongo.NewClient(options.Client().ApplyURI(viper.GetString("db.host")))
	if err != nil {
		s.logger.Fatal(s.txID, DBConnectionFailed).Fatalf("Failed to connect to database with error: %v", err)
	}

	ctx, cancel := context.WithTimeout(
		context.Background(),
		time.Duration(viper.GetInt("db.query_timeout_in_sec"))*time.Second,
	)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		s.logger.Fatal(s.txID, DBConnectionFailed).Fatalf("Failed to connect to database with error: %v", err)
	}

	database := viper.GetString("db.database")
	db := client.Database(database)
	s.logger.Info(s.txID).Infof("Successfully connected to database %s", database)

	return db, client
}
