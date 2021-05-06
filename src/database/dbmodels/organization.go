package dbmodels

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Organization model represents the organization collection in database
type Organization struct {
	ID                  primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	CreatedTimestampUTC time.Time          `json:"createdTimestampUTC,omitempty" bson:"CreatedTimestampUTC,omitempty"`
	UpdatedTimestampUTC time.Time          `json:"updatedTimestampUTC,omitempty" bson:"UpdatedTimestampUTC,omitempty"`

	OrganizationName     string `json:"organizationName,omitempty" bson:"OrganizationName,omitempty"`
	OrganizationHeadline string `json:"organizationHeadline,omitempty" bson:"OrganizationHeadline,omitempty"`
	OrganizationWebsite  string `json:"organizationWebsite,omitempty" bson:"OrganizationWebsite,omitempty"`
	OrganizationLogo     string `json:"organizationLogo,omitempty" bson:"OrganizationLogo,omitempty"`
	OrganizationSize     string `json:"organizationSize,omitempty" bson:"OrganizationSize,omitempty"`

	Markets        []string         `json:"markets,omitempty" bson:"Markets,omitempty"`
	FundingInUSD   string           `json:"fundingInUSD,omitempty" bson:"FundingInUSD,omitempty"`
	FoundingDate   time.Time        `json:"foundingDate,omitempty" bson:"FoundingDate,omitempty"`
	SocialProfiles []SocialProfiles `json:"socialProfiles,omitempty" bson:"SocialProfiles,omitempty"`
	Locations      []string         `json:"locations,omitempty" bson:"Locations,omitempty"`

	IsVerified bool `json:"isVerified,omitempty" bson:"IsVerified,omitempty"`
	IsActive   bool `json:"isActive,omitempty" bson:"IsActive,omitempty"`
	IsDisabled bool `json:"isDisabled,omitempty" bson:"IsDisabled,omitempty"`

	CreatedBy primitive.ObjectID   `json:"createdBy,omitempty" bson:"CreatedBy,omitempty"`
	Admins    []primitive.ObjectID `json:"admins,omitempty" bson:"Admins,omitempty"`
}

// SocialProfiles holds the social profiles of a company
type SocialProfiles struct {
	Name string `json:"Name,omitempty" bson:"Name,omitempty"`
	URL  string `json:"URL,omitempty" bson:"URL,omitempty"`
}
