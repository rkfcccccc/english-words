package vocabulary

import (
	"context"
	"fmt"
)

type Service struct {
	repo  Repository
	state State
}

func NewService(repo Repository, state State) *Service {
	return &Service{repo, state}
}

func (service *Service) GetChallenge(ctx context.Context, userId int) (*WordData, error) {
	challengeNo, err := service.state.Get(ctx, userId)
	if err != nil {
		return nil, fmt.Errorf("state.Get: %v", err)
	}

	return service.repo.GetChallenge(ctx, userId, challengeNo)
}

func (service *Service) PromoteWord(ctx context.Context, userId int, wordId string) error {
	challengeNo, err := service.state.Increment(ctx, userId)
	if err != nil {
		return fmt.Errorf("state.Increment: %v", err)
	}

	return service.repo.PromoteWord(ctx, userId, wordId, challengeNo)
}

func (service *Service) ResistWord(ctx context.Context, userId int, wordId string) error {
	challengeNo, err := service.state.Increment(ctx, userId)
	if err != nil {
		return fmt.Errorf("state.Increment: %v", err)
	}

	return service.repo.ResistWord(ctx, userId, wordId, challengeNo)
}

func (service *Service) SetAlreadyLearned(ctx context.Context, userId int, wordId string, isAlreadyLearned bool) error {
	_, err := service.state.Increment(ctx, userId)
	if err != nil {
		return fmt.Errorf("state.Increment: %v", err)
	}

	return service.repo.SetAlreadyLearned(ctx, userId, wordId, isAlreadyLearned)
}

func (service *Service) AddWord(ctx context.Context, userId int, wordId string) error {
	return service.repo.AddCount(ctx, userId, wordId, 1)
}

func (service *Service) DeleteWord(ctx context.Context, userId int, wordId string) error {
	return service.repo.AddCount(ctx, userId, wordId, -1)
}
