package job

import (
	"medication-notifier/utils/logger"
)

func (j *Job) activityLogCleanupJob() {
	logger.Info("start activityLogCleanupJob")
	// Find all newest activity logs
	activities := j.activityLogData.FetchAllNewestGroupedByClientId()

	// For each check is refresh token exists in temp db
	for _, activity := range activities {
		logger.Info("activity_log token_hash: %s client_id: %s", activity.RefreshTokenHash, activity.ClientId)
		if _, err := j.tokenData.FindByTokenHash(activity.RefreshTokenHash, activity.ClientId); err != nil {
			if err.Error() == "not found" {
				logger.Info("activity_log removing token_hash: %s client_id: %s", activity.RefreshTokenHash, activity.ClientId)

				if removeErr := j.activityLogData.RemoveByClientId(activity.ClientId); removeErr != nil {
					logger.Warn("remove activity log by client_id %s failed, err: %s", activity.ClientId, removeErr)
				}
			}
		}
	}

	logger.Info("end activityLogCleanupJob")
}
