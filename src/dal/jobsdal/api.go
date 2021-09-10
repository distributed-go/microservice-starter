package jobsdal

import (
	"github.com/distributed-go/microservice-starter/database/dbmodels"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// JobsDal ...
type JobsDal interface {
	Create(txID string, account *dbmodels.Job) (*dbmodels.Job, error)
	Update(recruiter *dbmodels.Job) error
	GetByID(id primitive.ObjectID) (*dbmodels.Job, error)
}
