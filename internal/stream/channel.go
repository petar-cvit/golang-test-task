package stream

import (
	"twitch_chat_analysis/internal/models"
)

type Channel struct {
	stream chan models.Message
}

func NewChannel() Channel {
	return Channel{
		stream: make(chan models.Message, 10),
	}
}

func (c Channel) Send(message models.Message) error {
	c.stream <- message
	return nil
}

func (c Channel) GetMessages() chan models.Message {
	return c.stream
}
