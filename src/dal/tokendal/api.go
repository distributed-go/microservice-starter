package tokendal

import (
	"github.com/jobbox-tech/recruiter-api/database/dbmodels"
)

type TokenDal interface {
	Create(txID string, token *dbmodels.Token) (*dbmodels.Token, error)
	GetByUUID(uuid string) (*dbmodels.Token, error)
	Update(token *dbmodels.Token) error
}
