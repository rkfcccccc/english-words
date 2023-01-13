package dictionary

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/rkfcccccc/english_words/services/dictionary/pkg/dictionaryapi"
	"github.com/rkfcccccc/english_words/services/dictionary/pkg/lemmatizer"
	"github.com/rkfcccccc/english_words/shared_pkg/dsync"
	"github.com/segmentio/kafka-go"

	models "github.com/rkfcccccc/english_words/shared_pkg/services/dictionary/models"
)

type Service struct {
	repo       Repository
	sync       dsync.Client
	dict       dictionaryapi.Client
	lemmatizer *lemmatizer.Lemmatizer
	kafka      *kafka.Writer
}

var ErrNoDefinitionsFound = errors.New("no definitions found")

func NewService(repo Repository, sync dsync.Client, dict dictionaryapi.Client, lemm *lemmatizer.Lemmatizer, writer *kafka.Writer) *Service {
	return &Service{repo, sync, dict, lemm, writer}
}

func (service *Service) writeToPopulationQueue(wordId string) {
	if err := service.kafka.WriteMessages(context.Background(), kafka.Message{Value: []byte(wordId)}); err != nil {
		log.Printf("failed writing to population queue: %v", err)
	}
}

func (service *Service) Create(ctx context.Context, word string) (string, error) {
	word = service.lemmatizer.Lemma(word)

	mutex := service.sync.NewMutex(fmt.Sprintf("dictionary_%s", word))
	if err := mutex.Lock(); err != nil {
		return "", fmt.Errorf("mutex.Lock: %v", err)
	}

	defer mutex.Unlock()

	entry, err := service.repo.GetByWord(ctx, word)
	if err != nil {
		return "", fmt.Errorf("repo.GetByWord: %v", err)
	}

	if entry != nil {
		if entry.Pictures == nil {
			go service.writeToPopulationQueue(entry.Id)
		}

		return entry.Id, nil
	}

	dEntry, err := service.dict.GetWordEntry(ctx, "en", word)
	if errors.Is(err, dictionaryapi.ErrNoDefinitionsFound) {
		return "", ErrNoDefinitionsFound
	}

	if err != nil {
		return "", fmt.Errorf("dict.GetWordEntry: %v", err)
	}

	wordId, err := service.repo.Create(ctx, TransformFromApi(dEntry))
	if err != nil {
		return "", fmt.Errorf("repo.Create: %v", err)
	}

	go service.writeToPopulationQueue(wordId)
	return wordId, nil
}

func (service *Service) GetByWord(ctx context.Context, word string) (*models.WordEntry, error) {
	word = service.lemmatizer.Lemma(word)
	return service.repo.GetByWord(ctx, word)
}

func (service *Service) GetById(ctx context.Context, wordId string) (*models.WordEntry, error) {
	return service.repo.GetById(ctx, wordId)
}

func (service *Service) Delete(ctx context.Context, wordId string) error {
	return service.repo.Delete(ctx, wordId)
}

func (service *Service) SetPictures(ctx context.Context, wordId string, pictures []models.SourcedPicture) error {
	return service.repo.SetPictures(ctx, wordId, pictures)
}

func (service *Service) Search(ctx context.Context, query string) ([]*models.WordEntry, error) {
	return service.repo.Search(ctx, query)
}
