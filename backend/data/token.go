package data

type Token struct {
	UserId         string
	Token          string
	ExpirationTime int64
	ClientInfo     string
	ClientId       string
}

type TokenDataService interface {
	Add(Token) error
	FindByToken(string) (*Token, error)
	FindByClientId(string) (*Token, error)
	FindByUserId(string) []Token
	RemoveAllByUserId(string) error
	RemoveByToken(string) error
	RemoveByClientId(string) error
}
