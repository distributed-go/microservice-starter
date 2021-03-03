package recruitersdal

import (
	"github.com/jobbox-tech/recruiter-api/database/dbmodels"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RecruitersDal interface {
	Create(txID string, account *dbmodels.Recruiters) (primitive.ObjectID, error)
}
