package repository

import (
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/spf13/viper"
)

type RabbitC struct {
	uri string
}

func (r *RabbitC) PushPackage(data string) (bool, error) {
	conn, err := amqp.Dial(r.uri)
	FailOnError(err, "redis conn failed")
	channel, err := conn.Channel()
	FailOnError(err, "Failed to open a channel")
	defer conn.Close()
	defer channel.Close()
	q, err := channel.QueueDeclare(
		viper.GetString("Rabbit.TokenQueue"), // name
		false,                                // durable
		false,                                // delete when unused
		false,                                // exclusive
		false,                                // no-wait
		nil,                                  // arguments
	)

	FailOnError(err, "transport failed")
	err = channel.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(data),
		})

	FailOnError(err, "Failed to publish a message")

	return true, nil
}

func NewRabbitTransport(c *RabbitConfig) *RabbitC {
	uri := fmt.Sprintf("amqp://%s:%s@%s:%s/%s", c.Login, c.Password, c.Host, c.Port, c.VHost)
	return &RabbitC{uri}
}
