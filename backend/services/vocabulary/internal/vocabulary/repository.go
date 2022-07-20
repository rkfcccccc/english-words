package vocabulary

import "context"

type Repository interface {
	GetChallenge(ctx context.Context, userId int, challengeNo int) (*WordData, error)

	PromoteWord(ctx context.Context, userId int, wordId string, challengeNo int) error
	ResistWord(ctx context.Context, userId int, wordId string, challengeNo int) error

	AddCount(ctx context.Context, userId int, wordId string, delta int) error

	SetAlreadyLearned(ctx context.Context, userId int, wordId string, isAlreadyLearned bool) error
}
