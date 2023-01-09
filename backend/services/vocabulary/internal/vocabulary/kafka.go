package vocabulary

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/segmentio/kafka-go"
)

type Consumer struct {
	service   *Service
	kafkaAddr string
}

func NewKafkaConsumer(service *Service, kafkaAddr string) *Consumer {
	return &Consumer{service, kafkaAddr}
}

func (c *Consumer) Serve(conn *kafka.Conn) error {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{c.kafkaAddr},
		Topic:    "vocabulary",
		GroupID:  "group-1",
		MaxBytes: 10e6, // 10MB
	})

	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			log.Printf("failed to read message: %v", err)
			break
		}

		var data WordMessage

		if err := json.Unmarshal(m.Value, &data); err != nil {
			log.Printf("failed to unmarshal message at offset %d: %v\n", m.Offset, err)
			continue
		}

		if data.Add {
			err = c.service.AddWord(context.Background(), data.UserId, data.WordId)
		} else {
			err = c.service.DeleteWord(context.Background(), data.UserId, data.WordId)
		}

		if err != nil {
			log.Printf("got error when processing word: %v", err)
		}
	}

	if err := conn.Close(); err != nil {
		return fmt.Errorf("conn.Close: %v", err)
	}

	return nil
}
