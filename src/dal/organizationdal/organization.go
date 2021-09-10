package organizationdal

import (
	"context"
	"fmt"
	"time"

	"github.com/distributed-go/microservice-starter/database/connection"
	"github.com/distributed-go/microservice-starter/database/dbmodels"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/mgo.v2/bson"
)

type org struct {
	db connection.MongoStore
}

// NewOrganizationDal ...
func NewOrganizationDal() OrganizationDal {
	return &org{
		db: connection.NewMongoStore(),
	}
}

// Create creates a new account.
func (r *org) Create(txID string, organization *dbmodels.Organization) (*dbmodels.Organization, error) {
	rc := r.db.Database().Collection(viper.GetString("db.organizations_collection"))
	ctx, cancel := context.WithTimeout(
		context.Background(),
		time.Duration(viper.GetInt("db.query_timeout_in_sec"))*time.Second,
	)
	defer cancel()

	insertResult, err := rc.InsertOne(ctx, organization)
	if err != nil {
		return nil, fmt.Errorf("Failed to create organization with error %v", err)
	}

	organization.ID = insertResult.InsertedID.(primitive.ObjectID)
	return organization, nil
}

func (r *org) Update(organization *dbmodels.Organization) error {
	rc := r.db.Database().Collection(viper.GetString("db.organizations_collection"))
	ctx, cancel := context.WithTimeout(
		context.Background(),
		time.Duration(viper.GetInt("db.query_timeout_in_sec"))*time.Second,
	)
	defer cancel()

	if _, err := rc.ReplaceOne(ctx, bson.M{"_id": organization.ID}, organization); err != nil {
		return err
	}

	return nil
}

func (r *org) GetByID(id primitive.ObjectID) (*dbmodels.Organization, error) {
	rc := r.db.Database().Collection(viper.GetString("db.organizations_collection"))
	ctx, cancel := context.WithTimeout(
		context.Background(),
		time.Duration(viper.GetInt("db.query_timeout_in_sec"))*time.Second,
	)
	defer cancel()

	var rec dbmodels.Organization
	if err := rc.FindOne(ctx, bson.M{"_id": id}).Decode(&rec); err != nil {
		return nil, err
	}

	return &rec, nil
}
