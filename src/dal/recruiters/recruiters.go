package recruiters

import (
	"context"
	"fmt"
	"time"

	"github.com/jobbox-tech/recruiter-api/database"
	"github.com/jobbox-tech/recruiter-api/database/dbmodels"
	"github.com/jobbox-tech/recruiter-api/logging"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type recruiters struct {
	log logging.Logger
	db  database.MongoStore
}

func New() *recruiters {
	return &recruiters{
		log: logging.NewLogger(),
		db:  database.New(),
	}
}

// Create creates a new account.
func (r *recruiters) Create(txID string, account *dbmodels.Recruiters) (primitive.ObjectID, error) {
	rc := r.db.Database().Collection(viper.GetString("db.recruiters_collection"))
	ctx, cancel := context.WithTimeout(
		context.Background(),
		time.Duration(viper.GetInt("db.query_timeout_in_sec"))*time.Second,
	)
	defer cancel()

	insertResult, err := rc.InsertOne(ctx, account)
	if err != nil {
		return primitive.ObjectID{}, fmt.Errorf("Failed to create recruiter with error %v", err)
	}

	return insertResult.InsertedID.(primitive.ObjectID), nil
}
