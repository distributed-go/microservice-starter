package organizationdal

import (
	"github.com/distributed-go/microservice-starter/database/dbmodels"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// OrganizationDal ...
type OrganizationDal interface {
	Create(txID string, account *dbmodels.Organization) (*dbmodels.Organization, error)
	Update(recruiter *dbmodels.Organization) error
	GetByID(id primitive.ObjectID) (*dbmodels.Organization, error)
}
