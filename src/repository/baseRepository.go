package repository

import (
	"example.com/hello/src/models"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
	GetUserFromDb(login string, password string) (models.LoginUserStruct, error)
}

type RabbitTransport interface {
	PushPackage(data string) (bool, error)
}

type Repository struct {
	Authorization
	RabbitTransport
}

func NewRepository(db *sqlx.DB, rc *RabbitConfig) *Repository {
	return &Repository{
		Authorization:   NewAuthPostgres(db),
		RabbitTransport: NewRabbitTransport(rc),
	}
}
