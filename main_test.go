package main

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	gomock "go.uber.org/mock/gomock"
)

func TestInitSession(t *testing.T) {
	any := gomock.Any()

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	storageMock := NewMockStorage(mockCtrl)

	t.Run("create session success", func(t *testing.T) {
		storageMock.EXPECT().RecoverSession(any).Return(nil, nil).Times(1)
		storageMock.EXPECT().CreateSession(any, any).Return(&Session{ID: "session_id", UserID: 1}, nil).Times(1)

		session, err := InitSession(User{Name: "ivan", Pass: "secret"}, storageMock)
		assert.Equal(t, nil, err)
		assert.Equal(t, "session_id", session.ID)
	})

	t.Run("create session fail", func(t *testing.T) {
		storageMock.EXPECT().RecoverSession(any).Return(nil, nil).Times(1)
		storageMock.EXPECT().CreateSession(any, any).Return(nil, errors.New("fail to create session")).Times(1)

		_, err := InitSession(User{Name: "ivan", Pass: "secret"}, storageMock)
		assert.Equal(t, errors.New("fail to create session"), err)
	})

	t.Run("recover session success", func(t *testing.T) {
		storageMock.EXPECT().RecoverSession(any).Return(&Session{ID: "session_id", UserID: 1}, nil).Times(1)

		session, err := InitSession(User{SessionID: "session_id", Name: "ivan", Pass: "secret"}, storageMock)
		assert.Equal(t, nil, err)
		assert.Equal(t, "session_id", session.ID)
	})
}
