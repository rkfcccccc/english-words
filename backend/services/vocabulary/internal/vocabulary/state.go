package vocabulary

import "context"

type State interface {
	Get(ctx context.Context, userId int) (int, error)
	Increment(ctx context.Context, userId int) (int, error)
}
