package pwdless

import (
	"time"

	"github.com/jobbox-tech/recruiter-api/logging"
)

func (rs *Resource) choresTicker() {
	ticker := time.NewTicker(time.Hour * 1)
	go func() {
		for range ticker.C {
			if err := rs.Store.PurgeExpiredToken(); err != nil {
				logging.Logger.WithField("chore", "purgeExpiredToken").Error(err)
			}
		}
	}()
}
