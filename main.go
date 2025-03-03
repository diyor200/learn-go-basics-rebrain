package main

type User struct {
	SessionID string
	Name      string
	Pass      string
}

type Session struct {
	ID     string
	UserID string
}

type Storage interface {
	CreateSession(name, pass string) (*Session, error)
	RecoverSession(id string) (*Session, error)
}

func InitSession(user User, storage Storage) (*Session, error) {
	session, err := storage.RecoverSession(user.SessionID)
	if err != nil {
		return nil, err
	}

	if session == nil {
		return storage.CreateSession(user.Name, user.Pass)
	}

	return session, nil
}
