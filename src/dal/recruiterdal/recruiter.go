package recruiterdal

import (
	"context"
	"fmt"
	"time"

	"github.com/jobbox-tech/recruiter-api/database/connection"
	"github.com/jobbox-tech/recruiter-api/database/dbmodels"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/mgo.v2/bson"
)

type recruiter struct {
	db connection.MongoStore
}

// NewRecruiterDal ...
func NewRecruiterDal() RecruiterDal {
	return &recruiter{
		db: connection.NewMongoStore(),
	}
}

// Create creates a new account.
func (r *recruiter) Create(txID string, account *dbmodels.Recruiter) (*dbmodels.Recruiter, error) {
	rc := r.db.Database().Collection(viper.GetString("db.recruiters_collection"))
	ctx, cancel := context.WithTimeout(
		context.Background(),
		time.Duration(viper.GetInt("db.query_timeout_in_sec"))*time.Second,
	)
	defer cancel()

	if err := account.Validate(); err != nil {
		return nil, err
	}

	insertResult, err := rc.InsertOne(ctx, account)
	if err != nil {
		return nil, fmt.Errorf("Failed to create recruiter with error %v", err)
	}

	account.ID = insertResult.InsertedID.(primitive.ObjectID)
	return account, nil
}

func (r *recruiter) Update(recruiter *dbmodels.Recruiter) error {
	rc := r.db.Database().Collection(viper.GetString("db.recruiters_collection"))
	ctx, cancel := context.WithTimeout(
		context.Background(),
		time.Duration(viper.GetInt("db.query_timeout_in_sec"))*time.Second,
	)
	defer cancel()

	if err := recruiter.Validate(); err != nil {
		return err
	}

	if _, err := rc.ReplaceOne(ctx, bson.M{"_id": recruiter.ID}, recruiter); err != nil {
		return err
	}

	return nil
}

func (r *recruiter) GetByEmail(email string) (*dbmodels.Recruiter, error) {
	rc := r.db.Database().Collection(viper.GetString("db.recruiters_collection"))
	ctx, cancel := context.WithTimeout(
		context.Background(),
		time.Duration(viper.GetInt("db.query_timeout_in_sec"))*time.Second,
	)
	defer cancel()

	var rec dbmodels.Recruiter
	if err := rc.FindOne(ctx, bson.M{"Email": email}).Decode(&rec); err != nil {
		return nil, err
	}

	return &rec, nil
}

func (r *recruiter) GetByID(id primitive.ObjectID) (*dbmodels.Recruiter, error) {
	rc := r.db.Database().Collection(viper.GetString("db.recruiters_collection"))
	ctx, cancel := context.WithTimeout(
		context.Background(),
		time.Duration(viper.GetInt("db.query_timeout_in_sec"))*time.Second,
	)
	defer cancel()

	var rec dbmodels.Recruiter
	if err := rc.FindOne(ctx, bson.M{"_id": id}).Decode(&rec); err != nil {
		return nil, err
	}

	return &rec, nil
}
