package main

type User struct {
	SessionID string
	Name      string
	Pass      string
}

type Session struct {
	ID     string
	UserID int
}

type Storage interface {
	RecoverSession(id string) (*Session, error)
	CreateSession(name string, pass string) (*Session, error)
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

func main() {
}
