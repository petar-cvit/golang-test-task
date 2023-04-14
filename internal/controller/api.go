package controller

import (
	"fmt"
	"net/http"
	"twitch_chat_analysis/internal/models"
	"twitch_chat_analysis/internal/stream"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	stream stream.Producer
}

func New(stream stream.Producer) Controller {
	return Controller{
		stream: stream,
	}
}

func (c *Controller) Receive(ctx *gin.Context) {
	var message models.Message

	if err := ctx.BindJSON(&message); err != nil {
		fmt.Println("error binding request", err)
		ctx.Status(http.StatusBadRequest)
		return
	}

	fmt.Println(message)

	if err := c.stream.Send(message); err != nil {
		fmt.Println("error sending message", err)
		ctx.Status(http.StatusBadRequest)
		return
	}

	ctx.Status(http.StatusOK)
	return
}
