package processor

import (
	"encoding/json"
	"fmt"
	"twitch_chat_analysis/internal/models"
	"twitch_chat_analysis/internal/storage"
	"twitch_chat_analysis/internal/stream"
)

type Processor struct {
	messagesStorage storage.Messages
	stream          stream.Receiver
}

func New(messagesStorage storage.Messages, stream stream.Receiver) Processor {
	return Processor{
		messagesStorage: messagesStorage,
		stream:          stream,
	}
}

func (p Processor) Start() {
	fmt.Println("started processing messages...")

	messagesStream, err := p.stream.GetMessages()
	if err != nil {
		panic(err)
	}

	for {
		select {
		case m := <-messagesStream:
			var message models.Message
			if err := json.Unmarshal(m.Body, &message); err != nil {
				fmt.Print(err)
				continue
			}

			if err := p.messagesStorage.StoreMessage(message); err != nil {
				fmt.Println(err)
			}
		}
	}
}
