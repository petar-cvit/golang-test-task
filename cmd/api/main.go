package main

import (
	"github.com/gin-gonic/gin"
	"twitch_chat_analysis/internal/controller"
	"twitch_chat_analysis/internal/processor"
	"twitch_chat_analysis/internal/storage"
	"twitch_chat_analysis/internal/stream"
)

func main() {
	r := gin.Default()

	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, "worked")
	})

	messagesStore, err := storage.New()
	if err != nil {
		// panic because you would like the service to fail if this happens
		panic(err)
	}

	messagesStream, err := stream.NewRabbitMQ()
	if err != nil {
		panic(err)
	}

	ctrl := controller.New(messagesStream)

	messageProcessor := processor.New(messagesStore, messagesStream)

	// separate goroutine to process messages
	go messageProcessor.Start()

	r.POST("/message", ctrl.Receive)

	r.Run()
}
