package dbmodels

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Job model represents the job collection in database
type Job struct {
	ID                  primitive.ObjectID `json:"id" bson:"_id"`
	CreatedTimestampUTC time.Time          `json:"createdTimestampUTC" bson:"CreatedTimestampUTC"`
	UpdatedTimestampUTC time.Time          `json:"updatedTimestampUTC" bson:"UpdatedTimestampUTC"`

	Title    string                 `json:"title" bson:"Title"`
	Summary  map[string]interface{} `json:"summary" bson:"Summary"`
	SideNote string                 `json:"sideNote" bson:"SideNote"`

	Locations        []string `json:"locations" bson:"Locations"`
	MustHaveSkills   []string `json:"mustHaveSkills" bson:"MustHaveSkills"`
	GoodToHaveSkills []string `json:"goodToHaveSkills" bson:"GoodToHaveSkills"`
	MinExperience    int      `json:"minExperience" bson:"MinExperience"`
	MaxExperience    int      `json:"maxExperience" bson:"MaxExperience"`
	EmploymentType   []string `json:"employmentType" bson:"EmploymentType"`
	Category         string   `json:"category" bson:"Category"`
	Function         string   `json:"function" bson:"Function"`

	IsRemote        bool      `json:"isRemote" bson:"IsRemote"`
	RemoteTimezone  time.Time `json:"remoteTimezone" bson:"RemoteTimezone"`
	VisaSponsorShip bool      `json:"visaSponsorShip" bson:"VisaSponsorShip"`

	IsVerified bool `json:"isVerified" bson:"IsVerified"`
	Deleted    bool `json:"deleted" bson:"Deleted"`

	RecruiterID    primitive.ObjectID `json:"recruiterID" bson:"RecruiterID"`
	OrganizationID primitive.ObjectID `json:"organizationID" bson:"OrganizationID"`
}
