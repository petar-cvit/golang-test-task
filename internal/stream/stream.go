package stream

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"twitch_chat_analysis/internal/models"
)

type Producer interface {
	Send(message models.Message) error
}

type Receiver interface {
	GetMessages() (<-chan amqp.Delivery, error)
}
