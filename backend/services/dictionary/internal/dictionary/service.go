package dictionary

import (
	"context"
	"errors"
	"fmt"

	"github.com/rkfcccccc/english_words/services/dictionary/pkg/dictionaryapi"
	"github.com/rkfcccccc/english_words/services/dictionary/pkg/lemmatizer"
	"github.com/rkfcccccc/english_words/shared_pkg/dsync"
)

type Service struct {
	repo       Repository
	sync       dsync.Client
	dict       dictionaryapi.Client
	lemmatizer *lemmatizer.Lemmatizer
}

var ErrNoDefinitionsFound = errors.New("no definitions found")

func NewService(repo Repository, sync dsync.Client, dict dictionaryapi.Client, lemm *lemmatizer.Lemmatizer) *Service {
	return &Service{repo, sync, dict, lemm}
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

	return wordId, nil
}

func (service *Service) GetByWord(ctx context.Context, word string) (*WordEntry, error) {
	word = service.lemmatizer.Lemma(word)
	return service.repo.GetByWord(ctx, word)
}

func (service *Service) GetById(ctx context.Context, wordId string) (*WordEntry, error) {
	return service.repo.GetById(ctx, wordId)
}

func (service *Service) Delete(ctx context.Context, wordId string) error {
	return service.repo.Delete(ctx, wordId)
}

func (service *Service) SetPictures(ctx context.Context, wordId string, pictures []SourcedPicture) error {
	return service.repo.SetPictures(ctx, wordId, pictures)
}
