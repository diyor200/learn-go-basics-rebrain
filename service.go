package main

import "errors"

type UserService struct {
	storage Storage
}

func NewUserService(storage Storage) *UserService {
	return &UserService{
		storage: storage,
	}
}

func (u *UserService) Login(username, password string) (*Session, error) {
	user, err := u.storage.GetUserByUsername(username)
	if err != nil {
		return nil, err
	}

	if user.Password != password {
		return nil, errors.New("wrong password")
	}

	return u.storage.CreateSession(user.ID)
}
