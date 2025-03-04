package main

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

//func TestInitSession(t *testing.T) {
//	type mockBehavior func(s *MockStorage, user User)
//
//	testTable := []struct {
//		name            string
//		inputUser       User
//		mockBehavior    mockBehavior
//		expectedSession *Session
//		expectedError   error
//	}{
//		{
//			name: "Create session: OK",
//			inputUser: User{
//				Name: "Alex Sparrow",
//				Pass: "pass",
//			},
//			mockBehavior: func(s *MockStorage, user User) {
//				s.EXPECT().RecoverSession(user.SessionID).Return(nil, nil)
//
//				s.EXPECT().CreateSession(user.Name, user.Pass).Return(&Session{ID: "sessionID", UserID: "userID"}, nil)
//			},
//			expectedSession: &Session{ID: "sessionID", UserID: "userID"},
//			expectedError:   nil,
//		},
//		{
//			name: "Create session: Fail",
//			inputUser: User{
//				Name: "Alex Sparrow",
//				Pass: "alex",
//			},
//			mockBehavior: func(s *MockStorage, user User) {
//				s.EXPECT().RecoverSession(user.SessionID).Return(nil, nil)
//
//				s.EXPECT().CreateSession(user.Name, user.Pass).Return(nil, errors.New("fail"))
//			},
//			expectedSession: nil,
//			expectedError:   errors.New("fail"),
//		},
//	}
//
//	for _, tt := range testTable {
//		t.Run(tt.name, func(t *testing.T) {
//			c := gomock.NewController(t)
//			defer c.Finish()
//
//			storage := NewMockStorage(c)
//			tt.mockBehavior(storage, tt.inputUser)
//
//			session, err := InitSession(tt.inputUser, storage)
//
//			// check
//			assert.Equal(t, tt.expectedSession, session)
//			assert.Equal(t, tt.expectedError, err)
//		})
//
//	}
//}

func TestUserService_Login(t *testing.T) {
	//ctrl := gomock.NewController(t)
	//defer ctrl.Finish()
	//
	//mockStorage := NewMockStorage(ctrl)
	//userService := NewUserService(mockStorage)

	testCases := []struct {
		name            string
		user            User
		mockBehavior    func(s *MockStorage)
		expectedSession *Session
		expectedError   error
	}{
		{
			name: "Success login",
			user: User{Username: "Alex", Password: "pass"},
			mockBehavior: func(s *MockStorage) {
				s.EXPECT().GetUserByUsername("Alex").Return(&User{Username: "Alex", Password: "pass", ID: "1"}, nil)

				s.EXPECT().CreateSession("1").Return(&Session{ID: "sessionID", UserID: "1"}, nil)
			},
			expectedSession: &Session{ID: "sessionID", UserID: "1"},
			expectedError:   nil,
		},
		{
			name: "Invalid password",
			user: User{Username: "Alex", Password: "pass"},
			mockBehavior: func(s *MockStorage) {
				s.EXPECT().GetUserByUsername("Alex").Return(&User{Username: "Alex", Password: "123"}, nil)
			},
			expectedSession: nil,
			expectedError:   errors.New("wrong password"),
		},
		{
			name: "user not found",
			user: User{Username: "Felix", Password: "passFelix"},
			mockBehavior: func(s *MockStorage) {
				s.EXPECT().GetUserByUsername("Felix").Return(nil, errors.New("user not found"))
			},
			expectedSession: nil,
			expectedError:   errors.New("user not found"),
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockStorage := NewMockStorage(ctrl)
			mockUserService := NewUserService(mockStorage)

			tt.mockBehavior(mockStorage)

			session, err := mockUserService.Login(tt.user.Username, tt.user.Password)

			// check
			assert.Equal(t, tt.expectedSession, session)
			assert.Equal(t, tt.expectedError, err)
		})
	}
}
