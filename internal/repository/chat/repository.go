package chat

import (
	"context"
	"fmt"

	"github.com/Ghaarp/chat-server/internal/repository/chat/model"
	"github.com/jackc/pgx/v4/pgxpool"

	sq "github.com/Masterminds/squirrel"
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

func (repo *repo) Create(ctx context.Context, data *model.CreateRequest) (int64, error) {
	queryBuilder := sq.Insert("chats").PlaceholderFormat(sq.Dollar).
		Columns("author", "label").
		Values(data.Author, data.ChatName).
		Suffix("RETURNING id")

	query, args, err := queryBuilder.ToSql()
	if err != nil {
		return 0, err
	}

	queryBuilderUsernames := sq.Insert("chat_users").PlaceholderFormat(sq.Dollar).
		Columns("chat_id", "user_id")

	var chatid int64

	tx, err := repo.pool.Begin(ctx)
	defer func() {
		if err != nil {
			tx.Rollback(ctx)
		}
	}()

	if err != nil {
		return 0, err
	}

	err = tx.QueryRow(ctx, query, args...).Scan(&chatid)
	if err != nil {
		return 0, err
	}

	for _, user := range data.Users {
		queryBuilderUsernames = queryBuilderUsernames.Values(chatid, user)
	}

	query, args, err = queryBuilderUsernames.ToSql()
	if err != nil {
		return 0, err
	}

	_, err = tx.Exec(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return 0, err
	}

	return chatid, nil
}

func (repo *repo) Delete(ctx context.Context, id int64) error {
	requests := make([]request, 3)

	queryChatBuilder := sq.Delete("chats").PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{"id": id})

	var err error
	requests[0].query, requests[0].args, err = queryChatBuilder.ToSql()
	if err != nil {
		return err
	}

	queryUsersBuilder := sq.Delete("chat_users").PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{"chat_id": id})

	requests[1].query, requests[1].args, err = queryUsersBuilder.ToSql()
	if err != nil {
		return err
	}

	queryMessagesBuilder := sq.Delete("messages").PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{"chat_id": id})

	requests[2].query, requests[2].args, err = queryMessagesBuilder.ToSql()
	if err != nil {
		return err
	}

	tx, err := repo.pool.Begin(ctx)
	if err != nil {
		return err
	}

	for _, request := range requests {
		_, err = tx.Exec(ctx, request.query, request.args...)
		if err != nil {
			return err
		}
	}

	err = tx.Commit(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (repo *repo) SendMessage(ctx context.Context, data *model.SendMessageRequest) error {

	builder := sq.Insert("messages").PlaceholderFormat(sq.Dollar).
		Columns("chat_id", "author", "content").
		Values(data.ChatId, data.From, data.Text)

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	_, err = repo.pool.Query(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
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
