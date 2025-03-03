package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type fakeStorage struct {
	createReturns  fakeStorageReturns
	recoverReturns fakeStorageReturns
}

type fakeStorageReturns struct {
	Session *Session
	Err     error
}

func (fs *fakeStorage) RecoverSession(id string) (*Session, error) {
	return fs.recoverReturns.Session, fs.recoverReturns.Err
}

func (fs *fakeStorage) CreateSession(name, pass string) (*Session, error) {
	return fs.createReturns.Session, fs.createReturns.Err
}

func TestInit(t *testing.T) {
	t.Run("create sesion success", func(t *testing.T) {
		req := require.New(t)
		fs := fakeStorage{
			createReturns:  fakeStorageReturns{Session: &Session{ID: "session id", UserID: "user id"}, Err: nil},
			recoverReturns: fakeStorageReturns{Session: nil, Err: nil},
		}

		session, err := InitSession(User{Name: "Alex", Pass: "pass"}, &fs)
		req.NoError(err)
		req.Equal("session id", session.UserID)
	})
}
