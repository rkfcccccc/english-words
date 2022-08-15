package movie

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/rkfcccccc/english_words/services/movie/pkg/srt"
	"github.com/rkfcccccc/english_words/shared_pkg/services/dictionary"
	"github.com/rkfcccccc/english_words/shared_pkg/services/vocabulary"
)

type Service struct {
	repo Repository

	dictionary *dictionary.Client
	vocabulary *vocabulary.Client
}

func NewService(repo Repository, dictionary *dictionary.Client, vocabulary *vocabulary.Client) *Service {
	return &Service{repo, dictionary, vocabulary}
}

func (service *Service) CreateByUrl(ctx context.Context, movie *Movie, subtitlesUrl string) (int, error) {
	log.Printf("creating movie \"%s\"..\n", movie.Title)

	words, err := srt.GetWordsFromURL(subtitlesUrl)
	if err != nil {
		return -1, fmt.Errorf("srt.GetWordsFromURL: %v", err)
	}

	log.Printf("found %d words in that movie\n", len(words))
	seen := make(map[string]struct{}, len(words))
	for _, word := range words {
		wordId, err := service.dictionary.Create(ctx, word)

		if errors.Is(err, dictionary.ErrNoDefinitionsFound) {
			continue
		} else if err != nil {
			return -1, fmt.Errorf("dict.Create: %v", err)
		}

		log.Printf("%s -> %s\n", word, wordId)
		seen[wordId] = struct{}{}
	}

	log.Printf("got %d words ids\n", len(seen))

	i := 0
	wordsIds := make([]string, len(seen))
	for wordId := range seen {
		wordsIds[i] = wordId
		i++
	}

	return service.repo.Create(ctx, movie, wordsIds)
}

func (service *Service) Create(ctx context.Context, movie *Movie, wordsIds []string) (int, error) {
	return service.repo.Create(ctx, movie, wordsIds)
}

func (service *Service) Delete(ctx context.Context, movieId int) error {
	return service.repo.Delete(ctx, movieId)
}

func (service *Service) Get(ctx context.Context, movieId int) (*Movie, error) {
	return service.repo.Get(ctx, movieId)
}

func (service *Service) GetWords(ctx context.Context, movieId int) ([]string, error) {
	return service.repo.GetWords(ctx, movieId)
}

func (service *Service) AddUser(ctx context.Context, movieId int, userId int) error {
	if fav, err := service.repo.IsFavorite(ctx, movieId, userId); err != nil {
		return fmt.Errorf("repo.IsFavorite: %v", err)
	} else if fav {
		return errors.New("already favorite")
	}

	wordsIds, err := service.GetWords(ctx, movieId)
	if err != nil {
		return fmt.Errorf("service.GetWords: %v", err)
	}

	if err := service.vocabulary.AddWords(ctx, userId, wordsIds...); err != nil {
		return fmt.Errorf("vocabulary.AddWords: %v", err)
	}

	if err := service.repo.AddUser(ctx, movieId, userId); err != nil {
		return fmt.Errorf("repo.AddUser: %v", err)
	}

	return nil
}

func (service *Service) RemoveUser(ctx context.Context, movieId int, userId int) error {
	if fav, err := service.repo.IsFavorite(ctx, movieId, userId); err != nil {
		return fmt.Errorf("repo.IsFavorite: %v", err)
	} else if !fav {
		return errors.New("already unfavorite")
	}

	wordsIds, err := service.GetWords(ctx, movieId)
	if err != nil {
		return fmt.Errorf("service.GetWords: %v", err)
	}

	if err := service.vocabulary.DeleteWords(ctx, userId, wordsIds...); err != nil {
		return fmt.Errorf("vocabulary.AddWords: %v", err)
	}

	if err := service.repo.RemoveUser(ctx, movieId, userId); err != nil {
		return fmt.Errorf("repo.AddUser: %v", err)
	}

	return nil
}

func (service *Service) Search(ctx context.Context, query string) ([]Movie, error) {
	return service.repo.Search(ctx, query)
}

func (service *Service) GetUserFavorites(ctx context.Context, userId int) ([]Movie, error) {
	return service.repo.GetUserFavorites(ctx, userId)
}
