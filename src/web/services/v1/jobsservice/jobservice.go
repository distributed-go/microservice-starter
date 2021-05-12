package jobsservice

import (
	"github.com/jobbox-tech/recruiter-api/dal/jobsdal"
	"github.com/jobbox-tech/recruiter-api/dal/organizationdal"
	"github.com/jobbox-tech/recruiter-api/dal/recruiterdal"
	"github.com/jobbox-tech/recruiter-api/dal/tokendal"
	"github.com/jobbox-tech/recruiter-api/logging"
	"github.com/jobbox-tech/recruiter-api/web/middlewares"
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
