package recruiterdal

import (
	"github.com/jobbox-tech/recruiter-api/database/dbmodels"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// RecruiterDal ...
type RecruiterDal interface {
	Create(txID string, account *dbmodels.Recruiter) (*dbmodels.Recruiter, error)
	GetByEmail(email string) (*dbmodels.Recruiter, error)
	GetByID(id primitive.ObjectID) (*dbmodels.Recruiter, error)
	Update(recruiter *dbmodels.Recruiter) error
}
