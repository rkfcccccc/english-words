package pictureapi

import "context"

type Repository interface {
	Search(ctx context.Context, query string) ([]Picture, error)
	GetName() string
}
