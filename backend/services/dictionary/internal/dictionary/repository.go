package dictionary

import (
	"context"

	models "github.com/rkfcccccc/english_words/shared_pkg/services/dictionary/models"
)

type Repository interface {
	CreateWordIndex(ctx context.Context) error

	Create(ctx context.Context, entry *models.WordEntry) (string, error)
	SetPictures(ctx context.Context, wordId string, pictures []models.SourcedPicture) error

	GetById(ctx context.Context, wordId string) (*models.WordEntry, error)
	GetByWord(ctx context.Context, word string) (*models.WordEntry, error)

	Search(ctx context.Context, query string) ([]*models.WordEntry, error)

	Delete(ctx context.Context, wordId string) error
}
