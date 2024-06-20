package db

import (
	"errors"
	"medication-notifier/data"
	"slices"
)

type DummyTokenDataService struct {
	tokens []data.Token
}

func NewDummyTokenDataService() DummyTokenDataService {
	return DummyTokenDataService{
		tokens: []data.Token{},
	}
}

func (s *DummyTokenDataService) Add(token data.Token) error {
	s.tokens = append(s.tokens, token)
	return nil
}

func (s *DummyTokenDataService) FindByToken(token string) (*data.Token, error) {
	idx := slices.IndexFunc(s.tokens, func(u data.Token) bool {
		return u.Token == token
	})
	if idx < 0 {
		return nil, errors.New("token_not_exists")
	}

	tokenCopy := s.tokens[idx]
	return &tokenCopy, nil
}

func (s *DummyTokenDataService) FindByClientId(clientId string) (*data.Token, error) {
	idx := slices.IndexFunc(s.tokens, func(u data.Token) bool {
		return u.ClientId == clientId
	})
	if idx < 0 {
		return nil, errors.New("token_not_exists")
	}

	tokenCopy := s.tokens[idx]
	return &tokenCopy, nil
}

func (s *DummyTokenDataService) FindByUserId(userId string) []data.Token {
	tokens := []data.Token{}
	for _, token := range s.tokens {
		if token.UserId == userId {
			tokens = append(tokens, token)
		}
	}

	return tokens
}

func (s *DummyTokenDataService) RemoveAllByUserId(userId string) error {
	_ = slices.DeleteFunc(s.tokens, func(token data.Token) bool {
		return token.UserId == userId
	})

	return nil
}

func (s *DummyTokenDataService) RemoveByToken(token string) error {
	_ = slices.DeleteFunc(s.tokens, func(t data.Token) bool {
		return t.Token == token
	})

	return nil
}

func (s *DummyTokenDataService) RemoveByClientId(clientId string) error {
	_ = slices.DeleteFunc(s.tokens, func(t data.Token) bool {
		return t.ClientId == clientId
	})

	return nil
}
