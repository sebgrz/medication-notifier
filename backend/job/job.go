package job

import (
	"medication-notifier/data"
	"medication-notifier/utils/logger"

	"github.com/robfig/cron"
)

type Job struct {
	tokenData       data.TokenDataService
	activityLogData data.ActivityLogDataService

	cron *cron.Cron
}

func New(
	tokenData data.TokenDataService,
	activityLogData data.ActivityLogDataService,
) *Job {
	c := cron.New()
	j := &Job{
		tokenData:       tokenData,
		activityLogData: activityLogData,
		cron:            c,
	}

	// TODO change to 5m
	c.AddFunc("@every 1m", j.activityLogCleanupJob)

	return j
}

func (j *Job) Start() {
	j.cron.Start()
	logger.Info("cron jobs started")
}
