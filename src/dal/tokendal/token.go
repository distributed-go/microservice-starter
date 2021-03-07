package tokendal

import (
	"context"
	"fmt"
	"time"

	"github.com/jobbox-tech/recruiter-api/database/connection"
	"github.com/jobbox-tech/recruiter-api/database/dbmodels"
	"github.com/jobbox-tech/recruiter-api/logging"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type token struct {
	log logging.Logger
	db  connection.MongoStore
}

// NewTokenDal ...
func NewTokenDal() TokenDal {
	return &token{
		log: logging.NewLogger(),
		db:  connection.NewMongoStore(),
	}
}

// Create creates a new account.
func (r *token) Create(txID string, account *dbmodels.Recruiter) (primitive.ObjectID, error) {
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
