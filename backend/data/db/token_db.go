package db

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"medication-notifier/data"
	"time"

	redis "github.com/redis/go-redis/v9"
)

const TOKEN_KEY = "token_%s:client_id_%s"
const USER_ID_KEY = "user_id_%s"

type DbTokenDataService struct {
	client *redis.Client
}

func NewDbTokenDataService(address, password string) DbTokenDataService {
	client := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: password,
		DB:       0,
	})

	return DbTokenDataService{
		client,
	}
}

func (s *DbTokenDataService) Add(token data.Token) error {
	ctx := context.Background()
	key := fmt.Sprintf(TOKEN_KEY, token.Token, token.ClientId)
	jsonToken, _ := json.Marshal(token)
	err := s.client.SetArgs(ctx, key, string(jsonToken), redis.SetArgs{
		Mode:     "",
		KeepTTL:  false,
		Get:      false,
		TTL:      0,
		ExpireAt: time.Unix(token.ExpirationTime, 0),
	}).Err()

	if err != nil {
		fmt.Printf("add_token err: %s\n", err)
		return err
	}

	// TODO: save somehow user tokens list
	return nil
}

func (s *DbTokenDataService) FindByToken(token string, clientId string) (*data.Token, error) {
	ctx := context.Background()
	key := fmt.Sprintf(TOKEN_KEY, token, clientId)
	cmd := s.client.Get(ctx, key)
	if cmd.Err() == redis.Nil {
		return nil, errors.New("not found")
	}

	if cmd.Err() != nil {
		return nil, cmd.Err()
	}

	var tokenData *data.Token
	if err := json.Unmarshal([]byte(cmd.Val()), tokenData); err != nil {
		return nil, err
	}

	return tokenData, nil
}
func (s *DbTokenDataService) FindByUserId(string) []data.Token {
	panic("UNIMPLEMENTED")
}

func (s *DbTokenDataService) RemoveAllByUserId(string) error {
	panic("UNIMPLEMENTED")
}

func (s *DbTokenDataService) RemoveByToken(token string, clientId string) error {
	ctx := context.Background()
	key := fmt.Sprintf(TOKEN_KEY, token, clientId)
	s.client.Del(ctx, key)

	return nil
}
