package jobsdal

import (
	"github.com/jobbox-tech/recruiter-api/database/dbmodels"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// OrganizationDal ...
type OrganizationDal interface {
	Create(txID string, account *dbmodels.Job) (*dbmodels.Job, error)
	Update(recruiter *dbmodels.Job) error
	GetByID(id primitive.ObjectID) (*dbmodels.Job, error)
	DeleteByID(id primitive.ObjectID) error
}
