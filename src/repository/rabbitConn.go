package repository

import (
	"log"
)

type RabbitConfig struct {
	Login    string
	Password string
	Port     string
	Host     string
	VHost    string
}

func FailOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func NewRabbitConn(c RabbitConfig) (*RabbitConfig, error) {
	return &RabbitConfig{
		Login:    c.Login,
		Password: c.Password,
		Port:     c.Port,
		Host:     c.Host,
		VHost:    c.VHost,
	}, nil
}
