package dbmodels

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Organization model represents the organization collection in database
type Organization struct {
	ID                  primitive.ObjectID `json:"ID,omitempty" bson:"_id,omitempty"`
	CreatedTimestampUTC time.Time          `json:"CreatedTimestampUTC,omitempty" bson:"CreatedTimestampUTC,omitempty"`
	UpdatedTimestampUTC time.Time          `json:"UpdatedTimestampUTC,omitempty" bson:"UpdatedTimestampUTC,omitempty"`
	OrganizationName    time.Time          `json:"OrganizationName,omitempty" bson:"OrganizationName,omitempty"`
	Description         string             `json:"Description,omitempty" bson:"Description,omitempty"`
	OrganizationWebsite string             `json:"OrganizationWebsite,omitempty" bson:"OrganizationWebsite,omitempty"`
	OrganizationLogo    string             `json:"OrganizationLogo,omitempty" bson:"OrganizationLogo,omitempty"`

	Markets          []string         `json:"Markets,omitempty" bson:"Markets,omitempty"`
	FundingInUSD     string           `json:"FundingInUSD,omitempty" bson:"FundingInUSD,omitempty"`
	OrganizationSize string           `json:"OrganizationSize,omitempty" bson:"OrganizationSize,omitempty"`
	FoundingDate     time.Time        `json:"FoundingDate,omitempty" bson:"FoundingDate,omitempty"`
	SocialProfiles   []SocialProfiles `json:"SocialProfiles,omitempty" bson:"SocialProfiles,omitempty"`
	Locations        []string         `json:"Locations,omitempty" bson:"Locations,omitempty"`

	IsVerified bool `json:"IsVerified,omitempty" bson:"IsVerified,omitempty"`
	IsActive   bool `json:"IsActive,omitempty" bson:"IsActive,omitempty"`
	IsDisabled bool `json:"IsDisabled,omitempty" bson:"IsDisabled,omitempty"`

	CreatedBy primitive.ObjectID   `json:"CreatedBy,omitempty" bson:"CreatedBy,omitempty"`
	Admins    []primitive.ObjectID `json:"Admins,omitempty" bson:"Admins,omitempty"`
}

// SocialProfiles holds the social profiles of a company
type SocialProfiles struct {
	Name string `json:"Name,omitempty" bson:"Name,omitempty"`
	URL  string `json:"URL,omitempty" bson:"URL,omitempty"`
}
