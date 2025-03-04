package main

type User struct {
	ID       string
	Username string
	Password string
}

type Session struct {
	ID     string
	UserID string
}

type Storage interface {
	GetUserByUsername(username string) (*User, error)
	CreateSession(userID string) (*Session, error)
}
