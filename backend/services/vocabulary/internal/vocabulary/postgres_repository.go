package vocabulary

import (
	"context"
	"fmt"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4/pgxpool"
)

type postgresRepository struct {
	db *pgxpool.Pool
}

const vocabularyTbl = "vocabulary"

func NewPostgresRepository(db *pgxpool.Pool) Repository {
	return &postgresRepository{db}
}

func (repo *postgresRepository) GetChallenge(ctx context.Context, userId int, challengeNo int) (*WordData, error) {
	query1 := fmt.Sprintf("select * from %s where user_id=$1 and next_challenge is not null and next_challenge <= $2 and already_learned is false and count > 0 order by next_challenge limit 1", vocabularyTbl)
	query2 := fmt.Sprintf("select * from %s where user_id=$1 and next_challenge is null and already_learned is false and count > 0 order by count desc limit 1", vocabularyTbl)

	query := fmt.Sprintf("(%s) union (%s) order by next_challenge desc nulls last limit 1", query1, query2)

	var words []WordData
	if err := pgxscan.Select(ctx, repo.db, &words, query, userId, challengeNo); err != nil {
		return nil, fmt.Errorf("pgxscan.Get: %v", err)
	}

	if len(words) == 0 {
		return nil, nil
	}

	return &words[0], nil
}

func (repo *postgresRepository) PromoteWord(ctx context.Context, userId int, wordId string, challengeNo int) error {
	tx, err := repo.db.Begin(ctx)

	if err != nil {
		return fmt.Errorf("db.Exec: %v", err)
	}

	var learningStep int
	query := fmt.Sprintf("select learning_step from %s where user_id=$1 and word_id=$2", vocabularyTbl)
	if err := pgxscan.Get(ctx, tx, &learningStep, query, userId, wordId); err != nil {
		tx.Rollback(ctx)
		return fmt.Errorf("pgxscan.Get: %v", err)
	}

	learningStep += 1

	query = fmt.Sprintf("update %s set next_challenge=$3, learning_step=$4 where user_id=$1 and word_id=$2", vocabularyTbl)
	if _, err := tx.Exec(ctx, query, userId, wordId, challengeNo+learningStep*2, learningStep); err != nil {
		tx.Rollback(ctx)
		return fmt.Errorf("db.Exec: %v", err)
	}

	tx.Commit(ctx)
	return nil
}

func (repo *postgresRepository) ResistWord(ctx context.Context, userId int, wordId string, challengeNo int) error {
	query := fmt.Sprintf("update %s set next_challenge=$3, learning_step=0 where user_id=$1 and word_id=$2", vocabularyTbl)
	_, err := repo.db.Exec(ctx, query, userId, wordId, challengeNo)

	if err != nil {
		return fmt.Errorf("db.Exec: %v", err)
	}

	return nil
}

func (repo *postgresRepository) SetAlreadyLearned(ctx context.Context, userId int, wordId string, isAlreadyLearned bool) error {
	query := fmt.Sprintf("update %s set already_learned=$3 where user_id=$1 and word_id=$2", vocabularyTbl)
	_, err := repo.db.Exec(ctx, query, userId, wordId, isAlreadyLearned)

	if err != nil {
		return fmt.Errorf("db.Exec: %v", err)
	}

	return nil
}

func (repo *postgresRepository) AddCount(ctx context.Context, userId int, wordId string, delta int) error {
	tx, err := repo.db.Begin(ctx)

	if err != nil {
		return fmt.Errorf("db.Exec: %v", err)
	}

	var data []WordData
	query := fmt.Sprintf("select * from %s where user_id=$1 and word_id=$2", vocabularyTbl)
	if err := pgxscan.Select(ctx, tx, &data, query, userId, wordId); err != nil {
		tx.Rollback(ctx)
		return fmt.Errorf("pgxscan.Get: %v", err)
	}

	if len(data) == 0 && delta > 0 {
		query = fmt.Sprintf("insert into %s (user_id, word_id, count) values ($1, $2, $3)", vocabularyTbl)

		if _, err := tx.Exec(ctx, query, userId, wordId, delta); err != nil {
			tx.Rollback(ctx)
			return fmt.Errorf("pgxscan.Get: %v", err)
		}
	} else if len(data) != 0 && (data[0].Count+delta != 0 || data[0].Changed()) {
		count := data[0].Count + delta

		if count < 0 {
			count = 0
		}

		query = fmt.Sprintf("update %s set count=$3 where user_id=$1 and word_id=$2", vocabularyTbl)
		if _, err := tx.Exec(ctx, query, userId, wordId, count); err != nil {
			tx.Rollback(ctx)
			return fmt.Errorf("db.Exec: %v", err)
		}
	} else if len(data) != 0 && data[0].Count+delta == 0 && !data[0].Changed() {
		query = fmt.Sprintf("delete from %s where user_id=$1 and word_id=$2", vocabularyTbl)
		if _, err := tx.Exec(ctx, query, userId, wordId); err != nil {
			tx.Rollback(ctx)
			return fmt.Errorf("db.Exec: %v", err)
		}
	}

	tx.Commit(ctx)
	return nil
}
