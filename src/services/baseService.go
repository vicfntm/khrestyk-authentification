package services

import (
	"example.com/hello/src/models"
	"example.com/hello/src/repository"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
	LoginUser(user models.LoginUserStruct) (string, int64, error)
	ParseToken(token string) (int, error)
}

type Service struct {
	Authorization
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos),
	}
}
