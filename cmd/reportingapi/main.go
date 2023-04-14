package main

import (
	"github.com/gin-gonic/gin"
	"twitch_chat_analysis/internal/controller"
	"twitch_chat_analysis/internal/storage"
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

	ctrl := controller.NewReport(messagesStore)

	r.GET("/message/list", ctrl.ListMessages)

	r.Run()
}
