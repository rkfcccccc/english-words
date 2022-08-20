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

func (repo *postgresRepository) Create(ctx context.Context, movie *Movie, words []string) (int, error) {
	if len(words) == 0 {
		return -1, fmt.Errorf("no words given")
	}

	tx, err := repo.db.Begin(ctx)
	if err != nil {
		return -1, fmt.Errorf("db.Begin: %v", err)
	}

	var movieId int
	query := fmt.Sprintf("INSERT INTO %s (imdb_id, title, year, poster_url) VALUES ($1, $2, $3, $4) RETURNING id", moviesTbl)
	if err := pgxscan.Get(ctx, tx, &movieId, query, movie.ImdbId, movie.Title, movie.Year, movie.PosterUrl); err != nil {
		return -1, fmt.Errorf("pgxscan.Get: %v", err)
	}

	for chunk := 0; chunk*chunkSize < len(words); chunk++ {
		start := chunk * chunkSize

		values := []interface{}{}
		queryValues := []string{}

		for i := 0; start+i < len(words) && i < chunkSize; i++ {
			values = append(values, movieId, words[start+i])
			queryValues = append(queryValues, fmt.Sprintf("($%d, $%d)", 2*i+1, 2*i+2))
		}

		query = fmt.Sprintf("INSERT INTO %s (movie_id, word_id) VALUES %s", moviesWordsTbl, strings.Join(queryValues, ", "))
		if _, err = tx.Exec(ctx, query, values...); err != nil {
			tx.Rollback(ctx)
			return -1, fmt.Errorf("db.Exec: %v", err)
		}
	}

	if err := tx.Commit(ctx); err != nil {
		return -1, fmt.Errorf("tx.Commit: %v", err)
	}

	return movieId, nil
}

func (repo *postgresRepository) Get(ctx context.Context, movieId int) (*Movie, error) {
	var result Movie

	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", moviesTbl)
	err := pgxscan.Get(ctx, repo.db, &result, query, movieId)

	if errors.Is(err, pgx.ErrNoRows) {
		return nil, nil
	}

	if err != nil {
		return nil, fmt.Errorf("pgxscan.Get: %v", err)
	}

	return &result, nil
}

func (repo *postgresRepository) GetWords(ctx context.Context, movieId int) ([]string, error) {
	var result []string

	query := fmt.Sprintf("SELECT word_id FROM %s WHERE movie_id = $1", moviesWordsTbl)
	err := pgxscan.Select(ctx, repo.db, &result, query, movieId)

	if errors.Is(err, pgx.ErrNoRows) {
		return []string{}, nil
	}

	if err != nil {
		return nil, fmt.Errorf("pgxscan.Get: %v", err)
	}

	return result, nil
}

func (repo *postgresRepository) Delete(ctx context.Context, movieId int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1", moviesTbl)
	_, err := repo.db.Exec(ctx, query, movieId)

	if err != nil {
		return fmt.Errorf("db.Exec: %v", err)
	}

	return nil
}

func (repo *postgresRepository) AddUser(ctx context.Context, movieId int, userId int) error {
	query := fmt.Sprintf("INSERT INTO %s (movie_id, user_id) VALUES ($1, $2)", moviesUsersTbl)
	_, err := repo.db.Exec(ctx, query, movieId, userId)

	if err != nil {
		return fmt.Errorf("db.Exec: %v", err)
	}

	return nil
}

func (repo *postgresRepository) RemoveUser(ctx context.Context, movieId int, userId int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE movie_id=$1 and user_id=$2", moviesUsersTbl)
	_, err := repo.db.Exec(ctx, query, movieId, userId)

	if err != nil {
		return fmt.Errorf("db.Exec: %v", err)
	}

	return nil
}

func (repo *postgresRepository) IsFavorite(ctx context.Context, movieId int, userId int) (bool, error) {
	var result bool
	query := fmt.Sprintf("SELECT true FROM %s WHERE movie_id=$1 and user_id=$2", moviesUsersTbl)

	err := pgxscan.Get(ctx, repo.db, &result, query, movieId, userId)
	if errors.Is(err, pgx.ErrNoRows) {
		return false, nil
	}

	if err != nil {
		return false, fmt.Errorf("db.Exec: %v", err)
	}

	return result, nil
}

func (repo *postgresRepository) Search(ctx context.Context, searchQuery string) ([]Movie, error) {
	var result []Movie

	searchQuery = "%" + searchQuery + "%"
	query := fmt.Sprintf(`
	SELECT m.* FROM %s m LEFT JOIN %s mu ON m.id = mu.movie_id
	WHERE lower(title) LIKE $1
	GROUP BY id ORDER BY count(mu.user_id) desc`, moviesTbl, moviesUsersTbl)

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
