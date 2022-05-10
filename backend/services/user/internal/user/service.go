package user

import (
	"context"
)

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo}
}

// TODO: maybe i should add here some more validation
func (service *Service) Create(ctx context.Context, email, password string) (int, error) {
	return service.repo.Create(ctx, email, password)
}

func (service *Service) GetById(ctx context.Context, userId int) (*User, error) {
	return service.repo.GetById(ctx, userId)
}

func (service *Service) GetByEmail(ctx context.Context, email string) (*User, error) {
	return service.repo.GetByEmail(ctx, email)
}

func (service *Service) Delete(ctx context.Context, userId int) error {
	return service.repo.Delete(ctx, userId)
}
