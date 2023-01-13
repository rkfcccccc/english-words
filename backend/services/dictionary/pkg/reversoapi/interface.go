package reversoapi

import "context"

type Client interface {
	GetTranslation(ctx context.Context, input, from, to string) (*translationResponse, error)
}
