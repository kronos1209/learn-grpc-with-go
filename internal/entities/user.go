package entities

import (
	"crypto/sha256"
	"errors"
)

type User struct {
	ID       int64
	UserID   string
	Name     string
	Password string
}

var (
	ErrInvalidArgument = errors.New("invalid arguments")
)

func NewUser(userId string, name string, password string) (*User, error) {
	if userId == "" && name == "" && password == "" {
		return nil, ErrInvalidArgument
	}
	encodedPasswd := sha256.Sum256([]byte(password))
	u := User{
		UserID:   userId,
		Name:     name,
		Password: string(encodedPasswd[:]),
	}
	return &u, nil
}
