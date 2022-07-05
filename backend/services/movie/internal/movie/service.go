package movie

import (
	"context"
)

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo}
}

func (service *Service) Create(ctx context.Context, movie *Movie, words []string) error {
	return service.repo.Create(ctx, movie, words)
}

func (service *Service) Delete(ctx context.Context, imdbId string) error {
	return service.repo.Delete(ctx, imdbId)
}

func (service *Service) Get(ctx context.Context, imdbId string) (*Movie, error) {
	return service.repo.Get(ctx, imdbId)
}

func (service *Service) GetWords(ctx context.Context, imdbId string) ([]string, error) {
	return service.repo.GetWords(ctx, imdbId)
}

func (service *Service) AddUser(ctx context.Context, imdbId string, userId int) error {
	return service.repo.AddUser(ctx, imdbId, userId)
}

func (service *Service) RemoveUser(ctx context.Context, imdbId string, userId int) error {
	return service.repo.RemoveUser(ctx, imdbId, userId)
}

func (service *Service) Search(ctx context.Context, query string) ([]Movie, error) {
	return service.repo.Search(ctx, query)
}

func (service *Service) GetUserFavorites(ctx context.Context, userId int) ([]Movie, error) {
	return service.repo.GetUserFavorites(ctx, userId)
}
