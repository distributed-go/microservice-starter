package dbmodels

import (
	"time"

	"github.com/distributed-go/microservice-starter/models/authmodel"
	"github.com/distributed-go/microservice-starter/models/recruitermodel"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Recruiter model represents the recruiter collection in database
type Recruiter struct {
	ID                  primitive.ObjectID    `json:"ID,omitempty" bson:"_id,omitempty"`
	CreatedTimestampUTC time.Time             `json:"CreatedTimestampUTC,omitempty" bson:"CreatedTimestampUTC,omitempty"`
	UpdatedTimestampUTC time.Time             `json:"UpdatedTimestampUTC,omitempty" bson:"UpdatedTimestampUTC,omitempty"`
	LastLogin           time.Time             `json:"LastLogin,omitempty" bson:"LastLogin,omitempty"`
	Email               string                `json:"Email,omitempty" bson:"Email,omitempty"`
	FirstName           string                `json:"FirstName,omitempty" bson:"FirstName,omitempty"`
	LastName            string                `json:"LastName,omitempty" bson:"LastName,omitempty"`
	Designation         string                `json:"Designation,omitempty" bson:"Designation,omitempty"`
	Active              bool                  `json:"Active,omitempty" bson:"Active,omitempty"`
	Roles               []recruitermodel.Role `json:"Roles,omitempty" bson:"Roles,omitempty"`
	VerifiedAccount     bool                  `json:"VerifiedAccount,omitempty" bson:"VerifiedAccount,omitempty"`
	SignUpCompleted     bool                  `json:"SignUpCompleted,omitempty" bson:"SignUpCompleted,omitempty"`
	OrganizationID      primitive.ObjectID    `json:"OrganizationID,omitempty" bson:"OrganizationID,omitempty"`
}

// Validate validates struct
func (r Recruiter) Validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.Email, validation.Required, is.Email),
		validation.Field(&r.FirstName, validation.Required, validation.Length(1, 50)),
		validation.Field(&r.Roles, validation.Required, validation.Length(1, 3)),
	)
}

// CanLogin returns true if user is allowed to login.
func (r *Recruiter) CanLogin() bool {
	return r.Active
}

// Claims returns the account's claims to be signed
func (r *Recruiter) Claims() authmodel.AppClaims {
	return authmodel.AppClaims{
		ID:    r.ID.String(),
		Sub:   r.Email,
		Roles: r.Roles,
	}
}
