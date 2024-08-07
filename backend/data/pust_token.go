package data

type PushToken struct {
	Id     string
	UserId string
	Token  string
}

type PushTokenDataService interface {
	Add(PushToken) error
	FindByUserId(string) []PushToken
	RemoveByToken(string) error
	RemoveByUserId(string) error
}
