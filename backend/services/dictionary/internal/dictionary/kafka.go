package dictionary

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/rkfcccccc/english_words/services/dictionary/pkg/pictureapi"
	models "github.com/rkfcccccc/english_words/shared_pkg/services/dictionary/models"
	"github.com/segmentio/kafka-go"
)

const MAX_PICTURES = 10

type Consumer struct {
	service  *Service
	pictures *pictureapi.Service
}

func NewKafkaConsumer(service *Service, pictures *pictureapi.Service) *Consumer {
	return &Consumer{service, pictures}
}

func NewKafkaProducer() *kafka.Writer {
	return &kafka.Writer{Addr: kafka.TCP("localhost:9092"), Topic: "pictures"}
}

func (c *Consumer) Serve(conn *kafka.Conn) error {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{"localhost:9092"},
		Topic:    "pictures",
		GroupID:  "group-1",
		MaxBytes: 10e6, // 10MB
	})

	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			log.Printf("failed to read message: %v", err)
			break
		}

		ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
		wordId := string(m.Value)
		entry, err := c.service.GetById(ctx, wordId)
		if err != nil {
			log.Printf("service.GetById: %v", err)
			cancel()
			continue
		}

		pictures := c.pictures.Search(ctx, entry.Word)
		rand.Shuffle(len(pictures), func(i, j int) {
			pictures[i], pictures[j] = pictures[j], pictures[i]
		})

		count := MAX_PICTURES
		if len(pictures) < count {
			count = len(pictures)
		}

		sourced := make([]models.SourcedPicture, count)
		for i := 0; i < count; i++ {
			sourced[i] = models.SourcedPicture{
				Url:    pictures[i].Url,
				Source: pictures[i].Source,
			}
		}

		if err := c.service.SetPictures(ctx, wordId, sourced); err != nil {
			log.Printf("got error when processing word: %v", err)
		}

		fmt.Printf("populated %s with %d pictures\n", entry.Word, count)
		cancel()
	}

	if err := conn.Close(); err != nil {
		return fmt.Errorf("conn.Close: %v", err)
	}

	return nil
}
