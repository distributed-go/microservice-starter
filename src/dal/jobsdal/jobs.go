package jobsdal

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

type jobs struct {
	db connection.MongoStore
}

// NewJobsDal ...
func NewJobsDal() JobsDal {
	return &jobs{
		db: connection.NewMongoStore(),
	}
}

// Create creates a new account.
func (r *jobs) Create(txID string, job *dbmodels.Job) (*dbmodels.Job, error) {
	rc := r.db.Database().Collection(viper.GetString("db.jobs_collection"))
	ctx, cancel := context.WithTimeout(
		context.Background(),
		time.Duration(viper.GetInt("db.query_timeout_in_sec"))*time.Second,
	)
	defer cancel()

	insertResult, err := rc.InsertOne(ctx, job)
	if err != nil {
		return nil, fmt.Errorf("Failed to create organization with error %v", err)
	}

	job.ID = insertResult.InsertedID.(primitive.ObjectID)
	return job, nil
}

func (r *jobs) Update(job *dbmodels.Job) error {
	rc := r.db.Database().Collection(viper.GetString("db.jobs_collection"))
	ctx, cancel := context.WithTimeout(
		context.Background(),
		time.Duration(viper.GetInt("db.query_timeout_in_sec"))*time.Second,
	)
	defer cancel()

	if _, err := rc.ReplaceOne(ctx, bson.M{"_id": job.ID}, job); err != nil {
		return err
	}

	return nil
}

func (r *jobs) GetByID(id primitive.ObjectID) (*dbmodels.Job, error) {
	rc := r.db.Database().Collection(viper.GetString("db.jobs_collection"))
	ctx, cancel := context.WithTimeout(
		context.Background(),
		time.Duration(viper.GetInt("db.query_timeout_in_sec"))*time.Second,
	)
	defer cancel()

	var rec dbmodels.Job
	if err := rc.FindOne(ctx, bson.M{"_id": id}).Decode(&rec); err != nil {
		return nil, err
	}

	return &rec, nil
}

func (r *jobs) DeleteByID(id primitive.ObjectID) error {
	rc := r.db.Database().Collection(viper.GetString("db.jobs_collection"))
	ctx, cancel := context.WithTimeout(
		context.Background(),
		time.Duration(viper.GetInt("db.query_timeout_in_sec"))*time.Second,
	)
	defer cancel()

	if _, err := rc.UpdateOne(
		ctx,
		bson.M{"_id": id},
		bson.D{
			{"$set", bson.D{{"Deleted", true}}},
		},
	); err != nil {
		return err
	}

	return nil
}
