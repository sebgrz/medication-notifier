package data

type ActivityLogDataService interface {
	Add(clientId, userId, refreshTokenHash string, expiryTime int64) error
}
