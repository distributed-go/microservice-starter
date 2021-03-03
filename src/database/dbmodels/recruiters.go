package dbmodels

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Role string

const (
	ADMIN          Role = "ADMIN"
	USER           Role = "USER"
	PLATFORM_ADMIN Role = "PLATFORM_ADMIN"
)

// Recruiters model
type Recruiters struct {
	ID                  primitive.ObjectID `json:"ID,omitempty" bson:"_id,omitempty"`
	CreatedTimestampUTC *time.Time         `json:"CreatedTimestampUTC,omitempty" bson:"CreatedTimestampUTC,omitempty"`
	UpdatedTimestampUTC *time.Time         `json:"UpdatedTimestampUTC,omitempty" bson:"UpdatedTimestampUTC,omitempty"`
	LastLogin           *time.Time         `json:"LastLogin,omitempty" bson:"LastLogin,omitempty"`
	Email               string             `json:"Email,omitempty" bson:"Email,omitempty"`
	FirstName           string             `json:"FirstName,omitempty" bson:"FirstName,omitempty"`
	LastName            string             `json:"LastName,omitempty" bson:"LastName,omitempty"`
	Active              bool               `json:"Active,omitempty" bson:"Active,omitempty"`
	Roles               []Role             `json:"Roles,omitempty" bson:"Roles,omitempty"`
}

// Validate validates struct
func (r Recruiters) Validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.Email, validation.Required, is.Email),
		validation.Field(&r.FirstName, validation.Required, validation.Length(1, 50)),
		validation.Field(&r.Roles, validation.Required, validation.Length(1, 3)),
	)
}
