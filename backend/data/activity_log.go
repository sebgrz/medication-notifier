package data

type ActivityLog struct {
	ClientId         string
	UserId           string
	RefreshTokenHash string
	ExpireTime       int64
}

type ActivityLogDataService interface {
	Add(clientId, userId, refreshTokenHash string, expireTime int64) error
	FetchAllNewestGroupedByClientId() []ActivityLog
	RemoveByUserIdAndClientId(string, string) error
}
