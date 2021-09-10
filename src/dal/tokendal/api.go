package tokendal

import (
	"github.com/distributed-go/microservice-starter/database/dbmodels"
)

type TokenDal interface {
	Create(txID string, token *dbmodels.Token) (*dbmodels.Token, error)
	GetByUUID(uuid string) (*dbmodels.Token, error)
	Update(token *dbmodels.Token) error
	DeleteByAccessToken(token string) error
}
