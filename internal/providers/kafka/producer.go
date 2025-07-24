package kafka

import (
	"context"
	"encoding/json"
	"github.com/airo507/GoProjectCore/internal/config"
	"github.com/airo507/GoProjectCore/internal/entity/comment"
	"github.com/segmentio/kafka-go"
	"log/slog"
)

type KafkaJson struct {
	Event  string `json:"event"`
	UserId int    `json:"user_id"`
	Body   string `json:"body"`
}

type KafkaProducer struct {
	producer *kafka.Conn
	logger   *slog.Logger
}

func NewProducer(ctx context.Context, kafkaConfig config.KafkaConfig, logger *slog.Logger) *KafkaProducer {
	conn, err := kafka.DialLeader(ctx, "tcp", kafkaConfig.KafkaAddr, kafkaConfig.Topic, int(kafkaConfig.Partition))
	if err != nil {
		return nil
	}

	defer conn.Close()

	return &KafkaProducer{
		producer: conn,
		logger:   logger,
	}
}

func (k *KafkaProducer) WriteCommentMessage(ctx context.Context, comment comment.Message) error {

	messageValue := KafkaJson{
		Event:  "message in blog created",
		UserId: comment.Author,
		Body:   comment.Body,
	}

	bytes, err := json.Marshal(messageValue)
	if err != nil {
		return err
	}

	message := kafka.Message{
		Key:   []byte("user"),
		Value: bytes,
	}

	_, err = k.producer.WriteMessages(message)
	if err != nil {
		k.logger.Error("Failed to write message: %v", err)
	}

	k.logger.Debug("Message sent successfully!", err)
	return nil

}
