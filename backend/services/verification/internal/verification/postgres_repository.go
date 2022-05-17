package verification

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

const verificationTable = "verifications"

type PostgresRepository struct {
	db *pgxpool.Pool
}

func NewPostgresRepository(db *pgxpool.Pool) Repository {
	return &PostgresRepository{db}
}

func (repo *PostgresRepository) Create(ctx context.Context, email string, typeId int, code, attempts int, ttl time.Duration) (string, error) {
	var requestId string

	query := fmt.Sprintf("INSERT INTO %s (email, type_id, code, attempts, expire_time) VALUES ($1, $2, $3, $4, $5) RETURNING id", verificationTable)
	err := pgxscan.Get(ctx, repo.db, &requestId, query, email, typeId, code, attempts, time.Now().Add(ttl))

	if err != nil {
		return "", fmt.Errorf("pgxscan.Get: %v", err)
	}

	return requestId, nil
}

func (repo *PostgresRepository) GetById(ctx context.Context, requestId string) (*Entry, error) {
	var entry Entry

	query := fmt.Sprintf("SELECT * FROM %s WHERE id=$1 and expire_time > $2", verificationTable)
	err := pgxscan.Get(ctx, repo.db, &entry, query, requestId, time.Now())

	if errors.Is(err, pgx.ErrNoRows) {
		return nil, nil
	}

	if err != nil {
		return nil, fmt.Errorf("pgxscan.Get: %v", err)
	}

	return &entry, nil
}

func (repo *PostgresRepository) GetByEmail(ctx context.Context, email string, typeId int) ([]*Entry, error) {
	var entries []*Entry

	query := fmt.Sprintf("SELECT * FROM %s WHERE type_id=$1 and email=$2 and expire_time > $3", verificationTable)
	err := pgxscan.Select(ctx, repo.db, &entries, query, typeId, email, time.Now())

	if err != nil {
		return nil, fmt.Errorf("pgxscan.Get: %v", err)
	}

	return entries, nil
}

func (repo *PostgresRepository) SetAttempts(ctx context.Context, requestId string, attempts int) error {
	query := fmt.Sprintf("UPDATE %s SET attempts=$2 WHERE id=$1", verificationTable)
	_, err := repo.db.Exec(ctx, query, requestId, attempts)

	if err != nil {
		return fmt.Errorf("db.Exec: %v", err)
	}

	return nil
}

func (repo *PostgresRepository) Delete(ctx context.Context, requestId string) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1", verificationTable)
	_, err := repo.db.Exec(ctx, query, requestId)

	if err != nil {
		return fmt.Errorf("db.Exec: %v", err)
	}

	return nil
}

func (repo *PostgresRepository) DeleteAllExpired(ctx context.Context) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE expire_time < $1", verificationTable)
	_, err := repo.db.Exec(ctx, query, time.Now())

	if err != nil {
		return fmt.Errorf("db.Exec: %v", err)
	}

	return nil
}
