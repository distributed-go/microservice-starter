package jobsservice

import (
	"github.com/jobbox-tech/recruiter-api/dal/organizationdal"
	"github.com/jobbox-tech/recruiter-api/dal/recruiterdal"
	"github.com/jobbox-tech/recruiter-api/dal/tokendal"
	"github.com/jobbox-tech/recruiter-api/logging"
)

type jobsservice struct {
	logger logging.Logger

	tokenDal     tokendal.TokenDal
	recruiterDal recruiterdal.RecruiterDal
	orgDal       organizationdal.OrganizationDal
}

// NewJobService returns service impl
func NewJobService() JobService {
	return &jobsservice{
		logger: logging.NewLogger(),

		tokenDal:     tokendal.NewTokenDal(),
		recruiterDal: recruiterdal.NewRecruiterDal(),
		orgDal:       organizationdal.NewOrganizationDal(),
	}
}
