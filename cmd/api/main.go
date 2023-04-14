package main

import (
	"github.com/gin-gonic/gin"
	"twitch_chat_analysis/internal/controller"
	"twitch_chat_analysis/internal/stream"
)

func main() {
	r := gin.Default()

	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, "worked")
	})

	messagesStream, err := stream.NewRabbitMQ()
	if err != nil {
		panic(err)
	}

	ctrl := controller.New(messagesStream)

	r.POST("/message", ctrl.Receive)

	r.Run()
}
