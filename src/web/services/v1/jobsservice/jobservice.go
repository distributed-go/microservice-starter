package jobsservice

import (
	"github.com/distributed-go/microservice-starter/dal/jobsdal"
	"github.com/distributed-go/microservice-starter/dal/organizationdal"
	"github.com/distributed-go/microservice-starter/dal/recruiterdal"
	"github.com/distributed-go/microservice-starter/dal/tokendal"
	"github.com/distributed-go/microservice-starter/logging"
	"github.com/distributed-go/microservice-starter/web/middlewares"
)

type jobsservice struct {
	logger logging.Logger

	tokenDal     tokendal.TokenDal
	recruiterDal recruiterdal.RecruiterDal
	orgDal       organizationdal.OrganizationDal
	jobsDal      jobsdal.JobsDal
	middlewares  middlewares.Middlewares
}

// NewJobService returns service impl
func NewJobService() JobsService {
	return &jobsservice{
		logger: logging.NewLogger(),

		tokenDal:     tokendal.NewTokenDal(),
		recruiterDal: recruiterdal.NewRecruiterDal(),
		orgDal:       organizationdal.NewOrganizationDal(),
		jobsDal:      jobsdal.NewJobsDal(),
		middlewares:  middlewares.NewMiddlewares(),
	}
}
