package dictionaryapi

import (
	"context"
)

type Client interface {
	GetWordEntry(ctx context.Context, language, word string) (*Entry, error)
}
