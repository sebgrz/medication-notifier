package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type DbActivityLogDataService struct {
	conn *pgxpool.Pool
}

func NewDbActivityLogDataService(conn *pgxpool.Pool) DbActivityLogDataService {
	return DbActivityLogDataService{
		conn: conn,
	}
}

func (s *DbActivityLogDataService) Add(clientId, userId, refreshTokenHash string, expiryTime int64) error {
	sql := "insert into med.login_activity_log(client_id, user_id, refresh_token_hash, expire_time) values ($1, $2, $3, $4)"
	_, err := s.conn.Exec(
		context.Background(),
		sql,
		clientId,
		userId,
		refreshTokenHash,
		expiryTime,
	)
	return err
}
