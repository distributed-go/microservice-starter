package tokendal

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/jobbox-tech/recruiter-api/database/connection"
	"github.com/jobbox-tech/recruiter-api/database/dbmodels"
	"github.com/jobbox-tech/recruiter-api/logging"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/mgo.v2/bson"
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
func (t *token) Create(txID string, token *dbmodels.Token) (*dbmodels.Token, error) {
	tc := t.db.Database().Collection(viper.GetString("db.access_tokens_collection"))
	ctx, cancel := context.WithTimeout(
		context.Background(),
		time.Duration(viper.GetInt("db.query_timeout_in_sec"))*time.Second,
	)
	defer cancel()

	if err := token.Validate(); err != nil {
		return nil, fmt.Errorf("Failed to create the access token with the error %v", err)
	}

	insertResult, err := tc.InsertOne(ctx, token)
	if err != nil {
		return nil, fmt.Errorf("Failed to create recruiter with error %v", err)
	}

	token.ID = insertResult.InsertedID.(primitive.ObjectID)
	return token, nil
}

func (t *token) Update(token *dbmodels.Token) error {
	tc := t.db.Database().Collection(viper.GetString("db.access_tokens_collection"))
	ctx, cancel := context.WithTimeout(
		context.Background(),
		time.Duration(viper.GetInt("db.query_timeout_in_sec"))*time.Second,
	)
	defer cancel()

	if err := token.Validate(); err != nil {
		return fmt.Errorf("Failed to create the access token with the error %v", err)
	}

	if _, err := tc.ReplaceOne(ctx, bson.M{"_id": token.ID}, token); err != nil {
		return err
	}

	return nil
}

func (t *token) DeleteByAccessToken(token string) error {
	tc := t.db.Database().Collection(viper.GetString("db.access_tokens_collection"))
	ctx, cancel := context.WithTimeout(
		context.Background(),
		time.Duration(viper.GetInt("db.query_timeout_in_sec"))*time.Second,
	)
	defer cancel()

	if _, err := tc.DeleteOne(ctx, bson.M{"AccessToken": strings.TrimSpace(token)}); err != nil {
		return err
	}

	return nil
}

func (t *token) GetByUUID(uuid string) (*dbmodels.Token, error) {
	tc := t.db.Database().Collection(viper.GetString("db.access_tokens_collection"))
	ctx, cancel := context.WithTimeout(
		context.Background(),
		time.Duration(viper.GetInt("db.query_timeout_in_sec"))*time.Second,
	)
	defer cancel()

	var rec dbmodels.Token
	if err := tc.FindOne(ctx, bson.M{"TokenUUID": uuid}).Decode(&rec); err != nil {
		return nil, err
	}

	return &rec, nil
}
