package data

type Token struct {
	UserId         string `json:"user_id"`
	Token          string `json:"token"`
	ExpirationTime int64  `json:"exp"`
	ClientInfo     string `json:"client_info"`
	ClientId       string `json:"client_id"`
}

type TokenDataService interface {
	Add(Token) error
	FindByToken(string, string) (*Token, error)
	FindByUserId(string) []Token
	RemoveAllByUserId(string) error
	RemoveByToken(string, string) error
}
