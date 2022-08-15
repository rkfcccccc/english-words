package movie

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/rkfcccccc/english_words/services/movie/pkg/srt"
	"github.com/rkfcccccc/english_words/shared_pkg/services/dictionary"
)

type Service struct {
	repo Repository
	dict *dictionary.Client
}

func NewService(repo Repository, dict *dictionary.Client) *Service {
	return &Service{repo, dict}
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
		wordId, err := service.dict.Create(ctx, word)

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
	return service.repo.AddUser(ctx, movieId, userId)
}

func (service *Service) RemoveUser(ctx context.Context, movieId int, userId int) error {
	return service.repo.RemoveUser(ctx, movieId, userId)
}

func (service *Service) Search(ctx context.Context, query string) ([]Movie, error) {
	return service.repo.Search(ctx, query)
}

func (service *Service) GetUserFavorites(ctx context.Context, userId int) ([]Movie, error) {
	return service.repo.GetUserFavorites(ctx, userId)
}
