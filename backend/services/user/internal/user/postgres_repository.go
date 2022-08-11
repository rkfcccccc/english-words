package user

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

const usersTable = "users"

type PostgresRepository struct {
	db *pgxpool.Pool
}

func NewPostgresRepository(db *pgxpool.Pool) Repository {
	return &PostgresRepository{db}
}

func (repo *PostgresRepository) Create(ctx context.Context, email, password string) (int, error) {
	var userId int

	query := fmt.Sprintf("INSERT INTO %s (email, password, register_date) VALUES ($1, $2, $3) RETURNING id", usersTable)
	err := pgxscan.Get(ctx, repo.db, &userId, query, email, password, time.Now())

	if err != nil {
		return -1, fmt.Errorf("pgxscan.Get: %v", err)
	}

	return userId, nil
}

func (repo *PostgresRepository) UpdatePassword(ctx context.Context, userId int, password string) error {
	query := fmt.Sprintf("UPDATE %s SET password=$2 WHERE id=$1", usersTable)
	_, err := repo.db.Exec(ctx, query, userId, password)

	if err != nil {
		return fmt.Errorf("db.Exec: %v", err)
	}

	return nil
}

func (repo *PostgresRepository) GetById(ctx context.Context, userId int) (*User, error) {
	var user User

	query := fmt.Sprintf("SELECT * FROM %s WHERE id=$1", usersTable)
	err := pgxscan.Get(ctx, repo.db, &user, query, userId)

	if errors.Is(err, pgx.ErrNoRows) {
		return nil, nil
	}

	if err != nil {
		return nil, fmt.Errorf("pgxscan.Get: %v", err)
	}

	return &user, nil
}

func (repo *PostgresRepository) GetByEmail(ctx context.Context, email string) (*User, error) {
	var user User

	query := fmt.Sprintf("SELECT * FROM %s WHERE email=$1", usersTable)
	err := pgxscan.Get(ctx, repo.db, &user, query, email)

	if errors.Is(err, pgx.ErrNoRows) {
		return nil, nil
	}

	if err != nil {
		return nil, fmt.Errorf("pgxscan.Get: %v", err)
	}

	return &user, nil
}

func (repo *PostgresRepository) Delete(ctx context.Context, userId int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1", usersTable)
	_, err := repo.db.Exec(ctx, query, userId)

	if err != nil {
		return fmt.Errorf("repo.db.Exec: %v", err)
	}

	return nil
}
