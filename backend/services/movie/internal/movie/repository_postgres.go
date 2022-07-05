package movie

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

const moviesTbl = "movies"
const moviesWordsTbl = "movies_words"
const moviesUsersTbl = "movies_users"
const chunkSize = 1000

type postgresRepository struct {
	db *pgxpool.Pool
}

func NewPostgresRepository(db *pgxpool.Pool) Repository {
	return &postgresRepository{db}
}

func (repo *postgresRepository) Create(ctx context.Context, movie *Movie, words []string) error {
	if words == nil || len(words) == 0 {
		return fmt.Errorf("no words given")
	}

	tx, err := repo.db.Begin(ctx)
	if err != nil {
		return fmt.Errorf("db.Begin: %v", err)
	}

	query := fmt.Sprintf("INSERT INTO %s (imdb_id, title, year, poster_url) VALUES ($1, $2, $3, $4)", moviesTbl)
	if _, err = tx.Exec(ctx, query, movie.ImdbId, movie.Title, movie.Year, movie.PosterUrl); err != nil {
		return fmt.Errorf("db.Exec: %v", err)
	}

	for chunk := 0; chunk*chunkSize < len(words); chunk++ {
		start := chunk * chunkSize

		values := []interface{}{}
		queryValues := []string{}

		for i := 0; start+i < len(words) && i < chunkSize; i++ {
			values = append(values, movie.ImdbId, words[start+i])
			queryValues = append(queryValues, fmt.Sprintf("($%d, $%d)", 2*i+1, 2*i+2))
		}

		query = fmt.Sprintf("INSERT INTO %s (imdb_id, word_id) VALUES %s", moviesWordsTbl, strings.Join(queryValues, ", "))
		if _, err = tx.Exec(ctx, query, values...); err != nil {
			tx.Rollback(ctx)
			return fmt.Errorf("db.Exec: %v", err)
		}
	}

	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("tx.Commit: %v", err)
	}

	return err
}

func (repo *postgresRepository) Get(ctx context.Context, imdbId string) (*Movie, error) {
	var result Movie

	query := fmt.Sprintf("SELECT * FROM %s WHERE imdb_id = $1", moviesTbl)
	err := pgxscan.Get(ctx, repo.db, &result, query, imdbId)

	if errors.Is(err, pgx.ErrNoRows) {
		return nil, nil
	}

	if err != nil {
		return nil, fmt.Errorf("pgxscan.Get: %v", err)
	}

	return &result, nil
}

func (repo *postgresRepository) GetWords(ctx context.Context, imdbId string) ([]string, error) {
	var result []string

	query := fmt.Sprintf("SELECT word_id FROM %s WHERE imdb_id = $1", moviesWordsTbl)
	err := pgxscan.Select(ctx, repo.db, &result, query, imdbId)

	if errors.Is(err, pgx.ErrNoRows) {
		return []string{}, nil
	}

	if err != nil {
		return nil, fmt.Errorf("pgxscan.Get: %v", err)
	}

	return result, nil
}

func (repo *postgresRepository) Delete(ctx context.Context, imdbId string) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE imdb_id=$1", moviesTbl)
	_, err := repo.db.Exec(ctx, query, imdbId)

	if err != nil {
		return fmt.Errorf("db.Exec: %v", err)
	}

	return nil
}

func (repo *postgresRepository) AddUser(ctx context.Context, imdbId string, userId int) error {
	query := fmt.Sprintf("INSERT INTO %s (imdb_id, user_id) VALUES ($1, $2)", moviesUsersTbl)
	_, err := repo.db.Exec(ctx, query, imdbId, userId)

	if err != nil {
		return fmt.Errorf("db.Exec: %v", err)
	}

	return nil
}

func (repo *postgresRepository) RemoveUser(ctx context.Context, imdbId string, userId int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE imdb_id=$1 and user_id=$2", moviesUsersTbl)
	_, err := repo.db.Exec(ctx, query, imdbId, userId)

	if err != nil {
		return fmt.Errorf("db.Exec: %v", err)
	}

	return nil
}

func (repo *postgresRepository) Search(ctx context.Context, searchQuery string) ([]Movie, error) {
	var result []Movie

	searchQuery = "%" + searchQuery + "%"
	query := fmt.Sprintf("select * from %s where lower(title) like $1", moviesTbl)

	err := pgxscan.Select(ctx, repo.db, result, query, searchQuery)

	if errors.Is(err, pgx.ErrNoRows) {
		return []Movie{}, nil
	}

	if err != nil {
		return nil, fmt.Errorf("pgxscan.Select: %v", err)
	}

	return result, nil
}

func (repo *postgresRepository) GetUserFavorites(ctx context.Context, userId int) ([]Movie, error) {
	var result []Movie

	query := fmt.Sprintf("select m.* from %s u inner join %s m on m.imdb_id = u.imdb_id where user_id=$1", moviesUsersTbl, moviesTbl)
	err := pgxscan.Select(ctx, repo.db, &result, query, userId)

	if errors.Is(err, pgx.ErrNoRows) {
		return []Movie{}, nil
	}

	if err != nil {
		return nil, fmt.Errorf("pgxscan.Select: %v", err)
	}

	return result, nil
}
