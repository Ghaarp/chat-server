package chat

import (
	"context"

	sq "github.com/Masterminds/squirrel"
)

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
