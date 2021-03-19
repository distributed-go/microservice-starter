package dbmodels

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/jobbox-tech/recruiter-api/models/authmodel"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Token model represents the recruiter collection in database
type Token struct {
	ID                  primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	CreatedTimestampUTC time.Time          `json:"CreatedTimestampUTC,omitempty" bson:"CreatedTimestampUTC,omitempty"`
	UpdatedTimestampUTC time.Time          `json:"UpdatedTimestampUTC,omitempty" bson:"UpdatedTimestampUTC,omitempty"`

	ExpiryTimestampUTC time.Time          `json:"-" bson:"ExpiryTimestampUTC,omitempty"`
	AccountID          primitive.ObjectID `json:"-" bson:"AccountID,omitempty"`
	AccessToken        string             `json:"-" bson:"AccessToken,omitempty"`
	ResfreshToken      string             `json:"-" bson:"ResfreshToken,omitempty"`
	Claimed            bool               `json:"-" bson:"Claimed"`

	Mobile    bool   `json:"Mobile,omitempty" bson:"Mobile,omitempty"`
	TokenUUID string `json:"TokenUUID,omitempty" bson:"TokenUUID,omitempty"`
	UserAgent string `json:"UserAgent,omitempty" bson:"UserAgent,omitempty"`
}

// Validate validates struct
func (t Token) Validate() error {
	return validation.ValidateStruct(&t,
		validation.Field(&t.CreatedTimestampUTC, validation.Required),
		validation.Field(&t.UpdatedTimestampUTC, validation.Required),
		validation.Field(&t.ExpiryTimestampUTC, validation.Required),

		validation.Field(&t.AccountID, validation.Required),
		validation.Field(&t.TokenUUID, validation.Required),
	)
}

// Claims returns the token claims to be signed
func (t *Token) Claims() authmodel.RefreshClaims {
	return authmodel.RefreshClaims{
		ID:        t.ID.String(),
		TokenUUID: t.TokenUUID,
	}
}
