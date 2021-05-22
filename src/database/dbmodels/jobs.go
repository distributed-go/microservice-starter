package dbmodels

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Job model represents the job collection in database
type Job struct {
	ID                  primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	CreatedTimestampUTC time.Time          `json:"createdTimestampUTC,omitempty" bson:"CreatedTimestampUTC,omitempty"`
	UpdatedTimestampUTC time.Time          `json:"updatedTimestampUTC,omitempty" bson:"UpdatedTimestampUTC,omitempty"`

	Title    string                 `json:"title,omitempty" bson:"Title,omitempty"`
	Summary  map[string]interface{} `json:"summary,omitempty" bson:"Summary,omitempty"`
	SideNote string                 `json:"sideNote,omitempty" bson:"SideNote,omitempty"`

	Locations        []string `json:"locations,omitempty" bson:"Locations,omitempty"`
	MustHaveSkills   []string `json:"mustHaveSkills,omitempty" bson:"MustHaveSkills,omitempty"`
	GoodToHaveSkills []string `json:"goodToHaveSkills,omitempty" bson:"GoodToHaveSkills,omitempty"`
	MinExperience    int      `json:"minExperience,omitempty" bson:"MinExperience,omitempty"`
	MaxExperience    int      `json:"maxExperience,omitempty" bson:"MaxExperience,omitempty"`
	EmploymentType   []string `json:"employmentType,omitempty" bson:"EmploymentType,omitempty"`
	Category         string   `json:"category,omitempty" bson:"Category,omitempty"`
	Function         string   `json:"function,omitempty" bson:"Function,omitempty"`

	IsRemote        bool      `json:"isRemote,omitempty" bson:"IsRemote,omitempty"`
	RemoteTimezone  time.Time `json:"remoteTimezone,omitempty" bson:"RemoteTimezone,omitempty"`
	VisaSponsorShip bool      `json:"visaSponsorShip,omitempty" bson:"VisaSponsorShip,omitempty"`

	IsVerified bool `json:"isVerified,omitempty" bson:"IsVerified,omitempty"`
	Deleted    bool `json:"deleted,omitempty" bson:"Deleted,omitempty"`

	RecruiterID    primitive.ObjectID `json:"recruiterID,omitempty" bson:"RecruiterID,omitempty"`
	OrganizationID primitive.ObjectID `json:"organizationID,omitempty" bson:"OrganizationID,omitempty"`
}
