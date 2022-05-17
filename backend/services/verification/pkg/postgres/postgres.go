package postgres

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v4/pgxpool"
)

func NewPool(ctx context.Context, username, password, host, port, database string) *pgxpool.Pool {
	url := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", username, password, host, port, database)
	pool, err := pgxpool.Connect(ctx, url)
	if err != nil {
		log.Fatal(err)
	}

	return pool
}
