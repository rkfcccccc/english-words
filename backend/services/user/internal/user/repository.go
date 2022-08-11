package user

import "context"

type Repository interface {
	Create(ctx context.Context, email, password string) (int, error)

	UpdatePassword(ctx context.Context, userId int, password string) error

	GetById(ctx context.Context, userId int) (*User, error)
	GetByEmail(ctx context.Context, email string) (*User, error)

	Delete(ctx context.Context, userId int) error
}
