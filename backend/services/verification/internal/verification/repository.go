package verification

import (
	"context"
	"time"
)

type Repository interface {
	Create(ctx context.Context, email string, typeId int, code, attempts int, ttl time.Duration) (string, error)

	GetByEmail(ctx context.Context, email string, typeId int) ([]*Entry, error)
	GetById(ctx context.Context, requestId string) (*Entry, error)

	SetAttempts(ctx context.Context, requestId string, attempts int) error
	Delete(ctx context.Context, requestId string) error

	DeleteAllExpired(ctx context.Context) error
}
