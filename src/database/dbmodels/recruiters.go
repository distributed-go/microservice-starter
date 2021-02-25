package dbmodels

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Role string

const (
	ADMIN Role = "ADMIN"
	USER  Role = "USER"
)

// Recruiters model
type Recruiters struct {
	ID                  primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	CreatedTimestampUTC *time.Time         `json:"CreatedTimestampUTC,omitempty" bson:"CreatedTimestampUTC,omitempty"`
	UpdatedTimestampUTC *time.Time         `json:"UpdatedTimestampUTC,omitempty" bson:"UpdatedTimestampUTC,omitempty"`
	LastLogin           *time.Time         `json:"LastLogin,omitempty" bson:"LastLogin,omitempty"`
	Email               string             `json:"Email,omitempty" bson:"Email,omitempty"`
	Name                string             `json:"Name,omitempty" bson:"Name,omitempty"`
	Active              bool               `json:"Active,omitempty" bson:"Active,omitempty"`
	Roles               []Role             `json:"Roles,omitempty" bson:"Roles,omitempty"`
}
