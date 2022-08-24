package dictionary

import (
	"context"

	. "github.com/rkfcccccc/english_words/shared_pkg/services/dictionary/models"
)

type Repository interface {
	CreateWordIndex(ctx context.Context) error

	Create(ctx context.Context, entry *WordEntry) (string, error)
	SetPictures(ctx context.Context, wordId string, pictures []SourcedPicture) error

	GetById(ctx context.Context, wordId string) (*WordEntry, error)
	GetByWord(ctx context.Context, word string) (*WordEntry, error)

	Search(ctx context.Context, query string) ([]*WordEntry, error)

	Delete(ctx context.Context, wordId string) error
}
