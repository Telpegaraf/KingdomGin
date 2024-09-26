package auth

import (
	"kingdom/model"
	"time"
)

type Database interface {
	GetUserByName(name string) (*model.User, error)
	GetUserByID(id uint) (*model.User, error)
	UpdateClientTokensLastUsed(tokens []string, t *time.Time) error
	UpdateApplicationTokenLastUsed(token string, t *time.Time) error
}

type Auth struct {
	DB Database
}
