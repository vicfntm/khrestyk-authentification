package repository

import (
	"fmt"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitConfig struct {
	Login    string
	Password string
	Port     string
	Host     string
	VHost    string
}

func RabbitConnect(c RabbitConfig) (*amqp.Channel, error) {
	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%s/%s", c.Login, c.Password, c.Host, c.Port, c.VHost))

	FailOnError(err, "redis conn failed")
	ch, err := conn.Channel()
	FailOnError(err, "Failed to open a channel")
	// defer conn.Close()
	// defer ch.Close()
	return ch, nil
}

func FailOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}
