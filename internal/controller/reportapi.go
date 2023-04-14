package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"twitch_chat_analysis/internal/storage"
)

type ReportController struct {
	messageStorage storage.Messages
}

func NewReport(messageStorage storage.Messages) ReportController {
	return ReportController{
		messageStorage: messageStorage,
	}
}

func (c *ReportController) ListMessages(ctx *gin.Context) {
	receiver := ctx.Query("receiver")
	sender := ctx.Query("sender")

	fmt.Println(receiver, sender)

	if len(sender) == 0 || len(receiver) == 0 {
		ctx.String(http.StatusBadRequest, "specify sender and receiver")
		return
	}

	messages, err := c.messageStorage.ListMessages(receiver, sender)
	if err != nil {
		fmt.Println("error fetching messages", err)
		ctx.Status(http.StatusBadRequest)
		return
	}

	ctx.JSON(http.StatusOK, messages)
	return
}
