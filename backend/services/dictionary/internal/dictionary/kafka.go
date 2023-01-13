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
	service   *Service
	pictures  *pictureapi.Service
	kafkaAddr string
}

func NewKafkaConsumer(service *Service, pictures *pictureapi.Service, kafkaAddr string) *Consumer {
	return &Consumer{service, pictures, kafkaAddr}
}

func NewKafkaProducer(kafkaAddr string) *kafka.Writer {
	return &kafka.Writer{Addr: kafka.TCP(kafkaAddr), Topic: "dictionary"}
}

func (c *Consumer) populateWithPictures(ctx context.Context, entry *models.WordEntry) {
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

	if err := c.service.SetPictures(ctx, entry.Id, sourced); err != nil {
		log.Printf("got error when processing word: %v", err)
	}

	log.Printf("populated %s with %d pictures\n", entry.Word, count)
}

func (c *Consumer) Serve(conn *kafka.Conn) error {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{c.kafkaAddr},
		Topic:    "dictionary",
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

		entry, err := c.service.GetById(ctx, string(m.Value))
		if err != nil {
			log.Printf("service.GetById: %v", err)
			cancel()
			continue
		}

		c.populateWithPictures(ctx, entry)
		cancel()
	}

	if err := conn.Close(); err != nil {
		return fmt.Errorf("conn.Close: %v", err)
	}

	return nil
}
