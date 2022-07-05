package movie

import "context"

type Repository interface {
	Create(ctx context.Context, movie *Movie, words []string) error
	Delete(ctx context.Context, imdbId string) error

	Get(ctx context.Context, imdbId string) (*Movie, error)
	GetWords(ctx context.Context, imdbId string) ([]string, error)

	AddUser(ctx context.Context, imdbId string, userId int) error
	RemoveUser(ctx context.Context, imdbId string, userId int) error

	Search(ctx context.Context, query string) ([]Movie, error)
	GetUserFavorites(ctx context.Context, userId int) ([]Movie, error)
}
