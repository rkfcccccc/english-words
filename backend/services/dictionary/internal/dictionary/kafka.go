package dictionary

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"

	"github.com/rkfcccccc/english_words/services/dictionary/pkg/pictureapi"
	"github.com/rkfcccccc/english_words/services/dictionary/pkg/reversoapi"
	models "github.com/rkfcccccc/english_words/shared_pkg/services/dictionary/models"
	"github.com/segmentio/kafka-go"
	"golang.org/x/sync/semaphore"
)

const MAX_CONCURRENT = 10
const MAX_PICTURES = 10

type Consumer struct {
	service   *Service
	pictures  *pictureapi.Service
	reverso   reversoapi.Client
	kafkaAddr string
}

func NewKafkaConsumer(service *Service, pictures *pictureapi.Service, reverso reversoapi.Client, kafkaAddr string) *Consumer {
	return &Consumer{service, pictures, reverso, kafkaAddr}
}

func NewKafkaProducer(kafkaAddr string) *kafka.Writer {
	return &kafka.Writer{Addr: kafka.TCP(kafkaAddr), Topic: "dictionary"}
}

func (c *Consumer) populateWithPictures(ctx context.Context, entry *models.WordEntry) {
	if entry.Pictures != nil {
		return
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

	if err := c.service.SetPictures(ctx, entry.Id, sourced); err != nil {
		log.Printf("got error when service.SetPictures: %v", err)
	}

	log.Printf("populated %s with %d pictures\n", entry.Word, count)
}

func (c *Consumer) populateWithTranslations(ctx context.Context, entry *models.WordEntry) {
	if entry.Translations != nil {
		return
	}

	result, err := c.reverso.GetTranslation(ctx, entry.Word, "eng", "rus")
	if err != nil {
		log.Printf("got error when reverso.GetTranslation of %s: %v", entry.Word, err)
		return
	}

	translations := make([]string, len(result.ContextResults.Results))
	for i, result := range result.ContextResults.Results {
		translations[i] = result.Translation
	}

	if err := c.service.SetTranslations(ctx, entry.Id, translations); err != nil {
		log.Printf("got error when service.SetTranslations: %v", err)
	}

	log.Printf("populated %s with %d translations\n", entry.Word, len(translations))
}

func (c *Consumer) populateAll(ctx context.Context, entry *models.WordEntry) {
	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		c.populateWithPictures(ctx, entry)
		wg.Done()
	}()

	go func() {
		c.populateWithTranslations(ctx, entry)
		wg.Done()
	}()

	wg.Wait()
}

func (c *Consumer) processWord(wordId string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	entry, err := c.service.GetById(ctx, wordId)
	if err != nil {
		log.Printf("service.GetById: %v", err)
		cancel()
		return
	}

	c.populateAll(ctx, entry)
	cancel()
}

func (c *Consumer) Serve(conn *kafka.Conn) error {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{c.kafkaAddr},
		Topic:    "dictionary",
		GroupID:  "group-1",
		MaxBytes: 10e6, // 10MB
	})

	sem := semaphore.NewWeighted(20)

	for {
		message, err := r.ReadMessage(context.Background())
		if err != nil {
			log.Printf("failed to read message: %v", err)
			break
		}

		if err := sem.Acquire(context.Background(), 1); err != nil {
			log.Printf("failed to acquire semaphore: %v", err)
			break
		}

		c.processWord(string(message.Value))
		sem.Release(1)
	}

	if err := conn.Close(); err != nil {
		return fmt.Errorf("conn.Close: %v", err)
	}

	return nil
}
