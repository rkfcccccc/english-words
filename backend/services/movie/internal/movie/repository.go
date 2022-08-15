package movie

import "context"

type Repository interface {
	Create(ctx context.Context, movie *Movie, words []string) (int, error)
	Delete(ctx context.Context, movieId int) error

	Get(ctx context.Context, movieId int) (*Movie, error)
	GetWords(ctx context.Context, movieId int) ([]string, error)

	AddUser(ctx context.Context, movieId int, userId int) error
	RemoveUser(ctx context.Context, movieId int, userId int) error

	Search(ctx context.Context, query string) ([]Movie, error)
	GetUserFavorites(ctx context.Context, userId int) ([]Movie, error)
}
