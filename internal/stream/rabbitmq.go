package stream

import (
	"context"
	"encoding/json"
	amqp "github.com/rabbitmq/amqp091-go"
	"time"

	"twitch_chat_analysis/internal/models"
)

type RabbitMQ struct {
	stream chan models.Message
	conn   *amqp.Connection
}

func NewRabbitMQ() (RabbitMQ, error) {
	conn, err := amqp.Dial("amqp://user:password@localhost:7001")
	if err != nil {
		return RabbitMQ{}, err
	}

	return RabbitMQ{
		stream: make(chan models.Message, 10),
		conn:   conn,
	}, nil
}

func (r RabbitMQ) Send(message models.Message) error {
	ch, err := r.conn.Channel()
	if err != nil {
		return err
	}

	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	body, err := json.Marshal(message)
	if err != nil {
		return err
	}

	if err := ch.PublishWithContext(ctx,
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        body,
		}); err != nil {
		return err
	}

	r.stream <- message
	return nil
}

func (r RabbitMQ) GetMessages() (<-chan amqp.Delivery, error) {
	ch, err := r.conn.Channel()
	if err != nil {
		return nil, err
	}

	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)

	msgs, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	return msgs, nil
}
