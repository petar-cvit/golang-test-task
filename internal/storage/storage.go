package storage

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	"time"
	"twitch_chat_analysis/internal/models"
)

type Messages struct {
	client *redis.Client
}

func New() (Messages, error) {
	client := redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
		DB:   0,
	})

	if err := client.Ping().Err(); err != nil {
		return Messages{}, err
	}

	return Messages{
		client: client,
	}, nil
}

func (m Messages) StoreMessage(message models.Message) error {
	message.Timestamp = time.Now().Unix()

	bytes, err := json.Marshal(message)
	if err != nil {
		return err
	}

	return m.client.ZAdd(messagesKey(message.Sender, message.Receiver), redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: bytes,
	}).Err()
}

func (m Messages) ListMessages(receiver, sender string) ([]models.Message, error) {
	messagesString, err := m.client.ZRangeByScore(messagesKey(sender, receiver), redis.ZRangeBy{
		Min: "-inf",
		Max: "+inf",
	}).Result()
	if err != nil {
		return nil, err
	}

	messages := make([]models.Message, 0, len(messagesString))

	for _, messageString := range messagesString {
		var message models.Message
		if err := json.Unmarshal([]byte(messageString), &message); err != nil {
			return nil, err
		}

		messages = append(messages, message)
	}

	return messages, nil
}

func messagesKey(sender, receiver string) string {
	return fmt.Sprintf("messages:%s:%s", sender, receiver)
}
