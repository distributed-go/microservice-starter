package dbmodels

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Tokens model
type Tokens struct {
	ID                  primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	CreatedTimestampUTC *time.Time         `json:"CreatedTimestampUTC,omitempty" bson:"CreatedTimestampUTC,omitempty"`
	UpdatedTimestampUTC *time.Time         `json:"UpdatedTimestampUTC,omitempty" bson:"UpdatedTimestampUTC,omitempty"`
	AccountID           primitive.ObjectID `json:"AccountID,omitempty" bson:"AccountID,omitempty"`
	Token               string             `json:"Token,omitempty" bson:"Token,omitempty"`
	Expiry              *time.Time         `json:"Expiry,omitempty" bson:"Expiry,omitempty"`
	Mobile              bool               `json:"Mobile,omitempty" bson:"Mobile,omitempty"`
	Identifier          string             `json:"Identifier,omitempty" bson:"Identifier,omitempty"`
}
