package dictionary

import (
	"context"
)

type Repository interface {
	CreateWordIndex(ctx context.Context) error

	Create(ctx context.Context, entry *WordEntry) (string, error)

	SetPictures(ctx context.Context, wordId string, pictures []SourcedPicture) error

	GetById(ctx context.Context, wordId string) (*WordEntry, error)
	GetByWord(ctx context.Context, word string) (*WordEntry, error)

	Delete(ctx context.Context, wordId string) error
}
