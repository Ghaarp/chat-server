package chat

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
)

type repo struct {
	pool *pgxpool.Pool
}

type request struct {
	query string
	args  []interface{}
}

func CreateRepository(ctx context.Context, dsn string) (*repo, error) {
	repository := &repo{}
	err := repository.openPool(ctx, dsn)
	return repository, err
}

func (rep *repo) openPool(ctx context.Context, dsn string) error {
	pool, err := pgxpool.Connect(ctx, dsn)
	if err != nil {
		return fmt.Errorf("failed to connect to database: %v", err)
	}
	rep.pool = pool
	return nil
}

func (repo *repo) ClosePool(ctx context.Context) {
	if repo.pool != nil {
		repo.pool.Close()
	}
}
