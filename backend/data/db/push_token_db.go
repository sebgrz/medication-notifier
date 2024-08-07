package db

import (
	"context"
	"medication-notifier/data"
	"medication-notifier/utils/logger"

	"github.com/jackc/pgx/v5/pgxpool"
)

type DbPushTokenDataService struct {
	conn *pgxpool.Pool
}

func NewDbPushTokenDataService(conn *pgxpool.Pool) DbPushTokenDataService {
	return DbPushTokenDataService{
		conn: conn,
	}
}

func (s *DbPushTokenDataService) Add(pushToken data.PushToken) error {
	sql := "insert into med.push_token(user_id, token) values ($1, $2)"
	_, err := s.conn.Exec(
		context.Background(),
		sql,
		pushToken.UserId,
		pushToken.Token,
	)
	return err
}

func (s *DbPushTokenDataService) FindByUserId(userId string) []data.PushToken {
	sql := `
	select id, user_id, token from med.push_token
	where user_id=$1
	`
	result := []data.PushToken{}
	rows, err := s.conn.Query(context.Background(), sql, userId)
	if err != nil {
		logger.Error("fetch push_token by userId failed, err: %s", err)
		return result
	}

	for rows.Next() {
		var id string
		var userId string
		var token string
		if err := rows.Scan(&id, &userId, &token); err != nil {
			logger.Error("fetch push_token by userId scan failed, err: %s", err)
			return []data.PushToken{}
		}
		result = append(result, data.PushToken{
			Id:     id,
			UserId: userId,
			Token:  token,
		})
	}

	return result
}

func (s *DbPushTokenDataService) RemoveByToken(token string) error {
	sql := "delete from med.push_token where token=$1"
	_, err := s.conn.Exec(context.Background(), sql, token)

	return err
}

func (s *DbPushTokenDataService) RemoveByUserId(userId string) error {
	sql := "delete from med.push_token where user_id=$1"
	_, err := s.conn.Exec(context.Background(), sql, userId)

	return err
}
