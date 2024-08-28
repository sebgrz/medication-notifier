package db

import (
	"context"

	"medication-notifier/data"
	"medication-notifier/utils/logger"

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

func (s *DbActivityLogDataService) FetchAllNewestGroupedByClientId() []data.ActivityLog {
	sql := `
	select
		client_id,
		user_id,
		refresh_token_hash,
		expire_time
	from
	(
		select
			client_id,
			user_id,
			refresh_token_hash,
			expire_time,
			row_number() over
				(partition by user_id order by expire_time desc)
				as row_number
		from
			med.login_activity_log
	) data
	where
		data.row_number = 1
	`
	result := []data.ActivityLog{}
	rows, err := s.conn.Query(context.Background(), sql)
	if err != nil {
		logger.Error("fetch activity_log failed, err: %s", err)
		return result
	}

	for rows.Next() {
		var clientId string
		var userId string
		var refreshTokenHash string
		var expireTime int64
		if err := rows.Scan(&clientId, &userId, &refreshTokenHash, &expireTime); err != nil {
			logger.Error("fetch activity_log scan failed, err: %s", err)
			return []data.ActivityLog{}
		}
		result = append(result, data.ActivityLog{
			ClientId:         clientId,
			UserId:           userId,
			RefreshTokenHash: refreshTokenHash,
			ExpireTime:       expireTime,
		})
	}

	return result
}
