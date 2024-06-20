package db

import (
	"errors"
	"medication-notifier/data"
	"slices"
	"strconv"
	"strings"
	"time"
)

type DummyUsersDataService struct {
	users []data.User
}

func NewDummyUsersDataService() DummyUsersDataService {
	return DummyUsersDataService{
		users: []data.User{},
	}
}

func (s *DummyUsersDataService) Add(username, passwordHash string) error {
	username = strings.ToLower(username)

	dbUser := s.fetchUserByUsername(username)
	if dbUser != nil {
		return errors.New("user_exists") // TODO move to enums
	}

	user := data.User{
		Id:           strconv.Itoa(len(s.users) + 1),
		Username:     username,
		PasswordHash: passwordHash,
		CreatedAt:    time.Now().UnixMicro(),
	}
	s.users = append(s.users, user)

	return nil
}

func (s *DummyUsersDataService) FindByUsername(username string) (*data.User, error) {
	username = strings.ToLower(username)
	user := s.fetchUserByUsername(username)
	if user == nil {
		return nil, errors.New("user_not_exists")
	}

	return user, nil
}

func (s *DummyUsersDataService) FindById(id string) (*data.User, error) {
	idx := slices.IndexFunc(s.users, func(u data.User) bool {
		return u.Id == id
	})
	if idx == 0 {
		return nil, errors.New("user_not_exists")
	}

	userCopy := s.users[idx]
	return &userCopy, nil
}

func (s *DummyUsersDataService) fetchUserByUsername(username string) *data.User {
	idx := slices.IndexFunc(s.users, func(u data.User) bool {
		return u.Username == username
	})
	if idx < 0 {
		return nil
	}

	userCopy := s.users[idx]
	return &userCopy
}
