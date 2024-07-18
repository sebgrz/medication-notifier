package db

import (
	"context"
	"fmt"
	"medication-notifier/data"

	"github.com/jackc/pgx/v5/pgxpool"
)

type DbUsersDataService struct {
	conn *pgxpool.Pool
}

func NewDbUsersDataService(address string) DbUsersDataService {
	conn, err := pgxpool.New(context.Background(), address)
	if err != nil {
		panic(fmt.Sprintf("psql connection failed: %s", err))
	}

	return DbUsersDataService{
		conn: conn,
	}
}

func (s *DbUsersDataService) Add(username, passwordHash string, creationTime int64) error {
	sql := "insert into med.users(username, password, created_at) values ($1, $2, $3)"
	_, err := s.conn.Exec(context.Background(), sql, username, passwordHash, creationTime)
	return err
}

func (s *DbUsersDataService) FindByUsername(username string) (*data.User, error) {
	sql := "select id, username, password, created_at from med.users where username=$1"
	row := s.conn.QueryRow(context.Background(), sql, username)

	var id string
	var password string
	var creationTime int64
	if err := row.Scan(&id, &username, &password, &creationTime); err != nil {
		return nil, err
	}

	return &data.User{
		Id:           id,
		Username:     username,
		PasswordHash: password,
		CreatedAt:    creationTime,
	}, nil
}

func (s *DbUsersDataService) FindById(id string) (*data.User, error) {
	sql := "select id, username, created_at from med.users where id=$1"
	row := s.conn.QueryRow(context.Background(), sql, id)

	var username string
	var creationTime int64
	if err := row.Scan(&id, &username, &creationTime); err != nil {
		return nil, err
	}

	return &data.User{
		Id:           id,
		Username:     username,
		PasswordHash: "",
		CreatedAt:    creationTime,
	}, nil
}
