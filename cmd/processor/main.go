package main

import (
	"twitch_chat_analysis/internal/processor"
	"twitch_chat_analysis/internal/storage"
	"twitch_chat_analysis/internal/stream"
)

func main() {
	messagesStore, err := storage.New()
	if err != nil {
		// panic because you would like the service to fail if this happens
		panic(err)
	}

	messagesStream, err := stream.NewRabbitMQ()
	if err != nil {
		panic(err)
	}

	messageProcessor := processor.New(messagesStore, messagesStream)

	messageProcessor.Start()
}
