package repository

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitC struct {
	channel *amqp.Channel
}

func (r *RabbitC) PushPackage(data string) (bool, error) {
	q, err := r.channel.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)

	FailOnError(err, "transport failed")

	// defer r.channel.Close()

	err = r.channel.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(data),
		})

	FailOnError(err, "Failed to publish a message")

	log.Printf(" [x] Sent %s\n", data)

	return true, nil
}

func NewRabbitTransport(c *amqp.Channel) *RabbitC {
	return &RabbitC{channel: c}
}
